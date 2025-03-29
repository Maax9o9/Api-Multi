package application

import (
    "Multi/src/user/domain/entities"
    "Multi/src/user/domain"
)

type CreateUserUseCase struct {
    repo domain.UserRepository
}

func NewCreateUserUseCase(repo domain.UserRepository) *CreateUserUseCase {
    return &CreateUserUseCase{
        repo: repo,
    }
}

func (uc *CreateUserUseCase) CreateUser(user *entities.User) error {
    return uc.repo.Create(user)
}