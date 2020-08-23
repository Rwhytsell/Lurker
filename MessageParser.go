package main

import (
	"Rwhytsell.dev/lurker/models"
	"regexp"
	"strconv"
	"strings"
	"time"
)


// ParseMessage parses the message received from twitch websocket and parse it into json
func ParseMessage(s string) models.Message {
	parsedMessage := models.Message{}

	infoMessageSplit := strings.Split(s, "user-type=")
	test := strings.SplitAfter(infoMessageSplit[1], "PRIVMSG")
	message := strings.SplitN(test[1], ":", 2)
	re := regexp.MustCompile(`\r?\n`)
	cleaned := re.ReplaceAllString(message[1], "")
	parsedMessage.MessageText = cleaned
	parsedMessage.Channel = strings.Trim(message[0], " #")

	splitMessageInfo := strings.Split(infoMessageSplit[0], ";")

	for _, item := range splitMessageInfo {
		splitItem := strings.Split(item, "=")
		switch splitItem[0] {
		case "display-name":
			parsedMessage.Username = splitItem[1]
		case "user-id":
			parsedMessage.UserId = splitItem[1]
		case "mod":
			parsedMessage.IsMod, _ = strconv.ParseBool(splitItem[1])
		case "subscriber":
			parsedMessage.IsSub, _ = strconv.ParseBool(splitItem[1])
		case "turbo":
			parsedMessage.IsTurbo, _ = strconv.ParseBool(splitItem[1])
		}
	}
	parsedMessage.DateSent = time.Now().UTC()

	return parsedMessage
}
