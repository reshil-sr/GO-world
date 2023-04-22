package mock

import (
	"database/sql"
	"errors"
	"testing"
)

func TestOpenDB(t *testing.T) {
	mockError := errors.New("Its an error")
	subTests := []struct {
		name        string
		u, p, a, db string
		sqlOpener   func(string, string) (*sql.DB, error)
		expectedErr error
	}{
		{
			name: "happy-path",
			u:    "u",
			p:    "p",
			a:    "a",
			db:   "db",
			sqlOpener: func(driver string, source string) (*sql.DB, error) {
				if source != "u:p@a/db" {
					return nil, errors.New("wrong connection String!")
				}
				return nil, nil
			},
		},
		{
			name: "error-from-sqlOpener",
			sqlOpener: func(driver string, source string) (*sql.DB, error) {
				return nil, mockError
			},
			expectedErr: mockError,
		},
	}

	for _, subTest := range subTests {
		//approach 2
		//SQLOpen = subTest.sqlOpener
		t.Run(subTest.name, func(t *testing.T) {
			_, err := OpenDB(subTest.u, subTest.p, subTest.a, subTest.db, subTest.sqlOpener)
			if !errors.Is(err, subTest.expectedErr) {
				t.Errorf("expected error (%v), got error (%v)", subTest.expectedErr, err)
			}
		})
	}
}
