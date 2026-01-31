package repository

import "github.com/MehmoodNadeemKhan1/CRUD_API/model"

type AuthRepository interface {
	Save(auth model.AuthModel) error
	Update(auth model.AuthModel) error
	Delete(id uint) error
	GetAll() ([]model.AuthModel, error)
	GetByID(id uint) (model.AuthModel, error)
}


