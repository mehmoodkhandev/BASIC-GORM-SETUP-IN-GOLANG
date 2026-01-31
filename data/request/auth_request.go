package request

type AuthCreateRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
	Email    string `validate:"required,email" json:"email"`
	Phone    string `validate:"required,numeric,len=11" json:"phone"`
	Role     string `validate:"required,oneof=user admin moderator" json:"role"`
	Status   string `validate:"required,oneof=ACTIVE INACTIVE BANNED" json:"status"`
}

type AuthUpdateRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
	Email    string `validate:"required,email" json:"email"`
	Phone    string `validate:"required,numeric,len=11" json:"phone"`
	Role     string `validate:"required,oneof=user admin moderator" json:"role"`
	Status   string `validate:"required,oneof=ACTIVE INACTIVE BANNED" json:"status"`
	ID       uint   `validate:"required" json:"id"`
}
