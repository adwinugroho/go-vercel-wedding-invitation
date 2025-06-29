package model

type (
	RequestNewWish struct {
		Name        string `json:"name"`
		Message     string `json:"message"`
		IsPublished bool   `json:"is_published"`
	}

	RequestNewRSVP struct {
		Name        string `json:"name"`
		GuestCount  int32  `json:"guest_count"`
		IsAttending *bool  `json:"is_attending"`
	}

	RequestNewGift struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Price       string `json:"price"`
		ImageURL    string `json:"image_url"`
		Category    string `json:"category"`
		StoreLink   string `json:"store_link"`
	}

	RequestMarkedGiftAsBooked struct {
		ID       int32 `json:"id"`
		IsBooked *bool `json:"is_booked"`
	}
)
