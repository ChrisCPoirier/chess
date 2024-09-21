package anthropic

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	log "github.com/sirupsen/logrus"
)

func New(name, color string) Anthropic {
	client := anthropic.NewClient(os.Getenv("ANTHROPIC_API_KEY"))
	return Anthropic{
		name:   name,
		color:  color,
		client: client,
	}
}

type Anthropic struct {
	name   string
	color  string
	client *anthropic.Client
}

const defaultPrompt = `We are playing chess and you are a chess GM. You MUST respond in Algebraic notation. Do not include additional numbers or periods in your response.`

func (a Anthropic) Ask(current string) (string, error) {
	prompt := fmt.Sprintf(`%s It is your turn. You are playing as %s. This is the current board in PGN format: %s`, defaultPrompt, a.color, current)

	if current == "\n *" {
		prompt = fmt.Sprintf(`%s It is your turn, the board is in the default position and you are the first to move`, defaultPrompt)
	}

	resp, err := a.client.Messages.Create(
		context.Background(),
		&anthropic.MessageCreateParams{
			Model: anthropic.Claude3Sonnet20240229,
			MaxTokens: 100,
			Messages: []anthropic.Message{
				{
					Role: anthropic.MessageRoleUser,
					Content: []anthropic.Content{
						{
							Type: anthropic.ContentTypeText,
							Text: prompt,
						},
					},
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	if len(resp.Content) == 0 || resp.Content[0].Type != anthropic.ContentTypeText {
		return "", fmt.Errorf("unexpected response format")
	}

	respText := a.SanitizeResponse(resp.Content[0].Text)
	log.Debugf("response is `%s`", respText)

	return respText, nil
}

func (a Anthropic) SanitizeResponse(s string) string {
	i := strings.LastIndex(s, " ")

	if i == -1 {
		return s
	}

	return s[i+1:]
}
