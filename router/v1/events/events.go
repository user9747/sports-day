package events_router

import (
	"sports-day/internal/middlewares"
	events_req "sports-day/internal/requesthandler/events"
	"sports-day/internal/responsehandler"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.With(middlewares.AuthFilter, events_req.Register).Post("/register/{eventId}", responsehandler.GenericRes)
	r.With(middlewares.AuthFilter, events_req.UnRegister).Post("/unregister/{eventId}", responsehandler.GenericRes)
	r.With(middlewares.AuthFilter, events_req.List).Get("/list", responsehandler.GenericRes)
	r.With(middlewares.AuthFilter, events_req.ListRegistered).Get("/list-registered", responsehandler.GenericRes)
}
