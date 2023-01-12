package telegram

import (
	"context"
	"fmt"
	"github.com/VeneLooool/BookHub/internal/entity"
	botApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io"
	"net/http"
	"strconv"
)

//createBook repoID Title Author NumberPages CurrentPage Desc

func (bot *TelegramBot) CreateBook(ctx context.Context, update botApi.Update) error {
	args, err := bot.parseArguments(update.Message.Text, 7)
	if err != nil {
		return fmt.Errorf("parseArguments: %w", err)
	}
	repoId, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("atoi: %w", err)
	}
	numberPages, err := strconv.Atoi(args[4])
	if err != nil {
		return fmt.Errorf("atoi: %w", err)
	}
	currentPage, err := strconv.Atoi(args[5])
	if err != nil {
		return fmt.Errorf("atoi: %w", err)
	}

	fileConfig, err := bot.bot.GetFile(botApi.FileConfig{FileID: update.Message.Document.FileID})
	if err != nil {
		return fmt.Errorf("GetFile: %w", err)
	}
	link := fileConfig.Link(bot.token)

	resp, err := http.Get(link)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("http.Get: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status: %d", resp.StatusCode)
	}

	file, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("readAll: %w", err)
	}

	bookId, err := bot.bookUseCase.CreateBook(ctx, int64(repoId), entity.Book{
		Title:       args[2],
		Author:      args[3],
		NumberPages: int64(numberPages),
		CurrentPage: int64(currentPage),
		Desc:        args[6],
		File: entity.File{
			Size: int64(fileConfig.FileSize),
			Type: entity.PDF,
			File: file,
		},
	})
	if err != nil {
		return fmt.Errorf("createBook: %w", err)
	}

	msg := botApi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Status code: 200; BookId is %d", bookId))
	if _, err = bot.bot.Send(msg); err != nil {
		return fmt.Errorf("send: %w", err)
	}

	return nil
}

//getBook bookId

func (bot *TelegramBot) GetBook(ctx context.Context, update botApi.Update) error {
	args, err := bot.parseArguments(update.Message.Text, 2)
	if err != nil {
		return fmt.Errorf("parseArguments: %w", err)
	}
	bookId, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("atoi: %w", err)
	}

	book, err := bot.bookUseCase.GetBook(ctx, int64(bookId))
	if err != nil {
		return fmt.Errorf("GetBook: %w", err)
	}

	msgText, err := bot.transformEntityInString(book)
	if err != nil {
		return fmt.Errorf("transform: %w", err)
	}

	msg := botApi.NewMessage(update.Message.Chat.ID, msgText)
	if _, err = bot.bot.Send(msg); err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}

//getBookFile bookId

func (bot *TelegramBot) GetBookFile(ctx context.Context, update botApi.Update) error {
	args, err := bot.parseArguments(update.Message.Text, 2)
	if err != nil {
		return fmt.Errorf("parseArguments: %w", err)
	}
	bookId, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("atoi: %w", err)
	}

	file, err := bot.bookUseCase.GetBookFile(ctx, int64(bookId))
	if err != nil {
		return fmt.Errorf("GetBookFile: %w", err)
	}
	msg := botApi.NewDocumentUpload(update.Message.Chat.ID, file.Path)
	msg.Caption = file.Name
	if _, err = bot.bot.Send(msg); err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}

//getBooksForRepo repoId

func (bot *TelegramBot) GetBooksForRepo(ctx context.Context, update botApi.Update) error {
	args, err := bot.parseArguments(update.Message.Text, 2)
	if err != nil {
		return fmt.Errorf("parseArguments: %w", err)
	}
	repoId, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("atoi: %w", err)
	}

	books, err := bot.bookUseCase.GetBooksForRepo(ctx, int64(repoId))
	if err != nil {
		return fmt.Errorf("GetBooksForRepo: %w", err)
	}

	msgText, err := bot.transformEntityInString(books)
	if err != nil {
		return fmt.Errorf("transform: %w", err)
	}

	msg := botApi.NewMessage(update.Message.Chat.ID, msgText)
	if _, err = bot.bot.Send(msg); err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}

//updateBook bookId Title Author NumberPages Desc

func (bot *TelegramBot) UpdateBook(ctx context.Context, update botApi.Update) error {
	args, err := bot.parseArguments(update.Message.Text, 6)
	if err != nil {
		return fmt.Errorf("parseArguments: %w", err)
	}
	bookId, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("atoi: %w", err)
	}
	numberPages, err := strconv.Atoi(args[4])
	if err != nil {
		return fmt.Errorf("atoi: %w", err)
	}

	book, err := bot.bookUseCase.UpdateBook(ctx, entity.Book{
		ID:          int64(bookId),
		Title:       args[2],
		Author:      args[3],
		NumberPages: int64(numberPages),
		Desc:        args[5],
	})
	if err != nil {
		return fmt.Errorf("updateBook: %w", err)
	}

	msgText, err := bot.transformEntityInString(book)
	if err != nil {
		return fmt.Errorf("transform: %w", err)
	}

	msg := botApi.NewMessage(update.Message.Chat.ID, msgText)
	if _, err = bot.bot.Send(msg); err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}

//deleteBook bookId

func (bot *TelegramBot) DeleteBook(ctx context.Context, update botApi.Update) error {
	args, err := bot.parseArguments(update.Message.Text, 2)
	if err != nil {
		return fmt.Errorf("parseArguments: %w", err)
	}
	bookId, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("atoi: %w", err)
	}

	err = bot.bookUseCase.DeleteBook(ctx, int64(bookId))
	if err != nil {
		return fmt.Errorf("deleteBook: %w", err)
	}

	msgText := "status code: 200"
	msg := botApi.NewMessage(update.Message.Chat.ID, msgText)
	if _, err = bot.bot.Send(msg); err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}
