package application

import (
    "Multi/src/user/domain"
    "Multi/src/user/domain/entities"
)

type ShowAllUsersUseCase struct {
    repo domain.UserRepository
}

func NewShowAllUsersUseCase(repo domain.UserRepository) *ShowAllUsersUseCase {
    return &ShowAllUsersUseCase{
        repo: repo,
    }
}

func (uc *ShowAllUsersUseCase) GetAllUsers() ([]entities.User, error) {
    return uc.repo.GetAll()
}