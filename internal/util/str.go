package util

import "database/sql"

func ToNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{Valid: true, String: *s}
}

func FromNullString(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}
