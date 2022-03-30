package domain

import "time"

type Meetup struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	MinAge      int       `json:"min_age"`
	CreatedAt   time.Time `json:"created_at"`
}
