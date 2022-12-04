package service

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
    
	"github.com/social-nyetwork/backend/internal/models"
)

func ChangeUserPassword(id int64, request *models.PasswordData) error {
	if err := validate.Struct(request); err != nil {
		return err
	}
	
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

func CreateUser(request *models.CreateUserRequest) (id int64, err error) {
	id = -1
	
	err = validate.Struct(request)
	if err == nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

		if err == nil {
			generatedId := dbpool.QueryRow(context.Background(),
				"INSERT INTO \"user\" (login, password, full_name, birthday, gender, phone, email, about, country, country_region, currency, language, timezone) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id",
				request.Login, hashedPassword, request.FullName, request.Birthday, request.Gender,
				trimString(request.Phone), trimString(request.Email), trimString(request.About),
				request.Country, trimString(request.CountryRegion), request.Currency, request.Language, request.Timezone)

			err = generatedId.Scan(&id)
		}
	}


	return id, err
}

func GetUser(id int64) (user *models.GetUserResponse, err error) {
	result := dbpool.QueryRow(context.Background(),
		"SELECT login, full_name, birthday, gender, phone, email, about, country, country_region, currency, language, timezone FROM \"user\" WHERE id = $1 AND active", id)

    var login, name, gender, country, currency, language, timezone string
    var phone, email, about, countryRegion pgtype.Text
    var birthday time.Time
    
	err = result.Scan(&login, &name,
		&birthday, &gender,
		&phone, &email, &about,
		&country, &countryRegion, &currency, &language, &timezone)

	if err == pgx.ErrNoRows {
		err = NotFoundError
	} else if err == nil {
        user = &models.GetUserResponse{ ID: id, 
            UserData: &models.UserData { 
                ShortUserData: &models.ShortUserData { Login: login, FullName: name, Birthday: birthday.Format("2006-01-02"), Gender: models.Gender(gender) },
                Phone: unwrap(phone), Email: unwrap(email), About: unwrap(about), 
                Country: country, CountryRegion: unwrap(countryRegion), Currency: currency, Language: language, Timezone: timezone } }
    }

	return user, err
}

func GetUsers(page int, pageSize int) (response *models.GetUsersResponse, err error) {
    result, err := dbpool.Query(context.Background(), "SELECT \"id\", login, full_name, birthday, gender FROM \"user\" WHERE active LIMIT $1 OFFSET $2", pageSize, page * pageSize)
    if err == nil {
        response = &models.GetUsersResponse{ Items: []*models.UserItem{}, Page: &models.Page{ PageNumber: page, PageSize: pageSize, TotalItems: 0 } }
        for result.Next() {
            var id int64
            var login, name, gender string
            var birthday time.Time
            if innerErr := result.Scan(&id, &login, &name, &birthday, &gender); innerErr == nil {
                response.Items = append(response.Items, &models.UserItem{ ID: id, 
                    ShortUserData: &models.ShortUserData { Login: login, FullName: name, Birthday: birthday.Format("2006-01-02"), Gender: models.Gender(gender) } })
            }
        }
    }
    
	return response, err
}

func EditUser(id int64, request *models.EditUserRequest) error {
	if err := validate.Struct(request); err != nil {
		return err
	}
	
	result, err := dbpool.Exec(context.Background(),
		"UPDATE \"user\" SET full_name = $2, birthday = $3, gender = $4, phone = $5, email = $6, about = $7, country = $8, country_region = $9, currency = $10, language = $11, timezone = $12, last_modified = $13 WHERE id = $1 AND active",
		id, request.FullName, request.Birthday, request.Gender,
		trimString(request.Phone), trimString(request.Email), trimString(request.About),
		request.Country, trimString(request.CountryRegion), request.Currency, request.Language, request.Timezone,
		time.Now())
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return NotFoundError
	}
	return nil
}

func StartSession(login string, password string) (id int64, err error) {
	result := dbpool.QueryRow(context.Background(),
							  "SELECT id, password FROM \"user\" WHERE login = $1 AND active", login)
	var storedPassword []byte
	err = result.Scan(&id, &storedPassword)
	if err == pgx.ErrNoRows {
		err = NotFoundError
	} else if err == nil {
		err = bcrypt.CompareHashAndPassword(storedPassword, []byte(password))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			err = WrongCredentialsError
		}
	}
	return id, err
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
