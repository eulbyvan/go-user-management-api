/*
 * Author : Ismail Ash Shidiq (https://www.eulbyvan.com)
 * Created on : Fri Mar 03 2023 9:56:04 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/eulbyvan/go-user-management/model"
)

type UserRepo interface {
	FindAll() any
	FindOne(id int) any
	Create(newUser *model.User) string
	Update(user *model.User) string
	Delete(id int) string
}

type userRepo struct {
	db *sql.DB
}

func (r *userRepo) FindAll() any {
	var users []model.User

	query := "SELECT id, first_name, last_name, email FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(users) == 0 {
		return "No data"
	}

	return users
}

func (r *userRepo) FindOne(id int) any {
	var userInDb model.User

	query := "SELECT id, first_name, last_name, email FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)

	err := row.Scan(&userInDb.ID, &userInDb.FirstName, &userInDb.LastName, &userInDb.Email)

	if err != nil {
		log.Println(err)
	}

	if userInDb.ID == 0 {
		return "User not found"
	}

	return userInDb
}

func (r *userRepo) Create(newUser *model.User) string {
	query := "INSERT INTO users (first_name, last_name, email) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(query, newUser.FirstName, newUser.LastName, newUser.Email)

	if err != nil {
		log.Println(err)
		return "Failed to create user"
	}

	return "User created successfully"
}

func (r *userRepo) Update(user *model.User) string {
	// Cari user di database
	res := r.FindOne(user.ID)

	// Jika tidak ada, return pesan
	if res == "User not found" {
		return res.(string)
	}

	// Jika ada maka update user
	query := "UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4"
	_, err := r.db.Exec(query, user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		log.Println(err)
	}

	// Jika berhasil update, return pesan
	return fmt.Sprintf("User with id %d updated successfully", user.ID)
}

func (r *userRepo) Delete(id int) string {
	// Cari user di database
	res := r.FindOne(id)

	// Jika tidak ada, return pesan
	if res == "User not found" {
		return res.(string)
	}

	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println(err)
		return "Failed to delete user"
	}

	return fmt.Sprintf("User with id %d deleted successfully", id)
}

func NewUserRepository(db *sql.DB) UserRepo {
	repo := new(userRepo)
	repo.db = db
	return repo
}