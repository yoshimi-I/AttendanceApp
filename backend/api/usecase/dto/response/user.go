package response

type UserDTO struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	UserKeyKey string `json:"user_key"`
}
