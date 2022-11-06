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

func TestRepoStorage_GetRepo(t *testing.T) {
	type args struct {
		ctx    context.Context
		repoID int64
	}
	type mockBehavior func(args args)

	mockDB, mockSQL, _ := sqlmock.New()
	defer mockDB.Close()

	bs := NewRepoStorage(sqlx.NewDb(mockDB, "sqlmock"))

	tests := []struct {
		name    string
		args    args
		mock    mockBehavior
		want    entity.Repo
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				ctx:    context.Background(),
				repoID: 1,
			},
			want: entity.Repo{ID: 1, Name: "a", Visibility: "true", Desc: "a", UserID: 1},
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"repo_id", "name", "visible", "repo_desc", "user_id"}).
					AddRow(1, "a", "true", "a", 1)

				mockSQL.ExpectQuery(`SELECT (.+)`).
					WithArgs(args.repoID).WillReturnRows(rows)
			},
		},
		{
			name: "error in bd",
			args: args{
				ctx:    context.Background(),
				repoID: 1,
			},
			wantErr: true,
			mock: func(args args) {
				mockSQL.ExpectQuery(`SELECT (.+)`).
					WithArgs(args.repoID).WillReturnError(errors.New("error in bd"))
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock(test.args)

			res, err := bs.GetRepo(test.args.ctx, test.args.repoID)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, res)
			}
		})
	}
}

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

func TestRepoStorage_DeleteRepo(t *testing.T) {
	type args struct {
		ctx    context.Context
		repoID int64
	}
	type mockBehavior func(args args)

	mockDB, mockSQL, _ := sqlmock.New()
	defer mockDB.Close()

	bs := NewRepoStorage(sqlx.NewDb(mockDB, "sqlmock"))

	tests := []struct {
		name    string
		args    args
		mock    mockBehavior
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				ctx:    context.Background(),
				repoID: 1,
			},
			mock: func(args args) {
				mockSQL.ExpectBegin()
				mockSQL.ExpectExec(`DELETE FROM repo_books WHERE (.+)`).
					WithArgs(args.repoID).
					WillReturnResult(sqlmock.NewResult(0, 10))
				mockSQL.ExpectExec(`DELETE FROM repos WHERE (.+)`).
					WithArgs(args.repoID).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mockSQL.ExpectCommit()
			},
		},
		{
			name: "error in bd",
			args: args{
				ctx:    context.Background(),
				repoID: 1,
			},
			wantErr: true,
			mock: func(args args) {
				mockSQL.ExpectBegin()
				mockSQL.ExpectExec(`DELETE FROM repo_books WHERE (.+)`).
					WithArgs(args.repoID).
					WillReturnError(errors.New("error in bd"))
				mockSQL.ExpectRollback()
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock(test.args)

			err := bs.DeleteRepo(test.args.ctx, test.args.repoID)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
