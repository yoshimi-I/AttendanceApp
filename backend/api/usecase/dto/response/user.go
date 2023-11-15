package response

type UserDTO struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	UserKey string `json:"user_key"`
}

type UserStatusDTO struct {
	Status string `json:"status"`
}
