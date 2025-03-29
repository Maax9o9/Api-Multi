package application

import (
    "Multi/src/user/domain/entities"
    "Multi/src/user/domain"
)

type ShowUsersUseCase struct {
    repo domain.UserRepository
}

func NewShowUsersUseCase(repo domain.UserRepository) *ShowUsersUseCase {
    return &ShowUsersUseCase{
        repo: repo,
    }
}

func (uc *ShowUsersUseCase) GetAllUsers() ([]entities.User, error) {
    return uc.repo.GetAll()
}

func (uc *ShowUsersUseCase) GetUserByID(id int) (*entities.User, error) {
    return uc.repo.GetByID(id)
}