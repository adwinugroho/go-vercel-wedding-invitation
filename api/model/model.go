package model

type (
	Wishes struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Message     string `json:"message"`
		IsPublished bool   `json:"is_published"`
		CreatedAt   string `json:"created_at"`
	}

	Reservation struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		IsAttending bool   `json:"is_attending"`
		GuestCount  int32  `json:"guest_count"`
		CreatedAt   string `json:"created_at"`
	}
)
