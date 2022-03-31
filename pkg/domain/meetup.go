package domain

import "time"

// Meetup represents a UpMeet meetup.
type Meetup struct {
	ID             string         `json:"id" gorm:"primaryKey"`
	Name           string         `json:"name"`
	Description    string         `json:"description,omitempty"`
	InviteOnly     bool           `json:"invite_only" gorm:"default:false"`
	MinAge         int            `json:"min_age" gorm:"default:-1"`
	MeetupLocation MeetupLocation `json:"location,omitempty" gorm:"embedded;embeddedPrefix:location_"`
	OwnerID        string         `json:"owner_id"`
	Owner          User           `json:"-" gorm:"foreignKey:OwnerID"`
	Participants   []User         `gorm:"many2many:participants;"`
	CreatedAt      time.Time      `json:"created_at"`
}

// MeetupLocation represents a meetup location.
type MeetupLocation struct {
	Name         string `json:"name,omitempty"`
	Country      string `json:"country,omitempty"`
	State        string `json:"state,omitempty"`
	City         string `json:"city,omitempty"`
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
}

// CreateMeetupDTO represents a meetup creation data transfer object.
type CreateMeetupDTO struct {
	Name           string         `json:"name"`
	Description    string         `json:"description,omitempty"`
	InviteOnly     bool           `json:"invite_only"`
	MinAge         int            `json:"min_age"`
	MeetupLocation MeetupLocation `json:"location,omitempty"`
}

// UpdateMeetupDTO represents a meetup update data transfer object.
type UpdateMeetupDTO struct {
	Name           string         `json:"name"`
	Description    string         `json:"description,omitempty"`
	InviteOnly     bool           `json:"invite_only"`
	MinAge         int            `json:"min_age"`
	MeetupLocation MeetupLocation `json:"location,omitempty"`
}

type MeetupService interface {
	CreateMeetup(uid string, dto *CreateMeetupDTO) (*Meetup, error)
	GetMeetupByID(uid string, id string) (*Meetup, error)
	UpdateMeetup(uid string, dto *UpdateMeetupDTO) (*Meetup, error)
	DeleteMeetup(uid string, id string) error
}

type MeetupRepository interface {
	CreateMeetup(m *Meetup) error
	GetMeetupByID(id string) (*Meetup, error)
	UpdateMeetup(m *Meetup) error
	DeleteMeetup(id string) error
}
