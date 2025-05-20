package model

import "time"

type Admin struct {
	Role          Role      `gorm:"foreignKey:RoleId"`
	AdminID       uint      `gorm:"primaryKey;autoIncrement" json:"admin_id"`
	Username      string    `gorm:"size:50;not null" json:"username"`
	Password      string    `gorm:"size:255;not null" json:"password"`
	Photo_profile string    `gorm:"size:255" json:"photo_profile"`
	Full_name     string    `gorm:"size:50;not null" json:"full_name"`
	Date_of_birth time.Time `json:"date_of_birth"`
	RoleId        uint      `json:"role_id"`
	CreatedAT     time.Time `json:"created_at"`
	UpdatedAT     time.Time `json:"updated_at"`
	DeletedAT     time.Time `json:"deleted_at"`
	
}
