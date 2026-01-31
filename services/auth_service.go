package services

import (
	"github.com/MehmoodNadeemKhan1/CRUD_API/data/request"
	"github.com/MehmoodNadeemKhan1/CRUD_API/data/response"
)

type AuthService interface {
	Create(auth request.AuthCreateRequest)
	Update(auth request.AuthUpdateRequest)
	Delete(authId int)
	FindById(authId int) response.AuthResponse
	FindAll() []response.AuthResponse
}
