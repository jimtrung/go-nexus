package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jimtrung/go-nexus/internal/domain"
	"github.com/jimtrung/go-nexus/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	AuthRepo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{
		AuthRepo: repo,
	}
}

func (s *AuthService) GetUserByID(userID uint) (*domain.User, error) {
	return s.AuthRepo.GetByID(userID)
}

func (s *AuthService) SignUp(req *domain.User) error {
	if !IsValidEmail(req.Email) || !HasMXRecord(req.Email) {
		return fmt.Errorf("Invalid email")
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return err
	}

	req.Password = hashedPassword
	token := GenerateToken()
	req.Token = token

	if err := s.AuthRepo.InsertIntoUsers(req); err != nil {
		return err
	}

	SendVerificationEmail(req.Email, token)
	return nil
}

func (s *AuthService) Login(req *domain.User) (string, error) {
	user, err := s.AuthRepo.GetByUsername(req.Username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", fmt.Errorf("Wrong password")
	}

	return s.CreateSignedToken(user.UserID)
}

func (s *AuthService) ForgotPassword(req *domain.User) error {
	if !IsValidEmail(req.Email) || !HasMXRecord(req.Email) {
		return fmt.Errorf("Invalid email")
	}

	token := GenerateToken()
	req.Token = token

	if err := s.AuthRepo.AddToken(req.Email, req.Token); err != nil {
		return fmt.Errorf("Failed to add token to user: %v", err)
	}

	ResetPasswordEmail(req.Email, req.Token)
	go func() {
		time.Sleep(time.Second * 300)
		s.AuthRepo.DeleteToken(req.Token)
	}()
	return nil
}

func (s *AuthService) CreateSignedToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 4).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", fmt.Errorf("Failed to sign the token")
	}

	return tokenString, nil
}

func (s *AuthService) SignupIfNotExist(email string) (*domain.User, error) {
	userInfo, err := s.AuthRepo.GetByEmail(email)
	if err == nil {
		return userInfo, nil
	}

	randomPassword := GenerateRandomPassword()
	hashedPassword, err := HashPassword(randomPassword)
	if err != nil {
		return &domain.User{}, err
	}

	userInfo = &domain.User{
		Username: email,
		Email:    email,
		Verified: true,
		Password: hashedPassword,
	}
	if err := s.AuthRepo.InsertIntoUsers(userInfo); err != nil {
		return &domain.User{}, err
	}

	return userInfo, nil
}
