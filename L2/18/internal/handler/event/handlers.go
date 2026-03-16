package event

import (
	"calendar-service/internal/errs"
	"calendar-service/internal/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

type Service interface {
	CreateEvent(event *models.Event) (*models.Event, error)
	DeleteEvent(eventID, userID string) error
	UpdateEvent(event *models.Event) (*models.Event, error)
	GetEventsForDay(userID string, date time.Time) ([]*models.Event, error)
	GetEventsForWeek(userID string, date time.Time) ([]*models.Event, error)
	GetEventsForMount(userID string, date time.Time) ([]*models.Event, error)
}

type handler struct {
	service Service
}

func New(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var req CreateEventRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "error to decoding request body", http.StatusBadRequest)
		return
	}

	date, err := ParseDate(req.Date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	event := &models.Event{
		UserID:      req.UserID,
		Date:        date,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Title:       req.Title,
		Description: req.Description,
	}

	eventNew, err := h.service.CreateEvent(event)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrInvalidArgument), errors.Is(err, errs.ErrInvalidPeriod):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(eventNew); err != nil {
		log.Printf("error to encoding reply body: %v", err)
	}
}

func (h *handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	var req DeleteEventRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "error to decoding request body", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteEvent(req.EventID, req.UserID)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrNotFoundEvent):
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var req UpdaleEventRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "error to decoding request body", http.StatusBadRequest)
		return
	}

	date, err := ParseDate(req.Date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	event := &models.Event{
		EventID:     req.EventID,
		UserID:      req.UserID,
		Date:        date,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Title:       req.Title,
		Description: req.Description,
	}

	eventNew, err := h.service.UpdateEvent(event)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrInvalidArgument), errors.Is(err, errs.ErrInvalidPeriod):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, errs.ErrNotFoundEvent):
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(eventNew); err != nil {
		log.Printf("error to encoding reply body: %v", err)
	}
}

func (h *handler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")

	dateStr := r.URL.Query().Get("date")
	date, err := ParseDate(dateStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	eventNew, err := h.service.GetEventsForDay(userID, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(eventNew); err != nil {
		log.Printf("error to encoding reply body: %v", err)
	}
}

func (h *handler) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")

	dateStr := r.URL.Query().Get("date")
	date, err := ParseDate(dateStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	eventNew, err := h.service.GetEventsForWeek(userID, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(eventNew); err != nil {
		log.Printf("error to encoding reply body: %v", err)
	}
}

func (h *handler) GetEventsForMount(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")

	dateStr := r.URL.Query().Get("date")
	date, err := ParseDate(dateStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	eventNew, err := h.service.GetEventsForMount(userID, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(eventNew); err != nil {
		log.Printf("error to encoding reply body: %v", err)
	}
}

func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}
