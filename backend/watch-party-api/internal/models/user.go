package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct represents the users table in postgresql
type User struct {
	// field	type	json_key_when_serialized	
	ID	uuid.UUID	`json:"id"`
	Username	string	`json:"username"`
	Email	string	`json:"email"`
	Password	string	`json:"-"` //Hides password in json responses
	CreatedAt	time.Time	`json:"created_at"`
}
