package repository

import (
	"book-management-api/structs"
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser - Simpan user baru dengan password yang sudah di-hash
func CreateUser(db *sql.DB, user structs.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (username, password, created_at, created_by) 
              VALUES ($1, $2, $3, $4)`

	_, err = db.Exec(query, user.Username, string(hashedPassword), time.Now(), user.CreatedBy)
	return err
}

// GetUserByUsername - Ambil user berdasarkan username
func GetUserByUsername(db *sql.DB, username string) (*structs.User, error) {
	var user structs.User
	query := `SELECT id, username, password FROM users WHERE username = $1`

	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("User not found in database")
			return nil, nil
		}
		fmt.Println("Database error:", err)
		return nil, err
	}

	fmt.Println("User found:", user.Username)
	fmt.Println("Stored Hash:", user.Password)
	return &user, nil
}
