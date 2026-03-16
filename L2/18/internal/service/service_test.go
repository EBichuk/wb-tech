package service

import (
	"testing"
	"time"

	"calendar-service/internal/errs"
	"calendar-service/internal/models"
	repoMock "calendar-service/internal/repository/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateEvent(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := repoMock.NewMockrepository(ctl)

	userID := "1"
	date := time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC)
	start := "10:00"
	end := "14:00"
	title := "Gym"
	desc := "Don't forget a water"

	eventInput := models.Event{
		UserID: userID, Date: date, StartTime: start,
		EndTime: end, Title: title, Description: desc,
	}

	eventID := uuid.NewString()

	evExp := models.Event{
		EventID: eventID,
		UserID:  userID, Date: date, StartTime: start,
		EndTime: end, Title: title, Description: desc,
	}

	repo.EXPECT().CreateEvent(gomock.Any()).Return(&evExp, nil)
	service := New(repo)

	event, err := service.CreateEvent(&eventInput)

	require.NoError(t, err)
	require.Equal(t, evExp, *event)
}

func TestCreateEventPeriodError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := repoMock.NewMockrepository(ctl)

	event := &models.Event{
		UserID:      "1",
		Date:        time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC),
		StartTime:   "15:00",
		EndTime:     "10:00",
		Title:       "sport",
		Description: "buy gym cards",
	}

	service := New(repo)

	event, err := service.CreateEvent(event)

	require.ErrorIs(t, err, errs.ErrInvalidPeriod)
	require.Nil(t, event)
}

func TestCreateEventValidationError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := repoMock.NewMockrepository(ctl)

	event := &models.Event{
		UserID:      "1",
		Date:        time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC),
		StartTime:   "10:00",
		EndTime:     "25:00",
		Title:       "test",
		Description: "buy gym cards",
	}

	service := New(repo)

	event, err := service.CreateEvent(event)

	require.ErrorIs(t, err, errs.ErrInvalidArgument)
	require.Nil(t, event)
}

func TestGetEventsForDay(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := repoMock.NewMockrepository(ctl)

	userID := "1"
	date := time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC)
	eventDay := []*models.Event{
		{
			EventID:     "12",
			UserID:      "1",
			Date:        date,
			StartTime:   "13:00",
			EndTime:     "14:00",
			Title:       "sport",
			Description: "buy gym cards",
		},
		{
			EventID:     "13",
			UserID:      "1",
			Date:        date,
			StartTime:   "11:00",
			EndTime:     "11:30",
			Title:       "sport",
			Description: "run",
		},
	}

	service := New(repo)

	repo.EXPECT().GetEventsForDay(userID, date).Return(eventDay, nil)
	event, err := service.GetEventsForDay(userID, date)

	require.NoError(t, err)
	require.Equal(t, eventDay, event)
}

func TestGetEventsForDayError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := repoMock.NewMockrepository(ctl)

	userID := "1"
	date := time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC)

	service := New(repo)

	repo.EXPECT().GetEventsForDay(userID, date).Return(nil, errs.ErrNotFoundEvent)
	event, err := service.GetEventsForDay(userID, date)

	require.ErrorIs(t, err, errs.ErrNotFoundEvent)
	require.Nil(t, event)
}

func TestGetEventsForWeek(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := repoMock.NewMockrepository(ctl)

	userID := "1"
	date := time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC)
	eventDay := []*models.Event{
		{
			EventID:     "12",
			UserID:      "1",
			Date:        date,
			StartTime:   "13:00",
			EndTime:     "14:00",
			Title:       "sport",
			Description: "buy gym cards",
		},
		{
			EventID:     "13",
			UserID:      "1",
			Date:        date,
			StartTime:   "11:00",
			EndTime:     "11:30",
			Title:       "sport",
			Description: "run",
		},
	}

	service := New(repo)

	repo.EXPECT().GetEventsForWeek(userID, date).Return(eventDay, nil)
	event, err := service.GetEventsForWeek(userID, date)

	require.NoError(t, err)
	require.Equal(t, eventDay, event)
}

func TestGetEventsForWeekError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := repoMock.NewMockrepository(ctl)

	userID := "1"
	date := time.Now()

	service := New(repo)

	repo.EXPECT().GetEventsForWeek(gomock.Any(), date).Return(nil, errs.ErrNotFoundEvent)
	event, err := service.GetEventsForWeek(userID, date)

	require.ErrorIs(t, err, errs.ErrNotFoundEvent)
	require.Nil(t, event)
}

func TestGetEventsForMountError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := repoMock.NewMockrepository(ctl)

	userID := "1"
	date := time.Now()

	service := New(repo)

	repo.EXPECT().GetEventsForMount(gomock.Any(), date).Return(nil, errs.ErrNotFoundEvent)
	event, err := service.GetEventsForMount(userID, date)

	require.ErrorIs(t, err, errs.ErrNotFoundEvent)
	require.Nil(t, event)
}
