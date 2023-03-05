/*
 * Author : Ismail Ash Shidiq (https://www.eulbyvan.com)
 * Created on : Fri Mar 03 2023 11:02:50 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package usecase

import (
	"github.com/eulbyvan/go-user-management/model/entity"
	"github.com/eulbyvan/go-user-management/repository"
)

type UserUsecase interface {
	GetAll() ([]entity.User, error)
	GetOne(id int) (entity.User, error)
	Add(newUser *entity.User) (entity.User, error)
	Edit(user *entity.User) (entity.User, error)
	Remove(id int) error
}

type userUsecase struct {
	userRepo repository.UserRepo
}

func (u *userUsecase) GetAll() ([]entity.User, error) {
	return u.userRepo.FindAll()
}

func (u *userUsecase) GetOne(id int) (entity.User, error) {
	return u.userRepo.FindOne(id)
}

func (u *userUsecase) Add(newUser *entity.User) (entity.User, error) {
	return u.userRepo.Create(newUser)
}

func (u *userUsecase) Edit(user *entity.User) (entity.User, error) {
	return u.userRepo.Update(user)
}

func (u *userUsecase) Remove(id int) error {
	return u.userRepo.Delete(id)
}

func NewUserUsecase(userRepo repository.UserRepo) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}