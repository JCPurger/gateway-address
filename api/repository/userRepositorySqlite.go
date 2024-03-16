package repository

import (
	"database/sql"
	"fmt"

	"github.com/gateway-address/user"
	_ "modernc.org/sqlite"
)

type RepositorySqlite struct {
	db *sql.DB
}

// NewRepository cria uma nova instância do Repository.
func NewRepositorySqlite() (*RepositorySqlite, error) {
	db, err := sql.Open("sqlite", "./db/main.db")
	if err != nil {
		return nil, fmt.Errorf("error opening sqlite database: %s", err)
	}
	return &RepositorySqlite{db: db}, nil
}

func (r *RepositorySqlite) Create(user user.User) (sql.Result, error) {
	stmt, err := r.db.Prepare("INSERT INTO user (first_name, last_name, username, password, email) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return nil, fmt.Errorf("error preparing SQL statement: %s", err)
	}
	defer stmt.Close()

	// Execute SQL statement to insert data
	result, resError := stmt.Exec(user.FirstName, user.LastName, user.UserName, user.Password, user.Email)
	if resError != nil {
		return nil, fmt.Errorf("error executing SQL statement: %s", err)
	}
	return result, nil
}

func (r *RepositorySqlite) GetAll() ([]user.User, error) {
	var users []user.User

	defer r.db.Close()

	allUsers, err := r.db.Query("SELECT * FROM user")
	if err != nil {
		fmt.Println(err)
		return users, err // Return empty slice and error
	}
	defer allUsers.Close()

	for allUsers.Next() {
		var u user.User
		if err := allUsers.Scan(&u.ID, &u.FirstName, &u.LastName, &u.UserName, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt); err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}
	if err := allUsers.Err(); err != nil {
		fmt.Println(err)
		return users, err // Return slice with retrieved users and error
	}
	return users, nil // Return slice of users and no error
}