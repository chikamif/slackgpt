package chatgpt

import (
	"context"
	"errors"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

// ErrorEmptyPrompt implements an Error raised by passing an empty prompt
var ErrorEmptyPrompt error = errors.New("Error empty prompt")

// GetStringResponse sends a completion request to the GPT-3 API to generate a response
// for a given conversation using the specified GPT-3 model. The function takes in a GPT-3
// client, a context, and a slice of strings representing the conversation.
//
// If the length of the conversation slice is 0, an error called ErrorEmptyPrompt is returned.
//
// The function returns the generated response text from the GPT-3 API as a string, with any leading
// or trailing spaces removed using strings.TrimSpace().
//
// Parameters:
// - client: a GPT-3 client object used to make API requests
// - ctx: a context object used to handle timeouts and cancellations
// - chat: a slice of strings representing the conversation
//
// Returns:
// - a string containing the generated response from the GPT-3 API
// - an error, if any
func GetStringResponse(client *openai.Client, ctx context.Context, chat []string) (string, error) {
	if len(chat) == 0 {
		return "", ErrorEmptyPrompt
	}

	req := openai.ChatCompletionRequest{
		Model: openai.GPT4Turbo1106,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a helpful chat bot assistant. Please answer shortly, and in Japanese.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: strings.Join(chat, " "),
			},
		},
		MaxTokens:   1000,
		Temperature: 0.5,
	}
	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(resp.Choices[0].Message.Content), nil
}
