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
	GetAll() any
	GetOne(id int) any
	Add(newUser *entity.User) any
	Edit(user *entity.User) any
	Remove(id int) any
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

func (u *userUsecase) Add(newUser *entity.User) any {
	return u.userRepo.Create(newUser)
}

func (u *userUsecase) Edit(user *entity.User) any {
	return u.userRepo.Update(user)
}

func (u *userUsecase) Remove(id int) any {
	return u.userRepo.Delete(id)
}

func NewUserUsecase(userRepo repository.UserRepo) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}