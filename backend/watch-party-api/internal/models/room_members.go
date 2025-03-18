package models

import (
	"time"

	"github.com/google/uuid"
)

// RoomMembers struct represents the room_members table in postgresql
type RoomMembers struct {
	// field	type	json_key_when_serialized	
	ID	uuid.UUID	`json:"id"`
	RoomID uuid.UUID `json:"room_id"`
	UserID uuid.UUID	`json:"user_id"`
	CreatedAt	time.Time	`json:"joined_at"`
}
