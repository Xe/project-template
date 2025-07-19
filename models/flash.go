package models

import (
	"encoding/gob"
	"time"

	"github.com/google/uuid"
)

func init() {
	gob.Register(Flash{})
}

type FlashKind string

const (
	FlashSuccess FlashKind = "success"
	FlashWarning FlashKind = "warning"
	FlashFailure FlashKind = "failure"
	FlashInfo    FlashKind = "info"
)

type Flash struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Kind      FlashKind `json:"kind"`
	Body      string    `json:"body"` // HTML content
}

// NewFlash creates a new flash message with a generated ID and current timestamp
func NewFlash(kind FlashKind, body string) Flash {
	return Flash{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		Kind:      kind,
		Body:      body,
	}
}
