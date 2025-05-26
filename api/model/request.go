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
)
