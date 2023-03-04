/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Sat Mar 04 2023 9:48:51 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package manager

import "github.com/eulbyvan/go-user-management/repository"

type RepoManager interface {
	UserRepo() repository.UserRepo
}

type repositoryManager struct {
	infraManager InfraManager
}

func (r *repositoryManager) UserRepo() repository.UserRepo {
	return repository.NewUserRepository(r.infraManager.DbConn())
}

func NewRepoManager(manager InfraManager) RepoManager {
	return &repositoryManager{
		infraManager: manager,
	}
}