/*
 * Author : Ismail Ash Shidiq (https://www.eulbyvan.com)
 * Created on : Fri Mar 03 2023 11:02:50 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package usecase

import (
	"github.com/eulbyvan/go-user-management/model"
	"github.com/eulbyvan/go-user-management/repository"
)

type UserUsecase interface {
	GetAll() any
	GetOne(id int) any
	Add(newUser *model.User) string
	Edit(user *model.User) string
	Remove(id int) string
}

type userUsecase struct {
	userRepo repository.UserRepo
}

func (u *userUsecase) GetAll() any {
	return u.userRepo.FindAll()
}

func (u *userUsecase) GetOne(id int) any {
	return u.userRepo.FindOne(id)
}

func (u *userUsecase) Add(newUser *model.User) string {
	return u.userRepo.Create(newUser)
}

func (u *userUsecase) Edit(user *model.User) string {
	return u.userRepo.Update(user)
}

func (u *userUsecase) Remove(id int) string {
	return u.userRepo.Delete(id)
}

func NewUserUsecase(userRepo repository.UserRepo) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}