package event

type CreateEventRequest struct {
	UserID      string `json:"user_id"`
	Date        string `json:"date"`
	StartTime   string `json:"start"`
	EndTime     string `json:"end"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdaleEventRequest struct {
	EventID     string `json:"event_id"`
	UserID      string `json:"user_id"`
	Date        string `json:"date"`
	StartTime   string `json:"start"`
	EndTime     string `json:"end"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type DeleteEventRequest struct {
	EventID string `json:"event_id"`
	UserID  string `json:"user_id"`
}
