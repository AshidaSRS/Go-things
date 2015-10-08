package one4all

// Send 
type Send struct {
	ChatID int
	Bot    *TelegramBot
}

type SendText struct {
	Send                  *Send
	Text                  string
	ReplyToMessageID      *int
}
