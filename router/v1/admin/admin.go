// admin_router has admin routes like create user and events...etc
package admin_router

import (
	"sports-day/internal/middlewares"
	admin_req "sports-day/internal/requesthandler/admin"
	"sports-day/internal/responsehandler"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.With(middlewares.AuthFilter, admin_req.CreateUser).Post("/create-user", responsehandler.GenericRes)
}
