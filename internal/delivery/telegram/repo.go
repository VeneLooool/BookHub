package telegram

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/VeneLooool/BookHub/internal/entity"
	botApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

//createRepo name visibility(true/false) desc

func (bot *TelegramBot) CreateRepo(ctx context.Context, update botApi.Update) error {
	tgUser := update.Message.From

	user, ok := bot.users[tgUser.ID]
	if !ok {
		return entity.ErrUserNotFound
	}

	args, err := bot.parseArguments(update.Message.Text, 4)
	if err != nil {
		return fmt.Errorf("parseArguments: %w", err)
	}

	repoId, err := bot.repoUseCase.CreateRepo(ctx, user.internalId, entity.Repo{
		UserID:     user.internalId,
		Name:       args[1],
		Visibility: args[2],
		Desc:       args[3],
	})
	if err != nil {
		return fmt.Errorf("createRepo: %w", err)
	}

	msg := botApi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Status code: 200; RepoId is %d", repoId))
	if _, err = bot.bot.Send(msg); err != nil {
		return fmt.Errorf("Send: %w", err)
	}

	return nil
}
func (bot *TelegramBot) GetReposForUser(ctx context.Context, update botApi.Update) error {
	tgUser := update.Message.From

	user, ok := bot.users[tgUser.ID]
	if !ok {
		return entity.ErrUserNotFound
	}

	repos, err := bot.repoUseCase.GetReposForUser(ctx, user.internalId)
	if err != nil {
		return fmt.Errorf("GetReposForUser: %w", err)
	}

	var msgText string
	for _, repo := range repos {
		str, err := bot.transformEntityInString(repo)
		if err != nil {
			return fmt.Errorf("transformEntityInString: %w", err)
		}
		msgText += str + "\n"
	}

	msg := botApi.NewMessage(update.Message.Chat.ID, msgText)
	if _, err = bot.bot.Send(msg); err != nil {
		return fmt.Errorf("send: %w", err)
	}

	return nil
}

//getRepo repoId

func (bot *TelegramBot) GetRepo(ctx context.Context, update botApi.Update) error {
	args, err := bot.parseArguments(update.Message.Text, 2)
	if err != nil {
		return fmt.Errorf("parseArguments: %w", err)
	}
	repoId, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("atoi: %w", err)
	}

	repo, err := bot.repoUseCase.GetRepo(ctx, int64(repoId))
	if err != nil {
		return fmt.Errorf("getRepo: %w", err)
	}

	msgText, err := bot.transformEntityInString(repo)
	if err != nil {
		return fmt.Errorf("transform: %w", err)
	}

	msg := botApi.NewMessage(update.Message.Chat.ID, msgText)
	if _, err = bot.bot.Send(msg); err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}

//updateRepo repoId name visibility Desc

func (bot *TelegramBot) UpdateRepo(ctx context.Context, update botApi.Update) error {
	tgUser := update.Message.From

	user, ok := bot.users[tgUser.ID]
	if !ok {
		return entity.ErrUserNotFound
	}

	args, err := bot.parseArguments(update.Message.Text, 5)
	if err != nil {
		return fmt.Errorf("parseArguments: %w", err)
	}
	repoId, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("atoi: %w", err)
	}

	newRepo, err := bot.repoUseCase.UpdateRepo(ctx, entity.Repo{
		ID:         int64(repoId),
		Name:       args[2],
		Visibility: args[3],
		Desc:       args[4],
		UserID:     user.internalId,
	})
	if err != nil {
		return fmt.Errorf("updateRepo: %w", err)
	}

	msgText, err := bot.transformEntityInString(newRepo)
	if err != nil {
		return fmt.Errorf("transform: %w", err)
	}

	msg := botApi.NewMessage(update.Message.Chat.ID, msgText)
	if _, err = bot.bot.Send(msg); err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}

//deleteRepo repoId

func (bot *TelegramBot) DeleteRepo(ctx context.Context, update botApi.Update) error {
	args, err := bot.parseArguments(update.Message.Text, 2)
	if err != nil {
		return fmt.Errorf("parseArguments: %w", err)
	}

	repoId, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("atoi: %w", err)
	}

	if err = bot.repoUseCase.DeleteRepo(ctx, int64(repoId)); err != nil {
		return fmt.Errorf("deleteRepo: %w", err)
	}

	msg := botApi.NewMessage(update.Message.Chat.ID, "status code: 200")
	if _, err = bot.bot.Send(msg); err != nil {
		return fmt.Errorf("send: %w", err)
	}

	return nil
}

//deleteBookFromRepo repoId bookId

func (bot *TelegramBot) DeleteBookFromRepo(ctx context.Context, update botApi.Update) error {
	args, err := bot.parseArguments(update.Message.Text, 3)
	if err != nil {
		return fmt.Errorf("parseArguments: %w", err)
	}

	repoId, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("atoi: %w", err)
	}
	bookId, err := strconv.Atoi(args[2])
	if err != nil {
		return fmt.Errorf("atoi: %w", err)
	}

	if err = bot.repoUseCase.DeleteBookFromRepo(ctx, int64(repoId), int64(bookId)); err != nil {
		return fmt.Errorf("deleteBookFromRepo: %w", err)
	}

	msg := botApi.NewMessage(update.Message.Chat.ID, "status code: 200")
	if _, err = bot.bot.Send(msg); err != nil {
		return fmt.Errorf("send: %w", err)
	}

	return nil
}

func (bot *TelegramBot) transformEntityInString(in any) (string, error) {
	json, err := json.Marshal(&in)
	if err != nil {
		return "", fmt.Errorf("marshal: %w", err)
	}

	return string(json), nil
}
