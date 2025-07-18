package middlewares

import (
	"net/http"

	"github.com/Pro100x3mal/go_basic_final_project/internal/config"
)

type TokenValidator interface {
	Validate(token string) (bool, error)
}

func NewAuthMiddleware(tv TokenValidator, cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if len(cfg.Password) > 0 {

				cookie, err := r.Cookie("token")
				if err != nil {
					http.Error(w, "token not found", http.StatusUnauthorized)
					return
				}

				valid, err := tv.Validate(cookie.Value)
				if err != nil || !valid {
					http.Error(w, "invalid token ", http.StatusUnauthorized)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
