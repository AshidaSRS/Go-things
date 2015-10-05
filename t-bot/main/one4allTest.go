package main

import (
	"fmt"
  "github.com/AshidaSRS/Go-things/t-bot/one4all"
)

func TestGetMe() {
  token := 
	bot := one4all.New(token)
	me := one4all.getMe(bot)
	fmt.Println(me.Id)
  fmt.Println(me.Username)
}

func main(){
  TestGetMe()
}
