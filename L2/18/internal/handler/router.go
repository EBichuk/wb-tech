package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler interface {
	DeleteEvent(w http.ResponseWriter, r *http.Request)
	CreateEvent(w http.ResponseWriter, r *http.Request)
	UpdateEvent(w http.ResponseWriter, r *http.Request)
	GetEventsForDay(w http.ResponseWriter, r *http.Request)
	GetEventsForWeek(w http.ResponseWriter, r *http.Request)
	GetEventsForMount(w http.ResponseWriter, r *http.Request)
}

func RegisterRouter(h Handler, logger *slog.Logger) http.Handler {
	handler := chi.NewRouter()

	handler.Use(middleware.RequestID)
	handler.Use(MiddlewareLogging(logger))

	handler.Post("/delete_event", h.DeleteEvent)
	handler.Post("/create_event", h.CreateEvent)
	handler.Post("/update_event", h.UpdateEvent)
	handler.Get("/events_for_day", h.GetEventsForDay)
	handler.Get("/events_for_week", h.GetEventsForWeek)
	handler.Get("/events_for_mount", h.GetEventsForMount)

	return handler
}
