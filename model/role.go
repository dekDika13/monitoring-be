package model

import "time"

type Role struct {
	RoleID    uint      `gorm:"primaryKey;autoIncrement" json:"role_id"`
	Role_Name string    `gorm:"size:50;not null" json:"role_name"`
	CreatedAT time.Time `json:"created_at"`
	UpdatedAT time.Time `json:"updated_at"`
	DeletedAT time.Time `json:"deleted_at"`
}
