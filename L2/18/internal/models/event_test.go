package models_test

import (
	"calendar-service/internal/errs"
	"calendar-service/internal/models"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestValidate_Success(t *testing.T) {
	tests := []struct {
		name  string
		event models.Event
		want  error
	}{
		{
			name: "success",
			event: models.Event{
				"1", "1", time.Date(2026, 01, 01, 0, 0, 0, 0, time.UTC), "10:00", "13:00", "test title", ""},
			want: nil,
		},
		{
			name: "borderline: start time == end time",
			event: models.Event{
				"1", "1", time.Date(2026, 01, 01, 0, 0, 0, 0, time.UTC), "10:00", "10:00", "test title", ""},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.event.Validate()
			require.NoError(t, err)
		})
	}
}

func TestValidate_PeriodError(t *testing.T) {
	tests := []struct {
		name  string
		event models.Event
		want  error
	}{
		{
			name: "startTime after endTime",
			event: models.Event{
				"1", "1", time.Date(2026, 01, 01, 0, 0, 0, 0, time.UTC), "13:00", "10:00", "test title", ""},
			want: errs.ErrInvalidPeriod,
		},
		{
			name: "startTime invalid format",
			event: models.Event{
				"1", "1", time.Date(2026, 01, 01, 0, 0, 0, 0, time.UTC), "10:00:01", "19:00", "test title", ""},
			want: errs.ErrInvalidArgument,
		},
		{
			name: "endTime invalid format",
			event: models.Event{
				"1", "1", time.Date(2026, 01, 01, 0, 0, 0, 0, time.UTC), "10:00", "190:00", "test title", ""},
			want: errs.ErrInvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.event.Validate()
			require.ErrorIs(t, err, tt.want)
		})
	}
}
