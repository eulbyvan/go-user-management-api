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