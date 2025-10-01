package auth

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type PhoneLoginRequest struct {
	Phone string `json:"phone" validate:"required"`
}

type PhoneLoginResponse struct {
	SessionId string `json:"session_id"`
}

type PhoneLoginVerifyRequest struct {
	SessionId string `json:"session_id"`
	Code      string `json:"code" validate:"required"`
}

type PhoneLoginVerifyResponse struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}
