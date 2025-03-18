package models

import (
	"time"

	"github.com/google/uuid"
)

// Room struct represents the rooms table in postgresql
type Room struct {
	// field	type	json_key_when_serialized	
	ID	uuid.UUID	`json:"id"`
	HostID	uuid.UUID	`json:"host_id"`
	Name	string	`json:"name"`
	CreatedAt	time.Time	`json:"created_at"`
}
