package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Follower struct {
	UserId     int64  `json:"user_id"`
	FollowerID int64  `json:"follower_id"`
	CreatedAt  string `json:"created_at"`
}

type FolloweStore struct {
	db *sql.DB
}

func (s *FolloweStore) Follow(ctx context.Context, userID, toFollowUserID int64) error {
	query := `
		INSERT INTO followers (user_id, follower_id) VALUES ($1, $2)
	`
	_, err := s.db.ExecContext(ctx, query, userID, toFollowUserID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrConflict
		}
	}

	return nil
}

func (s *FolloweStore) Unfollow(ctx context.Context, userID, toUnfollowUserID int64) error {
	query := `
		DELETE FROM followers WHERE user_id = $1 AND follower_id = $2
	`

	_, err := s.db.ExecContext(ctx, query, userID, toUnfollowUserID)

	return err
}
