package postgres

import (
	"database/sql"

	"github.com/deevarindu/final-project-2/httpserver/repositories"
	"github.com/deevarindu/final-project-2/httpserver/repositories/models"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetUsers() (*[]models.User, error) {
	query := `SELECT id, username, email, password, age FROM users`

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return &users, nil
}

func (u *userRepository) GetUser(id string) (*models.User, error) {
	query := `SELECT id, username, email, password, age FROM users WHERE id=$1`

	var user models.User
	err := u.db.QueryRow(query, id).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (id, username, email, password, age) 
	VALUES ($1, $2, $3, $4, $5)`

	stmt, err := u.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.Username, user.Email, user.Password, user.Age)

	return err
}

func (u *userRepository) UpdateUser(user *models.User) error {
	query := `UPDATE users SET username=$1, email=$2, password=$3, age=$4 WHERE id=$5`

	stmt, err := u.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, user.Password, user.Age, user.Id)

	return err
}

func (u *userRepository) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id=$1`

	stmt, err := u.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}
