package postgres

import (
	"bookhub/internal/entity"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepoStorage_GetReposForUser(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID int64
	}
	type mockBehavior func(args args)

	mockDB, mockSQL, _ := sqlmock.New()
	defer mockDB.Close()

	bs := NewRepoStorage(sqlx.NewDb(mockDB, "sqlmock"))

	tests := []struct {
		name    string
		args    args
		mock    mockBehavior
		want    []entity.Repo
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			want: []entity.Repo{
				{1, "a", "true", "a", 1},
				{2, "b", "true", "b", 1},
				{3, "c", "true", "c", 1},
			},
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"repo_id", "name", "visible", "repo_desc", "user_id"}).
					AddRow(1, "a", "true", "a", 1).
					AddRow(2, "b", "true", "b", 1).
					AddRow(3, "c", "true", "c", 1)

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

			res, err := bs.GetReposForUser(test.args.ctx, test.args.userID)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, res)
			}
		})
	}
}
