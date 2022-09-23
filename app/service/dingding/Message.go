package dingding

import (
	"account_check/app/utils"
	"log"
)

const GroupUrl = "http://api.open.netjoy.com/staff/send_message"

func SendGroup(message string, chatId string, title string) bool {
	data := make(map[string]string)
	data["text"] = message
	data["title"] = title
	data["chat_id"] = chatId

	res, _ := utils.HttpSendFormResJson(GroupUrl, "POST", data, nil, nil)
	log.Println(res.String())

	return true
}
