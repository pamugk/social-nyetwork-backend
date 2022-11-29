package models

import (
	"net/http"
	"time"
)

type CreateUserRequest struct {
	*PasswordData
	*UserData
}

func (*CreateUserRequest) Bind(r *http.Request) error {
	return nil
}

type EditUserRequest struct {
	*UserData
}

func (*EditUserRequest) Bind(r *http.Request) error {
	return nil
}

type GetUserResponse struct {
	ID int64 `json:"id"` // User ID
	*UserData
}

type GetUsersResponse struct {
	Items []*UserItem // Found items
	*Page
}

type Gender string

const (
	Undefined Gender = ""
	Male      Gender = "MALE"
	Female    Gender = "FEMALE"
)

// @Description Short user info
type ShortUserData struct {
	Login    string `example:"user" json:"login" validate:"required,min=3,max=255,alphanum"` // User login
	FullName string `example:"User Usersson" json:"name" validate:"required,max=1500"`                // User full name

	Birthday time.Time `example:"2022-01-01" json:"birthday" validate:"required"` // User birthday
	Gender   Gender    `example:"MALE" json:"gender" validate:"required"`         // User gender
}

// @Description Password definition
type PasswordData struct {
	Password string `example:"password" json:"password" validate:"required,min=5"` // User password
}

func (*PasswordData) Bind(r *http.Request) error {
	return nil
}

// @Description User full info
type UserData struct {
	*ShortUserData

	Phone string `example:"+78005553535" json:"phone" validate:"e164"`                 // User main phone number
	Email string `example:"example@example.com" json:"email" validate:"email,max=320"` // User main e-mail

	About         string `example:"Some useful and interesting info" json:"about" validate:"max=1000"` // Short user self-description
	Country       string `example:"RU" json:"country" validate:"required,iso3166_1_alpha2"`            // User country code
	CountryRegion string `example:"RU-PER" json:"country" validate:"iso3166_2"`                        // User country region code
	Currency      string `example:"RUB" json:"currency" validate:"required,iso4217"`                   // User preferred currency code
	Language      string `example:"ru_RU" json:"language" validate:"required,bcp47_language_tag"`      // User preferred language code
	Timezone      string `example:"UTC" json:"timezone" validate:"required,timezone"`                  // User timezone
}

// @Description User short info item
type UserItem struct {
	ID int64 `json:"id"` // User ID
	*ShortUserData
}
