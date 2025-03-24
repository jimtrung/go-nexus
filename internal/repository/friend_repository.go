package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jimtrung/go-nexus/internal/domain"
)

type FriendRepository struct {
	conn *pgx.Conn
}

func NewFriendRepository(conn *pgx.Conn) *FriendRepository {
	return &FriendRepository{
		conn: conn,
	}
}

func (r *FriendRepository) GetAll(userID uint) ([]domain.Friend, error) {
	rows, err := r.conn.Query(
		context.Background(), 
		`SELECT sender_id, receiver_id, status, created_at, updated_at FROM friends
		WHERE sender_id = $1 OR receiver_id = $1`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	friends := []domain.Friend{}
	for rows.Next() {
		var friend domain.Friend
		err := rows.Scan(&friend.SenderID, &friend.ReceiverID, &friend.Status, &friend.CreatedAt, &friend.UpdatedAt)
		if err != nil {
			return nil, err
		}
		friends = append(friends, friend)
	}

	return friends, nil
}

func (r *FriendRepository) Create(friend *domain.Friend) error {
	_, err := r.conn.Exec(
		context.Background(),
		`INSERT INTO friends (sender_id, receiver_id) VALUES ($1, $2)`,
		friend.SenderID,
		friend.ReceiverID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *FriendRepository) Update(friend *domain.Friend) error {
	_, err := r.conn.Exec(
		context.Background(),
		`UPDATE friends SET status = $1 
		WHERE sender_id = $2 AND receiver_id = $3
		OR sender_id = $3 AND receiver_id = $2`,
		friend.Status,
		friend.SenderID,
		friend.ReceiverID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *FriendRepository) Delete(friend *domain.Friend) error {
	_, err := r.conn.Exec(
		context.Background(),
		`DELETE FROM friends 
		WHERE sender_id = $1 AND receiver_id = $2
		OR sender_id = $2 AND receiver_id = $1`,
		friend.SenderID,
		friend.ReceiverID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *FriendRepository) Cancel(friend *domain.Friend) error {
	_, err := r.conn.Exec(
		context.Background(),
		`DELETE FROM friends WHERE sender_id = $1 AND receiver_id = $2`,
		friend.SenderID,
		friend.ReceiverID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *FriendRepository) GetRequests(userID uint) ([]domain.Friend, error) {
	rows, err := r.conn.Query(
		context.Background(),
		`SELECT sender_id, receiver_id, status, created_at, updated_at FROM friends
		WHERE receiver_id = $1 AND status = 'pending'`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	friends := []domain.Friend{}
	for rows.Next() {
		var friend domain.Friend
		err := rows.Scan(&friend.SenderID, &friend.ReceiverID, &friend.Status, &friend.CreatedAt, &friend.UpdatedAt)
		if err != nil {
			return nil, err
		}
		friends = append(friends, friend)
	}

	return friends, nil
}