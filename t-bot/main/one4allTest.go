package main

import (
	"fmt"
  "github.com/AshidaSRS/Go-things/t-bot/one4all"
)

func TestGetMe() {
  token := "138973547:AAE0mVpkk0Bu982ez-_GgEowyWtLC_JGC9U"
	bot := one4all.New(token)
	me := one4all.getMe(bot)
	fmt.Println(me.Id)
  fmt.Println(me.Username)
}

func main(){
  TestGetMe()
}
