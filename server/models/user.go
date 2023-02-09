package models

import "time"

type User struct {
	ID        int                   `json:"id"`
	Name      string                `json:"name" gorm:"type: varchar(255)"`
	Email     string                `json:"email" gorm:"type: varchar(255)"`
	Password  string                `json:"-" gorm:"type: varchar(255)"`
	Role      string                `json:"role" gorm:"type: varchar(255)"`
	Profile   ProfileResponse       `json:"profile"`
	Products  []ProductUserResponse `json:"products"`
	CreatedAt time.Time             `json:"-"`
	UpdatedAt time.Time             `json:"-"`
}

type UsersProfileResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
