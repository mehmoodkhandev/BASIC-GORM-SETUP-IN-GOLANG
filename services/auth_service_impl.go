package services

import (
	"errors"

	"github.com/MehmoodNadeemKhan1/CRUD_API/data/request"
	"github.com/MehmoodNadeemKhan1/CRUD_API/data/response"
	"github.com/MehmoodNadeemKhan1/CRUD_API/helper"
	"github.com/MehmoodNadeemKhan1/CRUD_API/model"
	"github.com/MehmoodNadeemKhan1/CRUD_API/repository"
	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	// You can add fields here if needed, such as a repository or database connection
	AuthRepository repository.AuthRepository
	validate       *validator.Validate
}

func NewAuthServiceImpl(authRepository repository.AuthRepository, validate *validator.Validate) AuthService {

	return &AuthServiceImpl{
		AuthRepository: authRepository,
		validate:       validate,
	}
}

// Create implements AuthService.
func (a *AuthServiceImpl) Create(auth request.AuthCreateRequest) {
	err := a.validate.Struct(auth)
	helper.ErrorPanic(err)
	authModel := model.AuthModel{
		Username: auth.Username,
		Password: auth.Password,
		Email:    auth.Email,
		Phone:    auth.Phone,
		Role:     model.UserRole(auth.Role),
		Status:   model.UserStatus(auth.Status),
	}
	a.AuthRepository.Save(authModel)

}

// Delete implements AuthService.
func (a *AuthServiceImpl) Delete(authId int) {
	a.AuthRepository.Delete(uint(authId))
}

// FindAll implements AuthService.
func (a *AuthServiceImpl) FindAll() []response.AuthResponse {
	result, err := a.AuthRepository.GetAll()
	helper.ErrorPanic(err)
	var authResponses []response.AuthResponse
	for _, value := range result {
		auth := response.AuthResponse{
			UserID:   int(value.ID),
			Username: value.Username,
			Email:    value.Email,
			Phone:    value.Phone,
			Role:     string(value.Role),
			Status:   string(value.Status),
		}
		authResponses = append(authResponses, auth)
	}
	return authResponses
}

// FindById implements AuthService.
func (a *AuthServiceImpl) FindById(authId int) response.AuthResponse {
	AuthData, err := a.AuthRepository.GetByID(uint(authId))
	helper.ErrorPanic(err)
	authResponse := response.AuthResponse{
		UserID:   int(AuthData.ID),
		Username: AuthData.Username,
		Email:    AuthData.Email,
		Phone:    AuthData.Phone,
		Status:   string(AuthData.Status),
		Role:     string(AuthData.Role),
	}
	if authResponse.UserID == 0 {
		helper.ErrorPanic(errors.New("auth not found"))
	}
	return authResponse
}

// Update implements AuthService.
func (a *AuthServiceImpl) Update(auth request.AuthUpdateRequest) {
	authData, err := a.AuthRepository.GetByID(auth.ID)
	helper.ErrorPanic(err)
	if authData.ID == 0 {
		helper.ErrorPanic(errors.New("auth not found"))
	}
	authData.Username = auth.Username
	authData.Email = auth.Email
	authData.Phone = auth.Phone
	authData.Role = model.UserRole(auth.Role)
	authData.Status = model.UserStatus(auth.Status)
	a.AuthRepository.Update(authData)

}
