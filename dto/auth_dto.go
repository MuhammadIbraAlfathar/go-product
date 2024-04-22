package dto

type RegisterRequest struct {
	Name                 string `json:"name"`
	UserName             string `json:"user_name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
	Gender               string `json:"gender"`
	Address              string `json:"address"`
}

type LoginRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type LoginResponse struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}
