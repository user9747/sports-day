package v1

import (
	admin_router "sports-day/router/v1/admin"
	auth_router "sports-day/router/v1/auth"
	events_router "sports-day/router/v1/events"
	user_router "sports-day/router/v1/users"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.Route("/auth", auth_router.Routes)
	r.Route("/admin", admin_router.Routes)
	r.Route("/users", user_router.Routes)
	r.Route("/events", events_router.Routes)
}
