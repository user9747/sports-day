package middlewares

import (
	"context"
	"net/http"
	"sports-day/internal/errorhandler"
	"sports-day/internal/usecases/auth"
)

var authService = auth.GetAuthService()

const HeaderToken = "token"
const MessageInvalidToken = "Invalid Token"

// AuthFilter is used as authentication middleware for token API calls
func AuthFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer errorhandler.Recovery(w, r, http.StatusForbidden)
		token := r.Header.Get(HeaderToken)

		if len(token) > 0 {
			userInstance, err := authService.GetUserFromToken(r.Context(), token)
			if err != nil {
				errorhandler.CustomError(w, http.StatusUnauthorized, MessageInvalidToken)
				return
			}

			ctx := context.WithValue(r.Context(), "user", userInstance)
			ctx = context.WithValue(ctx, "userId", userInstance.ID)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			errorhandler.CustomError(w, http.StatusUnauthorized, MessageInvalidToken)
			return
		}
	})
}
