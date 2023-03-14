package postgres

import (
	"context"
	"github.com/Mldlr/storety/internal/server/models"
	"github.com/Mldlr/storety/internal/server/storage"
	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v2"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestDB_CreateUser(t *testing.T) {
	id, err := uuid.NewRandom()
	assert.NoError(t, err)

	tests := []struct {
		name    string
		rows    *pgxmock.Rows
		wantErr error
	}{
		{
			name:    "Create user successfully",
			rows:    pgxmock.NewRows([]string{"id"}).AddRow(id.String()),
			wantErr: nil,
		},
		{
			name:    "Try to create user with duplicate name",
			rows:    pgxmock.NewRows([]string{"id"}),
			wantErr: storage.ErrUserExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPool, err := pgxmock.NewPool()
			if err != nil {
				t.Fatal(err)
			}
			defer mockPool.Close()

			mockPool.ExpectBegin()
			mockPool.ExpectQuery(regexp.QuoteMeta(`INSERT INTO users`)).
				WithArgs(id, "login", "password").WillReturnRows(tt.rows)
			mockPool.ExpectCommit()

			u := &models.User{
				ID:       id,
				Login:    "login",
				Password: "password",
			}
			db := &DB{conn: mockPool}
			err = db.CreateUser(context.Background(), u)
			assert.ErrorIs(t, err, tt.wantErr)

			if err := mockPool.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestDB_GetIdPassByName(t *testing.T) {
	id, err := uuid.NewRandom()
	assert.NoError(t, err)

	tests := []struct {
		name    string
		rows    *pgxmock.Rows
		wantID  uuid.UUID
		wantErr error
	}{
		{
			name:    "Get id",
			rows:    pgxmock.NewRows([]string{"id", "password"}).AddRow(id, "password"),
			wantID:  id,
			wantErr: nil,
		},
		{
			name:    "Try to get id for nonexistent user",
			rows:    pgxmock.NewRows([]string{"id", "password"}),
			wantID:  uuid.Nil,
			wantErr: storage.ErrUserNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock, err := pgxmock.NewPool()
			if err != nil {
				t.Fatal(err)
			}
			defer mock.Close()

			mock.ExpectQuery("SELECT id, password").WithArgs("login").WillReturnRows(tt.rows)
			db := &DB{conn: mock}
			uid, password, err := db.GetIdPassByName(context.Background(), "login")
			if tt.wantErr == nil {
				assert.Equal(t, "password", password)
			}
			assert.Equal(t, tt.wantID, uid)
			assert.ErrorIs(t, err, tt.wantErr)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
