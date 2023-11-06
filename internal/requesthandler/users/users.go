package users_req

import (
	"context"
	"net/http"
	"sports-day/internal/entity"
	"sports-day/internal/errorhandler"
	"sports-day/internal/usecases/users"
)

var userService = users.GetUserService()

// GetUserDetails returns users role and name of logged in user
func GetUserDetails(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer errorhandler.Recovery(w, r, http.StatusBadRequest)
		userId, ok := r.Context().Value("userId").(entity.ID)

		if !ok {
			panic("invalid user id")
		}
		u, err := userService.GetUser(r.Context(), userId)
		if err != nil {
			panic(err.Error())
		}
		response := map[string]interface{}{
			"name": u.UserName,
			"role": u.Role,
		}
		ctx := context.WithValue(r.Context(), "resData", response)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
