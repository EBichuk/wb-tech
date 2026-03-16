package models

import (
	"calendar-service/internal/errs"
	"fmt"
	"time"
)

const timeFormat string = "15:04"

type Event struct {
	EventID     string    `json:"event_id"`
	UserID      string    `json:"user_id"`
	Date        time.Time `json:"date"`
	StartTime   string    `json:"start"`
	EndTime     string    `json:"end"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

func (e *Event) Validate() error {
	startTime, err := time.Parse(timeFormat, e.StartTime)
	if err != nil {
		return fmt.Errorf("invalid start time: %w", errs.ErrInvalidArgument)
	}
	endTime, err := time.Parse(timeFormat, e.EndTime)
	if err != nil {
		return fmt.Errorf("invalid end time: %w", errs.ErrInvalidArgument)
	}
	if startTime.After(endTime) {
		return errs.ErrInvalidPeriod
	}
	return nil
}
