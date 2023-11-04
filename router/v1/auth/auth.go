package auth_router

import (
	"sports-day/internal/middlewares"
	auth_req "sports-day/internal/requesthandler/auth"
	"sports-day/internal/responsehandler"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.With(auth_req.Login).Post("/login", responsehandler.GenericRes)
	r.With(middlewares.AuthFilter, auth_req.Logout).Post("/logout", responsehandler.GenericRes)
}
