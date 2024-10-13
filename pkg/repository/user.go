package repository

import (
	"database/sql"
	"fmt"

	"github.com/Ivlay/upstore"
)

type UserSql struct {
	db *sql.DB
}

func NewUserSql(db *sql.DB) *UserSql {
	return &UserSql{db: db}
}

func (r *UserSql) GetAll() ([]upstore.User, error) {
	var users []upstore.User

	query := fmt.Sprintf("select user_id, chat_id, firstname, username, created_at from %s", usersTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user upstore.User

		err := rows.Scan(&user.UserId, &user.ChatId, &user.FirstName, &user.UserName, &user.CreatedAt)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserSql) GetByUserId(id int) (upstore.User, error) {
	var user upstore.User

	query := fmt.Sprintf("select user_id, chat_id, firstname, username, created_at from %s where user_id=$1", usersTable)

	err := r.db.QueryRow(query, id).Scan(&user.UserId, &user.ChatId, &user.FirstName, &user.UserName, &user.CreatedAt)

	return user, err
}

func (r *UserSql) Create(user upstore.User) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (userName, firstName, chat_id, user_id) values ($1, $2, $3, $4) returning user_id", usersTable)
	row := r.db.QueryRow(query, user.UserName, user.FirstName, user.ChatId, user.UserId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserSql) FindOrCreate(user upstore.User) (int, error) {
	u, err := r.GetByUserId(user.UserId)
	switch err {
	case sql.ErrNoRows:
		fmt.Printf("User not found, try to create\n %s", err.Error())
		return r.Create(user)
	default:
		return u.UserId, err
	}
}
