package model

import (
	"time"

	"github.com/google/uuid"
)

// Task represents a cron task
type Task struct {
	ID        int64         `json:"id"`
	UserID    int64         `json:"user_id"`
	URLID     uuid.UUID     `json:"url_uuid"`
	Frequency time.Duration `json:"frequency"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	DeletedAt time.Time     `json:"deleted_at"`
}
