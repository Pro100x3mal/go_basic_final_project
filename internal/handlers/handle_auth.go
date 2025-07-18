package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

type AuthServiceInterface interface {
	Authenticate(p *models.Password) (string, error)
	ValidateToken(tokenStr string) (bool, error)
}
type AuthHandler struct {
	auth AuthServiceInterface
}

func NewAuthHandler(auth AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		auth: auth,
	}
}

func (ah *AuthHandler) handleAuth(w http.ResponseWriter, r *http.Request) {
	var p models.Password

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		writeJson(w, &models.RespError{Error: "invalid request body"}, http.StatusBadRequest)
		return
	}

	var token models.RespToken
	token.Token, err = ah.auth.Authenticate(&p)
	if err != nil {
		writeJson(w, &models.RespError{Error: err.Error()}, http.StatusUnauthorized)
		return
	}

	writeJson(w, &token, http.StatusOK)
}

func (ah *AuthHandler) Validate(tokenStr string) (bool, error) {
	if len(tokenStr) == 0 {
		return false, errors.New("invalid token")
	}

	return ah.auth.ValidateToken(tokenStr)
}
