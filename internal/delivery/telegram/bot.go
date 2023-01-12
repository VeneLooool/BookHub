package telegram

import (
	"context"
	"errors"
	"fmt"
	"github.com/VeneLooool/BookHub/internal/config"
	"github.com/VeneLooool/BookHub/internal/entity"
	"github.com/VeneLooool/BookHub/internal/service"
	botApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/labstack/gommon/log"
	"strings"
	"sync"
)

//TODO DELETE THIS SHIT

type Router func(context.Context, botApi.Update) error

type TelegramUser struct {
	telegramId int
	internalId int64
	chatId     int
	update     botApi.Update
}

type TelegramBot struct {
	users         map[int]*TelegramUser
	bot           *botApi.BotAPI
	botMutex      sync.Mutex
	updatesConfig botApi.UpdateConfig
	updates       botApi.UpdatesChannel
	ctx           context.Context
	mux           map[string]Router
	token         string

	userUseCase service.UserUseCase
	repoUseCase service.RepoUseCase
	bookUseCase service.BookUseCase
}

func NewTelegramBot(ctx context.Context, config config.TelegramBotConfig, uc service.UserUseCase, rc service.RepoUseCase, bc service.BookUseCase) (*TelegramBot, error) {
	//TODO delete inmemory db
	var err error
	bot := TelegramBot{
		userUseCase: uc,
		repoUseCase: rc,
		bookUseCase: bc,
		token:       config.Token,
	}

	bot.users = make(map[int]*TelegramUser, 0)

	fmt.Println(config.Token)
	bot.bot, err = botApi.NewBotAPI(config.Token)
	if err != nil {
		return nil, fmt.Errorf("NewBotApi: %w", err)
	}

	bot.updatesConfig = botApi.NewUpdate(0)
	bot.updatesConfig.Timeout = config.Timeout

	bot.updates, err = bot.bot.GetUpdatesChan(bot.updatesConfig)
	if err != nil {
		return nil, fmt.Errorf("GetUpdatesChan: %w", err)
	}

	bot.InitRouter()
	return &bot, nil
}

func (bot *TelegramBot) InitRouter() {
	bot.mux = make(map[string]Router, 11)

	bot.mux["/createUser"] = bot.CreateUser
	bot.mux["/getUser"] = bot.GetUser
	bot.mux["/updateUser"] = bot.UpdateUser
	bot.mux["/deleteUser"] = bot.DeleteUser

	bot.mux["/createRepo"] = bot.CreateRepo
	bot.mux["/getReposForUser"] = bot.GetReposForUser
	bot.mux["/getRepo"] = bot.GetRepo
	bot.mux["/updateRepo"] = bot.UpdateRepo
	bot.mux["/deleteRepo"] = bot.DeleteRepo
	bot.mux["/deleteBookFromRepo"] = bot.DeleteBookFromRepo

	bot.mux["/createBook"] = bot.CreateBook
	bot.mux["/getBook"] = bot.GetBook
	bot.mux["/getBookFile"] = bot.GetBookFile
	bot.mux["/getBooksForRepo"] = bot.GetBooksForRepo
	bot.mux["/updateBook"] = bot.UpdateBook
	bot.mux["/deleteBook"] = bot.DeleteBook

}

func (bot *TelegramBot) StartTelegramBot() error {

	for update := range bot.updates {
		if update.Message != nil {
			fmt.Println(update.Message.Text)
			ctx := context.Background()

			args := strings.Split(update.Message.Text, " ")
			if update.Message.Text == "" && update.Message.Caption == "" {
				return errors.New("not a command")
			}
			if update.Message.Text == "" && update.Message.Caption != "" {
				args = strings.Split(update.Message.Caption, " ")
				update.Message.Text = update.Message.Caption
			}

			route, ok := bot.mux[args[0]]
			if !ok {
				msg := botApi.NewMessage(update.Message.Chat.ID, "command not found")
				if _, err := bot.bot.Send(msg); err != nil {
					return fmt.Errorf("send: %w", err)
				}
				continue
			}

			err := route(ctx, update)
			if err != nil {
				log.Errorf("%s: %w", args[0], err)
				msg := botApi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("error: %s", err))
				if _, err = bot.bot.Send(msg); err != nil {
					return fmt.Errorf("send: %w", err)
				}
			}
		}
	}
	return nil
}

func (bot *TelegramBot) findTelegramUser(ctx context.Context, telegramID int) (*TelegramUser, error) {
	user, ok := bot.users[telegramID]
	if !ok {
		return nil, entity.ErrUserNotFound
	}
	return user, nil
}

func (bot *TelegramBot) createTelegramUser(ctx context.Context, user *TelegramUser) error {
	if user.telegramId == 0 {
		return fmt.Errorf("telegramId is 0")
	}
	bot.users[user.telegramId] = user
	return nil
}

func (bot *TelegramBot) parseArguments(lineArgs string, amount int) ([]string, error) {
	args := strings.Split(lineArgs, " ")
	if len(args) != amount {
		return nil, fmt.Errorf("wrong number of arguments")
	}
	return args, nil
}
