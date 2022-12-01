package service

import (
    "context"

	"github.com/social-nyetwork/backend/internal/models"
)

func ChangeUserPassword(id int64, request* models.PasswordData) error {
    return nil
}

func CreateUser(request* models.CreateUserRequest) (int64, error) {
    return 0, nil
}

func GetUser(id int64) error {
    return nil
}

func GetUsers(page int, pageSize int) error {
    return nil
}

func EditUser(id int64, request* models.EditUserRequest) error {
    result ,err := dbpool.Exec(context.Background(),
                               "UPDATE \"user\" SET active = FALSE WHERE id = $1 AND active", id)
    if err != nil {
        return err
    }
    if result.RowsAffected() == 0 {
        return NotFoundError
    }
    return nil
}

func RemoveUser(id int64) error {
    result ,err := dbpool.Exec(context.Background(),
                               "UPDATE \"user\" SET active = FALSE WHERE id = $1 AND active", id)
    if err != nil {
        return err
    }
    if result.RowsAffected() == 0 {
        return NotFoundError
    }
    return nil
}
