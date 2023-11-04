package users_req

import (
	"context"
	"net/http"
	"sports-day/internal/entity"
	"sports-day/internal/errorhandler"
	"sports-day/internal/usecases/users"

	"github.com/go-chi/chi/v5"
)

var userService = users.GetUserService()

func GetUserDetails(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer errorhandler.Recovery(w, r, http.StatusBadRequest)
		id := chi.URLParam(r, "userId")
		userId, err := entity.StringToID(id)

		if err != nil {
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
