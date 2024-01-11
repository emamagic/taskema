package param

type UserRegisterRequest struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Avatar   *string `json:"avatar"`
	Password string  `json:"password"`
}

type UserRegisterResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
