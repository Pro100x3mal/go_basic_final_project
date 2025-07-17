package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

type AuthServiceInterface interface {
	Authenticate(p *models.Password) (string, error)
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

	fmt.Println(token.Token)

	writeJson(w, &token, http.StatusOK)
}
