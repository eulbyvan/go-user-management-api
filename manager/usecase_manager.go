/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Sat Mar 04 2023 9:48:44 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package manager

import "github.com/eulbyvan/go-user-management/usecase"

type UsecaseManager interface {
	UserUsecase() usecase.UserUsecase
}

type usecaseManager struct {
	repoManager RepoManager
}

func (u *usecaseManager) UserUsecase() usecase.UserUsecase {
	return usecase.NewUserUsecase(u.repoManager.UserRepo())
}

func NewUsecaseManager(rm RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: rm,
	}
}