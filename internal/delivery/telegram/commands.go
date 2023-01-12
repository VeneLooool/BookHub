package telegram

import (
	"context"
	"errors"
	"fmt"
	"github.com/VeneLooool/BookHub/internal/entity"
	botApi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *TelegramBot) commandStart(ctx context.Context, update *botApi.Update) error {
	bot.botMutex.Lock()
	defer bot.botMutex.Unlock()

	user := update.Message.From

	if _, ok := bot.users[user.ID]; ok {
		msg := botApi.NewMessage(update.Message.Chat.ID, "Hello")
		if _, err := bot.bot.Send(msg); err != nil {
			return fmt.Errorf("Send: %w", err)
		}
		return nil
	}

	internalId, err := bot.userUseCase.CreateUser(bot.ctx, entity.User{
		Name:     user.FirstName + " " + user.LastName,
		UserName: user.UserName,
	})
	if err != nil && !errors.Is(err, entity.ErrUserAlreadyExists) {
		return fmt.Errorf("CreateUser: %w", err)
	}

	bot.users[user.ID] = &TelegramUser{
		telegramId: user.ID,
		chatId:     int(update.Message.Chat.ID),
		internalId: internalId,
	}

	msg := botApi.NewMessage(update.Message.Chat.ID, "Hello")
	if _, err = bot.bot.Send(msg); err != nil {
		return fmt.Errorf("Send: %w", err)
	}

	return nil
}
