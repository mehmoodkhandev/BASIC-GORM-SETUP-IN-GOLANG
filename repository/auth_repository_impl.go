package repository

import (
	"errors"

	"github.com/MehmoodNadeemKhan1/CRUD_API/data/request"
	"github.com/MehmoodNadeemKhan1/CRUD_API/helper"
	"github.com/MehmoodNadeemKhan1/CRUD_API/model"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements AuthRepository.
func (a *AuthRepositoryImpl) Delete(id uint) error {
	var auth model.AuthModel
	result := a.Db.Where("id = ?", id).Delete(&auth)
	helper.ErrorPanic(result.Error)
	return result.Error
}

// GetAll implements AuthRepository.
func (a *AuthRepositoryImpl) GetAll() ([]model.AuthModel, error) {
	var auths []model.AuthModel
	result := a.Db.Find(&auths)
	helper.ErrorPanic(result.Error)
	return auths, result.Error
}

// GetByID implements AuthRepository.
func (a *AuthRepositoryImpl) GetByID(id uint) (model.AuthModel, error) {
	var auth model.AuthModel
	result := a.Db.Where("id = ?", id).First(&auth)
	if result != nil {
		return auth, nil
	} else {
		return auth, errors.New("record not found")
	}

}

// Save implements AuthRepository.
func (a *AuthRepositoryImpl) Save(auth model.AuthModel) error {
	result := a.Db.Create(&auth)
	helper.ErrorPanic(result.Error)
	return result.Error
}

// Update implements AuthRepository.
func (a *AuthRepositoryImpl) Update(auth model.AuthModel) error {
	var updatedAuth = request.AuthUpdateRequest{
		Username: auth.Username,
		Password: auth.Password,
		Email:    auth.Email,
		Phone:    auth.Phone,
		Role:     string(auth.Role),
		Status:   string(auth.Status),
		ID:       auth.ID,
	}
	result := a.Db.Model(&auth).Where("id = ?", updatedAuth.ID).Updates(updatedAuth)
	helper.ErrorPanic(result.Error)
	return result.Error
}

func NewAuthRepositoryImpl(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{Db: db}
}
