package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"uniqueIndex;not null"`
}
