package utils

import (
	"database/sql"
	"testing"
)

func TestGetNullStringValue(t *testing.T) {
	nullString := sql.NullString{Valid: false, String: ""}

	actual := GetNullStringValue(nullString)
	expected := ""

	if actual != expected {
		t.Errorf("actual: %s, expected: %s", actual, expected)
	}

	nullString = sql.NullString{Valid: true, String: "test"}

	actual = GetNullStringValue(nullString)
	expected = "test"

	if actual != expected {
		t.Errorf("actual: %s, expected: %s", actual, expected)
	}
}
