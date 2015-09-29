package main

import(
	"io"
	"http"
	"net/url"
	"strings"
	"sync/atomic"
)

type TgBot struct {
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

func New(token string) *TgBot {
	bot, err := NewWithError(token)
	if err != nil {
		panic(err)
	}
	return bot
}

