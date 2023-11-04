package events_req

import (
	"context"
	"net/http"
	"sports-day/internal/entity"
	"sports-day/internal/errorhandler"
	"sports-day/internal/usecases/events"
	"sports-day/internal/utils"
	"strconv"

	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slog"
)

var eventsService = events.GetService()

func List(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer errorhandler.Recovery(w, r, http.StatusBadRequest)
		eventsList, err := eventsService.GetAllEvents(r.Context(), 10, 1)
		if err != nil {
			slog.Error("unable to fetch events", err)
			panic("unable to fetch events")
		}
		response := map[string]interface{}{
			"events": eventsList,
		}
		ctx := context.WithValue(r.Context(), "resData", response)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ListRegistered(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer errorhandler.Recovery(w, r, http.StatusBadRequest)
		userId, ok := r.Context().Value("userId").(entity.ID)
		if !ok {
			panic("invalid user id")
		}
		eventsList, err := eventsService.GetAllRegisteredEvents(r.Context(), 10, 1, userId)
		if err != nil {
			slog.Error("unable to fetch events", err)
			panic("unable to fetch events")
		}
		response := map[string]interface{}{
			"events": eventsList,
		}

		ctx := context.WithValue(r.Context(), "resData", response)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Register(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer errorhandler.Recovery(w, r, http.StatusBadRequest)
		userId, ok := r.Context().Value("userId").(entity.ID)
		if !ok {
			panic("invalid user id")
		}

		eventId, err := strconv.Atoi(chi.URLParam(r, "eventId"))
		if err != nil {
			panic(err.Error())
		}
		err = eventsService.RegisterEvent(r.Context(), userId, eventId)
		if err != nil {
			if utils.IsUniqueConstraintViolation(err) {
				panic("event already registered for this user")
			}
			panic(err.Error())
		}

		ctx := context.WithValue(r.Context(), "resData", "success")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UnRegister(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer errorhandler.Recovery(w, r, http.StatusBadRequest)
		userId, ok := r.Context().Value("userId").(entity.ID)
		if !ok {
			panic("invalid user id")
		}

		eventId, err := strconv.Atoi(chi.URLParam(r, "eventId"))
		if err != nil {
			panic(err.Error())
		}
		err = eventsService.UnRegisterEvent(r.Context(), userId, eventId)
		if err != nil {
			panic(err.Error())
		}

		ctx := context.WithValue(r.Context(), "resData", "success")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
