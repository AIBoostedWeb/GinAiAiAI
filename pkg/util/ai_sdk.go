package util

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"go-gin/internal/config"
	"go-gin/internal/model"
)

type Role string

const (
	USER      Role = "user"
	ASSISTANT Role = "assistant"
)

// SYSTEM    Role = "system"

type AIClient struct {
	client *openai.Client
	config *config.AIConfig
}

func GenOwnClient(aiConf *config.AIConfig) *AIClient {
	llmConf := openai.DefaultConfig(aiConf.AuthKey)
	llmConf.BaseURL = aiConf.BaseURL

	customClient := openai.NewClientWithConfig(llmConf)

	return &AIClient{client: customClient, config: aiConf}
}

func (c *AIClient) RequestWithSdkMessage(messageList *[]openai.ChatCompletionMessage) (string, error) {

	req := openai.ChatCompletionRequest{
		Model:       c.config.Model,
		Messages:    *messageList,
		MaxTokens:   100,
		Temperature: 0.7, // 控制生成随机性（0-2，值越高越发散）
		//ResponseFormat: &openai.ChatCompletionResponseFormat{Type: openai.ChatCompletionResponseFormatTypeJSONSchema,&openai.ChatCompletionResponseFormatJSONSchema{Name: "name", Description: "", Schema: }},
	}

	resp, err := c.client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("create chat completion request: %w", err)
	}
	return resp.Choices[0].Message.Content, err
}

func (c *AIClient) RequestWithOwnMessage(_messageList *[]model.Message) (string, error) {
	messageList := MessageTransformer(_messageList)
	return c.RequestWithSdkMessage(messageList)
}

//-------------util--------------------

func MessageTransformer(msgList *[]model.Message) *[]openai.ChatCompletionMessage {
	var result []openai.ChatCompletionMessage
	for _, msg := range *msgList {
		result = append(result, openai.ChatCompletionMessage{Role: string(USER), Content: msg.Content})
		result = append(result, openai.ChatCompletionMessage{Role: string(ASSISTANT), Content: msg.ResponseContent})
	}
	return &result
}

func GenMessage(role Role, content string) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{Role: string(role), Content: content}

}
