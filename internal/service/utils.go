package service

import (
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
)

func trimString(value *string) *string {
	if value == nil {
		return nil
	}

	trimmedValue := strings.TrimSpace(*value)
	if trimmedValue == "" {
		return nil
	}
	return &trimmedValue
}

func unwrap(text pgtype.Text) *string {
	if text.Valid {
		return &text.String
	}
	return nil
}
