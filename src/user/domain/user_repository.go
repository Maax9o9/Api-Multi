package domain

import "Multi/src/user/domain/entities"

type UserRepository interface {
    Create(user *entities.User) error
    GetAll() ([]entities.User, error)
    GetByID(id int) (*entities.User, error)
    GetByUsername(username string) (*entities.User, error)
}