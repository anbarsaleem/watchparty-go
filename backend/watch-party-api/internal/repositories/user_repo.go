package repositories

//This is a repository layer program to allow interaction with PostgreSQL

import (
	"context"
	"log"

	"watch-party-api/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/google/uuid"
)

// function to create an entry in the users table in PostgreSQL
// takes in necessary details and returns ID of the created user and potential error
func CreateUser(db *pgx.Conn, username, email, hashedPassword string) (error) {
	var userID uuid.UUID
	ctx := context.Background()

	args := pgx.NamedArgs{
		"username": username,
		"email": email,
		"hashedPassword": hashedPassword,
	}

	query := `INSERT INTO users (username, emal, password) VALUES (@username, @email, @hashedPassword) RETURNING id` //uses named arguments for inserted vals for readability
	err := db.QueryRow(ctx, query, args).Scan(&userID) // executes query on row and saves userID 

	if err != nil {
		log.Printf("error inserting user into users table")
		return err
	}
	return nil
}

//function to get a user by their email
func GetUserByEmail(db *pgx.Conn, email string) (*models.User, error ){
	var user models.User
	ctx := context.Background()

	args := pgx.NamedArgs{
		"email": email,
	}
	
	query := `SELECT * FROM users WHERE email = @email`
	err := db.QueryRow(ctx, query, args).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		log.Printf("error fetching user by email")
		return nil, err
	}
	return &user, nil
}

// function to get a user by their UUID
func GetUserByID() {

}
