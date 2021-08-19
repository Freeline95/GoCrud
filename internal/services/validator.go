package services

import (
	"github.com/Freeline95/GoCrud/pkg/customTypes"
	"github.com/go-playground/validator/v10"
	"log"
	"reflect"
	"time"
)

/*func NewValidator() *validator.Validate {
	myValidator := validator.New()
	myValidator.RegisterValidation("adult", isAdult)

	return myValidator
}

func isAdult(fl validator.FieldLevel) bool {
	log.Println(123)
	birthTime := fl.Field().Interface().(customTypes.YmdTime)

	log.Println(birthTime)

	now := time.Now()
	birthTimeAdd18 := birthTime.AddDate(18, 0, 0)
	birthTimeAdd60 := birthTime.AddDate(60, 0, 0)

	return birthTimeAdd18.Before(now) && now.Before(birthTimeAdd60)
}*/

func NewValidator() *validator.Validate {
	myValidator := validator.New()
	myValidator.RegisterCustomTypeFunc(ymdTimeTypeValue, customTypes.YmdTime{})
	myValidator.RegisterValidation("adult", isAdult)

	return myValidator
}

func ymdTimeTypeValue(v reflect.Value) interface{} {
	return v.Interface().(customTypes.YmdTime).Time
}

func isAdult(fl validator.FieldLevel) bool {
	birthTime := fl.Field().Interface().(time.Time)

	log.Println(birthTime)

	now := time.Now()
	birthTimeAdd18 := birthTime.AddDate(18, 0, 0)
	birthTimeAdd60 := birthTime.AddDate(60, 0, 0)

	return birthTimeAdd18.Before(now) && now.Before(birthTimeAdd60)
}