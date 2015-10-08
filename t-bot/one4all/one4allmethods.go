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
  errormsg := "error getMe"
  if errs != nil{
    return types.User{},errors.New(errormsg)
  }
  var data *types.RGetUser
	dec := json.NewDecoder(strings.NewReader(body))
	dec.Decode(&data)

  if !data.Ok {
    return types.User{},errors.New(errormsg)
  }
  return data.Result, nil
}


func (bot TelegramBot) GetUpdates() ([]types.Update, error) {
  url := fmt.Sprintf(bot.BaseRequestURL, "getUpdates");
  errormsg := "error getUpdate"

  _, body, errs := gorequest.New().Get(url).End()

	if errs != nil {
		return []types.Update{}, errors.New(errormsg)
	}

	var data *types.RGetUpdates
	json.Unmarshal([]byte(body), &data)

	if !data.Ok {
		return []types.Update{}, errors.New(errormsg)
  }
	return data.Result, nil
}
