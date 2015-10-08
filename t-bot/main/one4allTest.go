package main

import (
	"fmt"
  "github.com/AshidaSRS/Go-things/t-bot/one4all"
)

func TestGetMe() {
  token :=
	bot := one4all.New(token)
	me, _ := bot.GetMe()
  msg,_ := bot.GetUpdates()
  fmt.Println(msg[1].UpdateId)
  fmt.Println(msg[1].Message)
	fmt.Println(me.Id)
  fmt.Println(me.Username)
}

func main(){
  TestGetMe()
}
