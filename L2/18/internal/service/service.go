package service

import (
	"calendar-service/internal/models"
	"time"
)

type Repository interface {
	DeleteEvent(eventID, userID string) error
	CreateEvent(event *models.Event) (*models.Event, error)
	UpdateEvent(event *models.Event) (*models.Event, error)
	GetEventsForDay(userID string, date time.Time) ([]*models.Event, error)
	GetEventsForWeek(userID string, date time.Time) ([]*models.Event, error)
	GetEventsForMount(userID string, date time.Time) ([]*models.Event, error)
}

type service struct {
	repo Repository
}

func New(r Repository) *service {
	return &service{
		repo: r,
	}
}

func (s *service) CreateEvent(event *models.Event) (*models.Event, error) {
	err := event.Validate()
	if err != nil {
		return nil, err
	}

	return s.repo.CreateEvent(event)
}

func (s *service) DeleteEvent(eventID, userID string) error {
	err := s.repo.DeleteEvent(eventID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateEvent(event *models.Event) (*models.Event, error) {
	err := event.Validate()
	if err != nil {
		return nil, err
	}

	return s.repo.UpdateEvent(event)
}

func (s *service) GetEventsForDay(userID string, date time.Time) ([]*models.Event, error) {
	events, err := s.repo.GetEventsForDay(userID, date)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (s *service) GetEventsForWeek(userID string, date time.Time) ([]*models.Event, error) {
	events, err := s.repo.GetEventsForWeek(userID, date)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (s *service) GetEventsForMount(userID string, date time.Time) ([]*models.Event, error) {
	events, err := s.repo.GetEventsForMount(userID, date)
	if err != nil {
		return nil, err
	}

	return events, nil
}
