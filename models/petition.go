package models

import "time"

// NOTE: Anonymous users can only send 2 requests in total.
// each request has to wait 24 hours.

type Petition struct {
	ID     string `json:"id" db:"id"`
	UserID string `json:"user_id" db:"user_id"`

	// NOTE: ask claude
	// UserID *string

	GameID string `json:"game_id" db:"game_id"`
	Status string `json:"status", db:"status"`

	// never use string for dates, always time.Time
	UpdatedAt time.Time `json:"updated_at db:"updated_at"`
	CreatedAt time.Time `json:"created_at db:"created_at"`
}
