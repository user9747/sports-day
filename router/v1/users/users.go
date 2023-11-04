package user_router

import (
	"sports-day/internal/middlewares"
	users_req "sports-day/internal/requesthandler/users"
	"sports-day/internal/responsehandler"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.With(middlewares.AuthFilter, users_req.GetUserDetails).Get("/{userId}", responsehandler.GenericRes)
}
