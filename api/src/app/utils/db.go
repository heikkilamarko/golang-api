package utils

import "database/sql"

// GetNullStringValue func
func GetNullStringValue(value sql.NullString) string {
	if value.Valid {
		return value.String
	}
	return ""
}
