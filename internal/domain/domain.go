package domain

import "time"

// User represents a human identity in a company.
type User struct {
	ID       string
	Name     string
	Language string
}

// Assistant represents the personal AI identity linked to one user.
type Assistant struct {
	ID     string
	UserID string
}

// Domain represents a company workspace boundary.
type Domain struct {
	ID   string
	Name string
}

// Message represents an A2A communication event.
type Message struct {
	ID          string
	From        string
	To          string
	Content     string
	CreatedAt   time.Time
}
