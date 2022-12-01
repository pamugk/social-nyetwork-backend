package service

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/social-nyetwork/backend/internal/models"
)

func ChangeUserPassword(id int64, request* models.PasswordData) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
    
    if err != nil {
        return err
    }
    
    result, err := dbpool.Exec(context.Background(), "UPDATE \"user\" SET password = $2 WHERE id = $1 AND active", id, hashedPassword)
    
    if err != nil {
        return err
    }
    if result.RowsAffected() == 0 {
        return NotFoundError
    }
    
    return nil
}

func CreateUser(request* models.CreateUserRequest) (int64, error) {
    id int64 := -1
    
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
    
    if err != nil {
        return id, err
    }
    
    generatedId := dbpool.QueryRow(context.Background(), "INSERT INTO \"user\" (login, password, full_name, birthday, gender, phone, email, about, country, country_region, currency, language, timezone) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id", request.Login, hashedPassword, request.FullName, request.Birthday, request.Gender, trimString(request.Phone), trimString(request.Email), trimString(request.About), request.Country, trimString(request.CountryRegion), request.Currency, request.Language, request.Timezone)
    
    err = generatedId.Scan(&id)
    
    return id, err
}

func GetUser(id int64) (*models.GetUserResponse, error) {
    result := dbpool.QueryRow("SELECT id, login, full_name, birthday, gender, phone, email, about, country, country_region, currency, language, timezone FROM \"user\" WHERE id = $1 AND active", id)
    
    return nil, nil
}

func GetUsers(page int, pageSize int) (*models.GetUsersResponse, error) {
    return nil, nil
}

func EditUser(id int64, request* models.EditUserRequest) error {
    result, err := dbpool.Exec(context.Background(),
                               "UPDATE \"user\" SET full_name = $2, birthday = $3, gender = $4, phone = $5, email = $6, about = $7, country = $8, country_region = $9, currency = $10, language = $11, timezone = $12 WHERE id = $1 AND active", id, request.FullName, request.Birthday, request.Gender, trimString(request.Phone), trimString(request.Email), trimString(request.About), request.Country, trimString(request.CountryRegion), request.Currency, request.Language, request.Timezone)
    if err != nil {
        return err
    }
    if result.RowsAffected() == 0 {
        return NotFoundError
    }
    return nil
}

func RemoveUser(id int64) error {
    result, err := dbpool.Exec(context.Background(),
                               "UPDATE \"user\" SET active = FALSE WHERE id = $1 AND active", id)
    if err != nil {
        return err
    }
    if result.RowsAffected() == 0 {
        return NotFoundError
    }
    return nil
}
