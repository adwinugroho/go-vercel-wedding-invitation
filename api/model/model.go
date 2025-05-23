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
		IsAttend    bool   `json:"is_attend"`
		NumberGuest int32  `json:"number_guest"`
		CreatedAt   string `json:"created_at"`
	}
)
