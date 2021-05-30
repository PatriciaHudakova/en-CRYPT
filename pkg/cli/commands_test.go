package cli_test

import (
	app "en-CRYPT/pkg/cli"
	"en-CRYPT/pkg/mocks"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

func TestGetAll(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{
			name: "should successfully retrieve db data",
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := mocks.NewDB()

			// escape special characters in query
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM credentials;")).
				WillReturnRows(sqlmock.NewRows([]string{"domain", "username", "password", "created_on", "user_id"}).
					AddRow("facebook", "myusername", "mypassword", time.Now(), 1))

			got := app.GetAll(db)

			assert.NotNil(t, got)
		})
	}
}
