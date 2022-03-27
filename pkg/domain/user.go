package domain

import (
	"time"
)

type User struct {
	ID               string    `json:"id" gorm:"primaryKey"`
	Email            string    `json:"email" gorm:"uniqueIndex"`
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
	Password         string    `json:"-"`
	Verified         bool      `json:"verified"`
	Disabled         bool      `json:"disabled"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type SignUpDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	Email            string `json:"email"`
	Username         string `json:"username"`
	Name             string `json:"name"`
	Bio              string `json:"bio"`
	Age              bool   `json:"age"`
	Password         string `json:"password"`
	InstagramProfile string `json:"instagram_profile"`
	FacebookProfile  string `json:"facebook_profile"`
	TwitterProfile   string `json:"twitter_profile"`
	DiscordTag       string `json:"discord_tag"`
}

type UserService interface {
	SignUp(dto *SignUpDTO) (*User, error)
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
