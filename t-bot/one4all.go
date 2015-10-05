package main

import(
	"io"
	"http"
	"net/url"
	"strings"
	"sync/atomic"
)
//Telegram constants
const(
  APIUrl = "https://api.telegram.org/bot%s/%s"
  FileUrl = "https://api.telegram.org/file/bot%s/%s"
)
//Struct to manage all the interactions with the api
type TelegramBot struct {
	Token                string
	FirstName            string
	ID                   int
	Username             string
	BaseRequestURL       string
	BaseFileRequestURL   string
	RelicCfg             *RelicConfig
	BotanIO              *botan.Botan
	MainListener         chan MessageWithUpdateID
	LastUpdateID         int64
	TestConditionalFuncs []ConditionCallStructure
	NoMessageFuncs       []NoMessageCall
	ChainConditionals    []*ChainStructure
	BuildingChain        bool
	DefaultOptions       DefaultOptionsBot
}
//New bot
func New(token string) (*TgBot, error) {
  aurl := fmt.Sprintf(baseURL, token, "%s")
	furl := fmt.Sprintf(fileURL, token, "%s")
  tgbot := &TelegramBot{
		Token:                token,
		BaseRequestURL:       aurl,
		BaseFileRequestURL:   furl,
		MainListener:         nil,
		RelicCfg:             nil,
		BotanIO:              nil,
		TestConditionalFuncs: make([]ConditionCallStructure, 0),
		NoMessageFuncs:       make([]NoMessageCall, 0),
		ChainConditionals:    make([]*ChainStructure, 0),
		BuildingChain:        false,
		DefaultOptions: DefaultOptionsBot{
			CleanInitialUsername:       true,
			AllowWithoutSlashInMention: true,
		},
	}

}

