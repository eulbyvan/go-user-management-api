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
	FindAll() ([]entity.User, error)
	FindOne(id int) (entity.User, error)
	Create(newUser *entity.User) (entity.User, error)
	Update(user *entity.User) (entity.User, error)
	Delete(id int) error
}

type userRepo struct {
	db *sql.DB
}

func (r *userRepo) FindAll() ([]entity.User, error) {
    var users []entity.User

    query := "SELECT id, first_name, last_name, email FROM users ORDER BY id"
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    for rows.Next() {
        var user entity.User
        if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}

func (r *userRepo) FindOne(id int) (entity.User, error) {
    var userInDb entity.User

    query := "SELECT id, first_name, last_name, email FROM users WHERE id = $1"
    row := r.db.QueryRow(query, id)

    err := row.Scan(&userInDb.ID, &userInDb.FirstName, &userInDb.LastName, &userInDb.Email)

    if err == sql.ErrNoRows {
        return entity.User{}, fmt.Errorf("user with id %d not found", id)
    } else if err != nil {
        return entity.User{}, err
    }

    return userInDb, nil
}

func (r *userRepo) Create(newUser *entity.User) (entity.User, error) {
	query := "INSERT INTO users (first_name, last_name, email) VALUES ($1, $2, $3) RETURNING id"
	var userID int
	err := r.db.QueryRow(query, newUser.FirstName, newUser.LastName, newUser.Email).Scan(&userID)
	if err != nil {
		log.Println(err)
		return entity.User{}, err
	}

	newUser.ID = userID

	return *newUser, nil
}

func (r *userRepo) Update(user *entity.User) (entity.User, error) {
	query := "UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4 RETURNING id, first_name, last_name, email"
	row := r.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.ID)

	var updatedUser entity.User
	err := row.Scan(&updatedUser.ID, &updatedUser.FirstName, &updatedUser.LastName, &updatedUser.Email)
	if err != nil {
		log.Println(err)
		return entity.User{}, err
	}

	return updatedUser, nil
}

func (r *userRepo) Delete(id int) error {
    query := "DELETE FROM users WHERE id = $1"
    result, err := r.db.Exec(query, id)
    if err != nil {
        log.Println(err)
        return err
    }
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Println(err)
        return err
    }
    if rowsAffected == 0 {
        return fmt.Errorf("user with id %d not found", id)
    }
    return nil
}


func NewUserRepository(db *sql.DB) UserRepo {
	repo := new(userRepo)
	repo.db = db
	return repo
}