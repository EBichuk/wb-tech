package repository

import (
	"calendar-service/internal/errs"
	"calendar-service/internal/models"
	"sync"
	"time"

	"github.com/google/uuid"
)

type mapStore struct {
	store map[string][]*models.Event
	mu    sync.Mutex
}

func New() *mapStore {
	return &mapStore{
		store: make(map[string][]*models.Event),
	}
}

func (m *mapStore) Full() {
	m.mu.Lock()
	defer m.mu.Unlock()

	eventDay := []*models.Event{
		{
			EventID:     "12",
			UserID:      "1",
			Date:        time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC),
			StartTime:   "10:00",
			EndTime:     "12:00",
			Title:       "sport",
			Description: "buy gym cards",
		},
		{
			EventID:     "13",
			UserID:      "1",
			Date:        time.Date(2026, 1, 4, 0, 0, 0, 0, time.UTC),
			StartTime:   "12:00",
			EndTime:     "14:00",
			Title:       "sport",
			Description: "run",
		},
	}

	m.store["1"] = eventDay
}

func (m *mapStore) CreateEvent(event *models.Event) (*models.Event, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	event.EventID = uuid.NewString()

	m.store[event.UserID] = append(m.store[event.UserID], event)

	return event, nil
}

func (m *mapStore) DeleteEvent(eventID, userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	listEvents, ok := m.store[userID]
	if !ok {
		return errs.ErrNotFoundUser
	}

	for i, e := range listEvents {
		if e.EventID == eventID {
			copy(listEvents[i:], listEvents[:i+1])
			listEvents[len(listEvents)-1] = nil
			return nil
		}
	}

	return errs.ErrNotFoundEvent
}

func (m *mapStore) UpdateEvent(event *models.Event) (*models.Event, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	listEvents, ok := m.store[event.UserID]
	if !ok {
		return &models.Event{}, errs.ErrNotFoundUser
	}

	for i, e := range listEvents {
		if e.EventID == event.EventID {
			listEvents[i] = event
			return event, nil
		}
	}

	return &models.Event{}, errs.ErrNotFoundEvent
}

func (m *mapStore) GetEventsForDay(userID string, date time.Time) ([]*models.Event, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	events, exists := m.store[userID]
	if !exists {
		return nil, errs.ErrNotFoundUser
	}

	eventDay := []*models.Event{}
	for _, e := range events {
		if e.Date.Equal(date) {
			eventDay = append(eventDay, e)
		}
	}

	if len(eventDay) < 1 {
		return eventDay, errs.ErrNotFoundEvent
	}

	return eventDay, nil
}

func (m *mapStore) GetEventsForWeek(userID string, date time.Time) ([]*models.Event, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	events, exists := m.store[userID]
	if !exists {
		return nil, errs.ErrNotFoundEvent
	}

	eventsForWeek := []*models.Event{}
	endDate := date.AddDate(0, 0, 7)

	for _, e := range events {
		if e.Date.Compare(date) > -1 && e.Date.Compare(endDate) < 1 {
			eventsForWeek = append(eventsForWeek, e)
		}
	}

	return eventsForWeek, nil
}

func (m *mapStore) GetEventsForMount(userID string, date time.Time) ([]*models.Event, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	events, exists := m.store[userID]
	if !exists {
		return nil, errs.ErrNotFoundEvent
	}

	eventsForWeek := []*models.Event{}

	dateMount, dateYear := date.Month(), date.Year()

	for _, e := range events {
		if e.Date.Month() == dateMount && e.Date.Year() == dateYear {
			eventsForWeek = append(eventsForWeek, e)
		}
	}

	return eventsForWeek, nil
}
