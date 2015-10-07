package one4all

import(
  "encoding/json"
  "errors"
  "strings"
  "fmt"

  "github.com/parnurzeal/gorequest"
  "github.com/AshidaSRS/Go-Things/t-bot/types"

)

func (bot TelegramBot) GetMe()(types.User,error){
  url := fmt.Sprintf(bot.BaseRequestURL, "getMe");
  _, body, errs := gorequest.New().Get(url).End()
  errormsg := "error"
  if errs != nil{
    return types.User{},errors.New(errormsg)
  }
  var data *types.ResultGetUser
	dec := json.NewDecoder(strings.NewReader(body))
	dec.Decode(&data)

  if !data.Ok {
    return types.User{},errors.New(errormsg)
  }
  return data.Result, nil
}
