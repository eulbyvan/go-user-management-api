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

	"github.com/eulbyvan/go-user-management/model/entity"
)

type UserRepo interface {
	FindAll() any
	FindOne(id int) any
	Create(newUser *entity.User) any
	Update(user *entity.User) any
	Delete(id int) any
}

type userRepo struct {
	db *sql.DB
}

func (r *userRepo) FindAll() any {
	var users []entity.User

	query := "SELECT id, first_name, last_name, email FROM users ORDER BY id"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	return users
}

func (r *userRepo) FindOne(id int) any {
	var userInDb entity.User

	query := "SELECT id, first_name, last_name, email FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)

	err := row.Scan(&userInDb.ID, &userInDb.FirstName, &userInDb.LastName, &userInDb.Email)

	if err != nil {
		log.Println(err)
	}

	// Jika tidak ada, return nil
	if userInDb.ID == 0 {
		return nil
	}

	return userInDb
}

func (r *userRepo) Create(newUser *entity.User) any {
	query := "INSERT INTO users (first_name, last_name, email) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(query, newUser.FirstName, newUser.LastName, newUser.Email)

	if err != nil {
		log.Println(err)
		return nil
	}

	return r.FindAll()
}

func (r *userRepo) Update(user *entity.User) any {
	// Cari user di database
	res := r.FindOne(user.ID)

	// Jika ada maka update user
	query := "UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4"
	_, err := r.db.Exec(query, user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		log.Println(err)
	}

	// Jika berhasil update, return hasil
	return res
}

func (r *userRepo) Delete(id int) any {
	// Cari user di database
	res := r.FindOne(id)

	if res == nil {
		return "User not found"
	}

	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println(err)
		return fmt.Sprintf("Failed to delete user with id: %d", id)
	}

	// return fmt.Sprintf("User with id %d deleted successfully", id)
	return r.FindAll()
}

func NewUserRepository(db *sql.DB) UserRepo {
	repo := new(userRepo)
	repo.db = db
	return repo
}