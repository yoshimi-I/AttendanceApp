package request

type UserDTO struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	UserKeyKey string `json:"user_key"`
}
