package response

type UserDTO struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	UserKey string `json:"user_key"`
}

type UserStatusDTO struct {
	Status string `json:"status"`
}
