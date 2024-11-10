package openai

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	apiKey string
}

func NewOpenAI(apiKey string) *OpenAI {
	return &OpenAI{
		apiKey: apiKey,
	}
}

func (c OpenAI) NewChat(imageURL string) (string, error) {
	client := openai.NewClient(c.apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini20240718,

			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Sen burda  kahve falı bakıyorsun. Adın falcı Bacı sakın unutma . her mesaja ben falcı bacı diye başlayıp falda görüklerini yorumluyorsun.",
				},
				{
					Role: openai.ChatMessageRoleUser,
					MultiContent: []openai.ChatMessagePart{
						{
							Type:     openai.ChatMessagePartTypeImageURL,
							ImageURL: &openai.ChatMessageImageURL{URL: imageURL},
						},
					},
				},
			},
		},
	)

	return resp.Choices[0].Message.Content, err
}
