package models

import "time"

// StatusPage used for statuspage.io response
type StatusPage struct {
	Page   PageDetails   `json:"page"`
	Status StatusDetails `json:"status"`
}

// PageDetails contains statuspage details
type PageDetails struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	TimeZone  string    `json:"time_zone"`
	UpdatedAt time.Time `json:"updated_at"`
}

// StatusDetails contains statuspage status details
type StatusDetails struct {
	Indicator   string `json:"indicator"`
	Description string `json:"description"`
}
