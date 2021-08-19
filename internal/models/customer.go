package models

import (
	"github.com/Freeline95/GoCrud/pkg/customTypes"
)

type Customer struct {
	Id uint `json:"id" db:"id"`
	FirstName string `json:"firstName" db:"first_name" validate:"required,max=100"`
	LastName string `json:"lastName" db:"last_name" validate:"required,max=100"`
	Gender string `json:"gender" db:"gender" validate:"required,oneof=Male Female"`
	Email string `json:"email" db:"email" validate:"required,email"`
	BirthDate customTypes.YmdTime `json:"birthDate" db:"birth_date" validate:"required,adult"`
	Address string `json:"address" db:"address" validate:"max=200"`
}