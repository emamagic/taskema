package param

type UserProfileRequest struct {
	UserID uint `json:"user_id"`
}

type UserProfileResponse struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Avatar *string `json:"avatar"`
	Email  string  `json:"email"`
}
