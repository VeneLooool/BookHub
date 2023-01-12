package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

var MainMenuButton = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("repositories"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("profile"),
		tgbotapi.NewKeyboardButton("books"),
	),
)

var RepoMenuButton = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("create new repository"),
		tgbotapi.NewKeyboardButton("get list of yours repositories"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("back to main menu"),
	),
)

var BookMenuButton = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("create new book"),
		tgbotapi.NewKeyboardButton("get list of books for repository"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("back to repositories"),
	),
)

var ProfileMenuButton = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("update your profile"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("back to main menu"),
	),
)
var (
	AddMailServiceDescription     = "_ADDMAIL"
	SettingMailServiceDescription = "_SETTINGS"
	TurnOnConstUpdateSettings     = "_TurnOnContUpdateSettings"
	TurnOffConstUpdateSettings    = "_TurnOffContUpdateSettings"
	GetLastMessageSettings        = "_GetLastMessageSettings"
)
