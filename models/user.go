package models

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id         string      `gorm:"primaryKey" json:"id" validate:"required,numeric,len=16"` //jika uint64 hanya dapat menampung angka, shg hanya dapat mengimputkan angka 16
	Email      string      `json:"email" validate:"required,email"`                         //tidak blh ada spasi stlh koma
	Password   string      `json:"password"`
	KonfPass   string      `json:"konfPass"`
	Role       Role        `gorm:"default:user" json:"role" `
	Masyarakat *Masyarakat `gorm:"foreignKey:Id;references:NIK" json:"masyarakat"`
}

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "User"
)

func ValidateUser(user *User) error {
	validate := validator.New()
	return validate.Struct(user)
}

func (User) TableName() string {
	return "user"
}
