package postgres

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/VeneLooool/BookHub/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserStorage_GetUser(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID int64
	}
	type mockBehavior func(args args)

	mockDB, mockSQL, _ := sqlmock.New()
	defer mockDB.Close()

	bs := NewUserStorage(sqlx.NewDb(mockDB, "sqlmock"))

	tests := []struct {
		name    string
		args    args
		mock    mockBehavior
		want    entity.User
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			want: entity.User{ID: 1, Name: "a", UserName: "a", Password: "a", Desc: "a"},
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"user_id", "name", "username", "password", "user_desc"}).
					AddRow(1, "a", "a", "a", "a")

				mockSQL.ExpectQuery(`SELECT (.+)`).
					WithArgs(args.userID).WillReturnRows(rows)
			},
		},
		{
			name: "error in bd",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			wantErr: true,
			mock: func(args args) {
				mockSQL.ExpectQuery(`SELECT (.+)`).
					WithArgs(args.userID).WillReturnError(errors.New("error in bd"))
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock(test.args)

			res, err := bs.GetUser(test.args.ctx, test.args.userID)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, res)
			}
		})
	}
}
