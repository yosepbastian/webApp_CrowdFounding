package manager

import (
	"web-app-crowdfounding/usecase"
)

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
}

type useCaseManager struct {
	repomanager RepositoryManager
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repomanager.UsersRepository())
}

func NewUseCaseManager(repoManager RepositoryManager) *useCaseManager {
	return &useCaseManager{repoManager}
}
