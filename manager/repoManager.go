package manager

import "web-app-crowdfounding/repository"

type RepositoryManager interface {
	UsersRepository() repository.UserRepository
}

type repositoryManager struct {
	infra InfraManager
}

func (r *repositoryManager) UsersRepository() repository.UserRepository {
	return repository.NewUserRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infraManager InfraManager) *repositoryManager {
	return &repositoryManager{infraManager}
}
