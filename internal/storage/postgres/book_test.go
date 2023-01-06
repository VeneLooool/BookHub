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

func TestBookStorage_CreateBook(t *testing.T) {
	type args struct {
		ctx    context.Context
		book   entity.Book
		repoId int64
	}
	type mockBehavior func(args args, id int64)

	mockDB, mockSQL, _ := sqlmock.New()
	defer mockDB.Close()

	bs := NewBookStorage(sqlx.NewDb(mockDB, "sqlmock"))

	tests := []struct {
		name    string
		args    args
		mock    mockBehavior
		want    int64
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				ctx: context.Background(),
				book: entity.Book{
					File: entity.File{Path: "abcd"},
				},
				repoId: 1,
			},
			want: 1,
			mock: func(args args, id int64) {
				mockSQL.ExpectBegin()
				mockSQL.ExpectExec(`INSERT INTO book`).
					WithArgs(args.book.Title, args.book.Author, args.book.NumberPages, args.book.Desc, args.book.Image.Path, args.book.File.Path).
					WillReturnResult(sqlmock.NewResult(id, 1))
				mockSQL.ExpectExec(`INSERT INTO repo_books`).
					WithArgs(1, args.book.CurrentPage, id).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mockSQL.ExpectCommit()
			},
		},
		{
			name: "error in bd",
			args: args{
				ctx: context.Background(),
				book: entity.Book{
					File: entity.File{Path: "fgh"},
				},
				repoId: 1,
			},
			want:    0,
			wantErr: true,
			mock: func(args args, id int64) {
				mockSQL.ExpectBegin()
				mockSQL.ExpectExec(`INSERT INTO book`).
					WithArgs(args.book.Title, args.book.Author, args.book.NumberPages, args.book.Desc, args.book.Image.Path, args.book.File.Path).
					WillReturnError(errors.New("error in bd"))
				mockSQL.ExpectRollback()
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock(test.args, test.want)

			res, err := bs.CreateBook(test.args.ctx, test.args.repoId, test.args.book)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, res)
			}
		})
	}
}

func TestBookStorage_GetBooksForRepo(t *testing.T) {
	type args struct {
		ctx    context.Context
		repoID int64
	}
	type mockBehavior func(args args)

	mockDB, mockSQL, _ := sqlmock.New()
	defer mockDB.Close()

	bs := NewBookStorage(sqlx.NewDb(mockDB, "sqlmock"))

	tests := []struct {
		name    string
		args    args
		mock    mockBehavior
		want    []entity.Book
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				ctx:    context.Background(),
				repoID: 1,
			},
			want: []entity.Book{
				{ID: 1, Title: "a", Author: "a", NumberPages: 100, CurrentPage: 0, Desc: "a", Image: entity.File{}, File: entity.File{}},
				{ID: 2, Title: "b", Author: "b", NumberPages: 100, CurrentPage: 0, Desc: "b", Image: entity.File{}, File: entity.File{}},
				{ID: 3, Title: "c", Author: "c", NumberPages: 100, CurrentPage: 0, Desc: "c", Image: entity.File{}, File: entity.File{}},
			},
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"book_id", "title", "author", "number_pages", "desc", "current_page"}).
					AddRow(1, "a", "a", 100, "a", 0).
					AddRow(2, "b", "b", 100, "b", 0).
					AddRow(3, "c", "c", 100, "c", 0)

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

			res, err := bs.GetBooksForRepo(test.args.ctx, test.args.repoID)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, res)
			}
		})
	}
}

func TestBookStorage_UpdateBook(t *testing.T) {
	type args struct {
		ctx  context.Context
		book entity.Book
	}
	type mockBehavior func(args args)

	mockDB, mockSQL, _ := sqlmock.New()
	defer mockDB.Close()

	bs := NewBookStorage(sqlx.NewDb(mockDB, "sqlmock"))

	tests := []struct {
		name    string
		args    args
		mock    mockBehavior
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				ctx: context.Background(),
				book: entity.Book{
					ID:   1,
					File: entity.File{Path: "abcd"},
				},
			},
			mock: func(args args) {
				mockSQL.ExpectExec(`UPDATE books SET (.+) WHERE (.+)`).
					WithArgs(args.book.Title, args.book.Author, args.book.NumberPages, args.book.Desc, args.book.Image.Path, args.book.File.Path, args.book.ID).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: "error in bd",
			args: args{
				ctx: context.Background(),
				book: entity.Book{
					ID:   1,
					File: entity.File{Path: "fgh"},
				},
			},
			wantErr: true,
			mock: func(args args) {
				mockSQL.ExpectBegin()
				mockSQL.ExpectExec(`UPDATE books SET (.+) WHERE (.+)`).
					WithArgs(args.book.Title, args.book.Author, args.book.NumberPages, args.book.Desc, args.book.Image.Path, args.book.File.Path, args.book.ID).
					WillReturnError(errors.New("error in bd"))
				mockSQL.ExpectRollback()
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock(test.args)

			err := bs.UpdateBook(test.args.ctx, test.args.book)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestBookStorage_DeleteBook(t *testing.T) {
	type args struct {
		ctx    context.Context
		bookID int64
	}
	type mockBehavior func(args args)

	mockDB, mockSQL, _ := sqlmock.New()
	defer mockDB.Close()

	bs := NewBookStorage(sqlx.NewDb(mockDB, "sqlmock"))

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
				bookID: 1,
			},
			mock: func(args args) {
				mockSQL.ExpectBegin()
				mockSQL.ExpectExec(`DELETE FROM books WHERE (.+)`).
					WithArgs(args.bookID).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mockSQL.ExpectExec(`DELETE FROM repo_books WHERE (.+)`).
					WithArgs(args.bookID).
					WillReturnResult(sqlmock.NewResult(0, 10))
				mockSQL.ExpectCommit()
			},
		},
		{
			name: "error in bd",
			args: args{
				ctx:    context.Background(),
				bookID: 1,
			},
			wantErr: true,
			mock: func(args args) {
				mockSQL.ExpectBegin()
				mockSQL.ExpectExec(`DELETE FROM books WHERE (.+)`).
					WithArgs(args.bookID).
					WillReturnError(errors.New("error in bd"))
				mockSQL.ExpectRollback()
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock(test.args)

			err := bs.DeleteBook(test.args.ctx, test.args.bookID)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
