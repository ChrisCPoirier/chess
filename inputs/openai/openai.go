package openai

import (
	"context"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
)

func New(name, color string) OpenAI {

	client := openai.NewClient(os.Getenv(`OPEN_AI_TOKEN`))
	return OpenAI{
		name:   name,
		color:  color,
		client: client,
	}
}

type OpenAI struct {
	name   string
	color  string
	client *openai.Client
}

const defaultPrompt = `We are playing chess and you are a chess GM. You MUST respond in Algebraic notation. Do not include additonal numbers or periords in your response.`

func (o OpenAI) Ask(current string) (string, error) {

	prompt := fmt.Sprintf(`%s It is your turn. You are playing as %s. This is the current board in PGN format: %s`, defaultPrompt, o.color, current)

	if current == "\n *" {
		prompt = fmt.Sprintf(`%s It is your turn, the board is in the default position and you are the first to move`, defaultPrompt)
	}

	resp, err := o.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return ``, err
	}

	respText := o.SanatizeResponse(resp.Choices[0].Message.Content)
	log.Debugf("response is `%s`", respText)

	return respText, nil
}

func (o OpenAI) SanatizeResponse(s string) string {
	i := strings.LastIndex(s, ` `)

	if i == -1 {
		return s
	}

	return s[i+1:]

}
