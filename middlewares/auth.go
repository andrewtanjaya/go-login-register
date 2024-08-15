package middlewares

import (
	"context"
	"go-login-register/responses"
	"go-login-register/utils"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")

		if accessToken == "" {
			responses.Response(w, 401, "Unauthorized", nil)
			return
		}

		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(accessToken, bearerPrefix) {
			responses.Response(w, 401, "Unauthorized", nil)
			return
		}

		accessToken = strings.TrimPrefix(accessToken, bearerPrefix)

		user, err := utils.ValidateToken(accessToken)
		if err != nil {
			responses.Response(w, 401, err.Error(), nil)
			return
		}

		ctx := context.WithValue(r.Context(), "current_user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
