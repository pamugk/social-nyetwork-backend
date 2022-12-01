package service

import (
    "strings"
)

func trimString(value* string) *string {
    if value == nil {
        return nil
    }
    
    trimmedValue := strings.TrimSpace(*value)
    if trimmedValue == "" {
        return nil
    }
    return &trimmedValue
}
