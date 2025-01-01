package models

import "time"

type URLData struct {
	OriginalURL string
	Expiry      time.Time
	AccessCount int
}
