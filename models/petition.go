package models

import "time"

// NOTE: Anonymous users can only send 2 requests in total.
// each request has to wait 24 hours.

type Petition struct {
	ID        string
	UserID    string // nullable, can be anonymous
	GameID    string // fk ulid
	Status    string // pending, received, testing, completed
	UpdatedAt time.Time
	CreatedAt time.Time
}
