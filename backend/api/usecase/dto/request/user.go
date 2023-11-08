package request

type UserDTO struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	UserKey string `json:"user_key"`
}
