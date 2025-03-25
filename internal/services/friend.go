package services

import (
	"github.com/jimtrung/go-nexus/internal/domain"
	"github.com/jimtrung/go-nexus/internal/repository"
)

type FriendService struct {
	friendRepository *repository.FriendRepository
}

func NewFriendService(friendRepository *repository.FriendRepository) *FriendService {
	return &FriendService{friendRepository: friendRepository}
}

func (s *FriendService) GetAllFriends(userID uint) ([]domain.Friend, error) {
	return s.friendRepository.GetAll(userID)
}

func (s *FriendService) CreateRequest(friend *domain.Friend) error {
	return s.friendRepository.Create(friend)
}

func (s *FriendService) AcceptRequest(friend *domain.Friend) error {
	return s.friendRepository.Update(friend)
}

func (s *FriendService) RejectRequest(friend *domain.Friend) error {
	return s.friendRepository.Update(friend)
}

func (s *FriendService) CancelRequest(friend *domain.Friend) error {
	return s.friendRepository.Cancel(friend)
}

func (s *FriendService) GetPendingRequests(userID uint) ([]domain.Friend, error) {
	return s.friendRepository.GetRequests(userID)
}

func (s *FriendService) RemoveFriend(friend *domain.Friend) error {
	return s.friendRepository.Delete(friend)
}

func (s *FriendService) GetSentRequests(userID uint) ([]domain.Friend, error) {
	return s.friendRepository.GetSentRequests(userID)
}
