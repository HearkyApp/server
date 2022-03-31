package domain

import "time"

type ParticipantPermissions struct {
	UserID     string    `json:"user_id" gorm:"primaryKey"`
	MeetupID   string    `json:"meetup_id" gorm:"primaryKey"`
	Permission string    `json:"permission" gorm:"primaryKey"`
	CreatedAt  time.Time `json:"created_at"`
}
