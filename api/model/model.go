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

	Gifts struct {
		ID          int32   `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       string  `json:"price"`
		ImageURL    string  `json:"image_url"`
		Category    string  `json:"category"`
		StoreLink   string  `json:"store_link"`
		IsBooked    bool    `json:"is_booked"`
		CreatedAt   string  `json:"created_at"`
		UpdatedAt   *string `json:"updated_at"`
	}
)
