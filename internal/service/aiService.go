package service

import (
	"go-gin/internal/model"
	"strings"
)

func ASK(msgList *[]model.Message) string {
	var result []string
	for _, msg := range *msgList {
		result = append(append(result, msg.Content), msg.ResponseContent)
	}

	question := strings.Join(result, " ")

}
