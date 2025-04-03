package application

import (
    "Multi/src/user/domain"
    "Multi/src/user/domain/entities"
)

type ShowUserByIDUseCase struct {
    repo domain.UserRepository
}

func NewShowUserByIDUseCase(repo domain.UserRepository) *ShowUserByIDUseCase {
    return &ShowUserByIDUseCase{
        repo: repo,
    }
}

func (uc *ShowUserByIDUseCase) GetUserByID(id int) (*entities.User, error) {
    return uc.repo.GetByID(id)
}