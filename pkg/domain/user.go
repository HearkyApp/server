package domain

import (
	"time"
)

// User is a user of the application.
type User struct {
	ID               string    `json:"id" gorm:"primaryKey"`
	Username         string    `json:"username" gorm:"uniqueIndex"`
	Name             string    `json:"name" gorm:"index"`
	ProfilePicture   string    `json:"profile_picture"`
	Age              int       `json:"age" gorm:"default:-1"`
	AgeVerified      bool      `json:"age_verified"`
	AgePrivate       bool      `json:"age_private"`
	Bio              string    `json:"bio"`
	InstagramProfile string    `json:"instagram_profile"`
	FacebookProfile  string    `json:"facebook_profile"`
	TwitterProfile   string    `json:"twitter_profile"`
	DiscordTag       string    `json:"discord_tag"`
	CreatedAt        time.Time `json:"created_at"`
}

const (
	// UsernameMinLength is the minimum length of a username.
	UsernameMinLength = 3
	// UsernameMaxLength is the maximum length of a username.
	UsernameMaxLength = 24
	// UserNameMaxLength is the maximum length of a users' name.
	UserNameMaxLength = 64
	// UserBioMaxLength is the maximum length of a users' bio.
	UserBioMaxLength = 256
	// UserMaxAge is the maximum age of a user. funfact: (03/29/2022 - current oldest person is Kane Tananka at age 119)
	UserMaxAge = 120
)

// CreateUserDTO is the data transfer object for creating a user.
type CreateUserDTO struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

// UpdateUserDTO is the data transfer object for updating a user.
type UpdateUserDTO struct {
	Username         string `json:"username,omitempty"`
	Name             string `json:"name,omitempty"`
	Bio              string `json:"bio,omitempty"`
	Age              int    `json:"age,omitempty"`
	AgePrivate       bool   `json:"age_private,omitempty"`
	InstagramProfile string `json:"instagram_profile,omitempty"`
	FacebookProfile  string `json:"facebook_profile,omitempty"`
	TwitterProfile   string `json:"twitter_profile,omitempty"`
	DiscordTag       string `json:"discord_tag,omitempty"`
}

type UserService interface {
	GetUserByID(id string) (*User, error)
	CreateUser(id string, dto *CreateUserDTO) (*User, error)
	UpdateUser(id string, dto *UpdateUserDTO) (*User, error)
}

type UserRepository interface {
	CreateUser(u *User) error
	GetUserByID(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	SearchUsersByName(name string) ([]*User, error)
	SearchUsersByUsername(username string) ([]*User, error)
	UpdateUser(u *User) error
	DeleteUser(id string) error
}
