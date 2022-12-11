package service

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("iso-date", validator.Func(func(fl validator.FieldLevel) bool {
		if value, ok := fl.Field().Interface().(string); ok {
			_, err := time.Parse("2006-01-02", value)
			return err == nil
		}
		return false
	}))
	validate.RegisterValidation("present", validator.Func(func(fl validator.FieldLevel) bool {
		switch val := fl.Field().Interface().(type) {
		case string:
			{
				valTime, err := time.Parse("2006-01-02", val)
				return err == nil && !valTime.After(time.Now())
			}
		case time.Time:
			return !val.After(time.Now())
		default:
			return false
		}
	}))
}
