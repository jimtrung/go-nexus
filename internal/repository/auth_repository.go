package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jimtrung/go-nexus/internal/domain"
)

type UserRepository struct {
    conn *pgx.Conn
}

func (u *UserRepository) InsertIntoUsers(user *domain.User) error {
    _, err := u.conn.Exec(
        context.Background(),
        `INSERT INTO users(username, email, password) VALUES ($1, $2, $3)`,
    )
    if err != nil {
        return fmt.Errorf("Error inserting user into database: %v", err)
    }
    return nil
}

func (u *UserRepository) GetByUsername(username string) (*domain.User, error) {
    row := u.conn.QueryRow(
        context.Background(),
        `SELECT user_id, username, email, role, verified, created_at, updated_at
        FROM users
        WHERE username = $1`,
        username,
    )

    user := &domain.User{}
    if err := row.Scan(
        &user.UserID, &user.Username, &user.Email, &user.Role, &user.Verified,
        &user.CreatedAt, &user.UpdatedAt,
    ); err != nil {
        return &domain.User{}, fmt.Errorf("Failed to get user %s: %v", username, err)
    }

    return user, nil
}

func (u *UserRepository) GetByEmail(email string) (*domain.User, error) {
    row := u.conn.QueryRow(
        context.Background(),
        `SELECT user_id, username, email, role, verified, created_at, updated_at
        FROM users
        WHERE email = $1`,
        email,
    )

    var user *domain.User
    if err := row.Scan(
        &user.UserID, &user.Username, &user.Email, &user.Role, &user.Verified,
        &user.CreatedAt, &user.UpdatedAt,
    ); err != nil {
        return &domain.User{}, fmt.Errorf("Failed to get user %s: %v", email, err)
    }

    return user, nil
}

func (u *UserRepository) AddToken(email, token string) error {
    _, err := u.conn.Exec(
        context.Background(),
        `UPDATE users SET token = $1 WHERE email = $2`,
        token, email,
    )
    return err
}

func (u *UserRepository) DeleteToken(token string) error {
    _, err := u.conn.Exec(
        context.Background(),
        `DELETE FROM users WHERE token = $1`,
        token,
    )
    return err
}

func (u *UserRepository) Verify(token string) error {
    _, err := u.conn.Exec(
        context.Background(),
        `UPDATE users SET verified = true, token = "" WHERE token = $1`,
        token,
    )
    return err
}

func (u *UserRepository) UpdatePassword(token, newPasswordHash string) error {
    _, err := u.conn.Exec(
        context.Background(),
        `UPDATE users
        SET token = "", password = $1
        WHERE token = $2`,
        newPasswordHash, token,
    )
    return err
}
