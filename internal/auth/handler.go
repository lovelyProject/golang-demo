package auth

import (
	"go/adv-example/configs"
	"go/adv-example/pkg/jwt"
	req "go/adv-example/pkg/req"
	"go/adv-example/pkg/res"
	"net/http"
)

type AuthHandler struct {
	Service *AuthService
	Config  *configs.Config
}

func NewAuthHandler(service *AuthService, config *configs.Config) *AuthHandler {
	return &AuthHandler{
		Service: service,
		Config:  config,
	}
}

func (h *AuthHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /auth/register", h.Register())
	router.HandleFunc("POST /auth/login", h.Login())
	router.HandleFunc("POST /auth/phone-login", h.PhoneLogin())
	router.HandleFunc("POST /auth/phone-login/verify", h.PhoneLoginVerify())
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			res.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := h.Service.Register(body.Email, body.Name, body.Password)
		if err != nil {
			res.Json(w, http.StatusInternalServerError, err.Error())
			return
		}

		token, err := jwt.NewJWT(h.Config.Auth.Secret).CreateToken(jwt.JWTData{Email: user.Email, ID: user.ID})
		if err != nil {
			res.Json(w, http.StatusInternalServerError, err.Error())
			return
		}

		res.Json(w, http.StatusCreated, RegisterResponse{Token: token})
	}
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			res.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		email, err := h.Service.Login(body.Email, body.Password)
		if err != nil {
			res.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		token, err := jwt.NewJWT(h.Config.Auth.Secret).CreateToken(jwt.JWTData{Email: email})
		if err != nil {
			res.Json(w, http.StatusInternalServerError, err.Error())
			return
		}

		res.Json(w, http.StatusOK, LoginResponse{Token: token})
	}
}

func (h *AuthHandler) PhoneLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[PhoneLoginRequest](w, r)
		if err != nil {
			res.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		sessionId, err := h.Service.LoginByPhone(body.Phone)
		if err != nil {
			res.Json(w, http.StatusInternalServerError, err.Error())
			return
		}

		res.Json(w, http.StatusOK, PhoneLoginResponse{SessionId: sessionId})
	}
}

func (h *AuthHandler) PhoneLoginVerify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[PhoneLoginVerifyRequest](w, r)
		if err != nil {
			res.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		email, err := h.Service.VerifyPhoneLogin(body.SessionId, body.Code)
		if err != nil {
			res.Json(w, http.StatusInternalServerError, err.Error())
			return
		}
		token, err := jwt.NewJWT(h.Config.Auth.Secret).CreateToken(jwt.JWTData{Email: email})
		if err != nil {
			res.Json(w, http.StatusInternalServerError, err.Error())
			return
		}

		res.Json(w, http.StatusOK, PhoneLoginVerifyResponse{Token: token})

	}
}
