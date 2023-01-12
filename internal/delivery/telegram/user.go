package telegram

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/VeneLooool/BookHub/internal/entity"
	botApi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//createUser Name UserName Password Desc

func (bot *TelegramBot) CreateUser(ctx context.Context, update botApi.Update) error {
	bot.botMutex.Lock()
	defer bot.botMutex.Unlock()
	user := update.Message.From

	if _, ok := bot.users[user.ID]; ok {
		msg := botApi.NewMessage(update.Message.Chat.ID, "status code: 200")
		if _, err := bot.bot.Send(msg); err != nil {
			return fmt.Errorf("send: %w", err)
		}
		return nil
	}

	args, err := bot.parseArguments(update.Message.Text, 5)
	if err != nil {
		return fmt.Errorf("parseArguments: %w", err)
	}

	internalId, err := bot.userUseCase.CreateUser(ctx, entity.User{
		Name:     args[1],
		UserName: args[2],
		Password: args[3],
		Desc:     args[4],
	})
	if err != nil && !errors.Is(err, entity.ErrUserAlreadyExists) {
		return fmt.Errorf("CreateUser: %w", err)
	}

	bot.users[user.ID] = &TelegramUser{
		telegramId: user.ID,
		chatId:     int(update.Message.Chat.ID),
		internalId: internalId,
	}

	msg := botApi.NewMessage(update.Message.Chat.ID, "Status code: 200")
	if _, err = bot.bot.Send(msg); err != nil {
		return fmt.Errorf("send: %w", err)
	}

	return nil
}

func (bot *TelegramBot) GetUser(ctx context.Context, update botApi.Update) error {
	tgUser := update.Message.From

	user, ok := bot.users[tgUser.ID]
	if !ok {
		return entity.ErrUserNotFound
	}

	internalUser, err := bot.userUseCase.GetUser(ctx, user.internalId)
	if err != nil {
		return fmt.Errorf("getUser: %w", err)
	}

	messageText, err := bot.transformInternalUserInString(internalUser)
	if err != nil {
		return fmt.Errorf("transformInternalUserInString: %w", err)
	}

	msg := botApi.NewMessage(update.Message.Chat.ID, messageText)
	if _, err = bot.bot.Send(msg); err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}

//updateUser Name UserName Password Desc

func (bot *TelegramBot) UpdateUser(ctx context.Context, update botApi.Update) error {
	tgUser := update.Message.From

	user, ok := bot.users[tgUser.ID]
	if !ok {
		return entity.ErrUserNotFound
	}

	args, err := bot.parseArguments(update.Message.Text, 5)
	if err != nil {
		return fmt.Errorf("parseArguments: %w", err)
	}

	newUser, err := bot.userUseCase.UpdateUser(ctx, entity.User{
		ID:       user.internalId,
		Name:     args[1],
		UserName: args[2],
		Password: args[3],
		Desc:     args[4],
	})

	messageText, err := bot.transformInternalUserInString(newUser)
	if err != nil {
		return fmt.Errorf("transformInternalUserInString: %w", err)
	}

	msg := botApi.NewMessage(update.Message.Chat.ID, messageText)
	if _, err = bot.bot.Send(msg); err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}

func (bot *TelegramBot) DeleteUser(ctx context.Context, update botApi.Update) error {
	bot.botMutex.Lock()
	defer bot.botMutex.Unlock()
	tgUser := update.Message.From

	user, ok := bot.users[tgUser.ID]
	if !ok {
		return entity.ErrUserNotFound
	}

	err := bot.userUseCase.DeleteUser(ctx, user.internalId)
	if err != nil {
		return fmt.Errorf("deleteUser: %w", err)
	}

	delete(bot.users, user.telegramId)
	return nil
}

func (bot *TelegramBot) transformInternalUserInString(user entity.User) (string, error) {
	jsonUser, err := json.Marshal(&user)
	if err != nil {
		return "", fmt.Errorf("marshal: %w", err)
	}
	return string(jsonUser), nil
}
