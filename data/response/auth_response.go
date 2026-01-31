package response

type AuthResponse struct {
	UserID       int    `json:"userid"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Status       string `json:"status"`
	Role         string `json:"role"`
	CreatedAt    string `json:"created_at"`
	LastModified string `json:"last_modified"`
}
