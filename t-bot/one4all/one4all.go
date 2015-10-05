// Golang (Go) implementation for the Telegram Bot API. (Simple one)
package one4all

import(
"encoding/json"
	"errors"
	"io"
	"strings"
  "fmt"
  "github.com/AshidaSRS/Go-things/t-bot/types"
)
//Telegram constants
const(
  APIUrl = "https://api.telegram.org/bot%s/%s"
  FileUrl = "https://api.telegram.org/file/bot%s/%s"
)
//Struct to manage interactions functions
type TelegramBot struct {
  	Token                string
	BaseRequestURL       string
	BaseFileRequestURL   string
 	Messages    	     chan (*types.Message)
}
//New bot
func New(token string) *TelegramBot {
    aurl := fmt.Sprintf(APIUrl, token, "%s")
	furl := fmt.Sprintf(FileUrl, token, "%s")

	tgbot := &TelegramBot{
	Token:                token,
	BaseRequestURL:       aurl,
	BaseFileRequestURL:   furl,
    Messages:             make(chan *types.Message),
  }
	return tgbot
}

// GetMe function :D
func (bot TgBot) GetMe() (types.User, error) {
}





