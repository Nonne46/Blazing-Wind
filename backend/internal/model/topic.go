package model

import "time"

type Topic struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	ForumID   uint64    `json:"forum_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}
