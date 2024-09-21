package anthropic

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	log "github.com/sirupsen/logrus"
)

func New(name, color string) Anthropic {
	client := anthropic.NewClient(option.WithAPIKey(os.Getenv(`ANTHROPIC_API_KEY`)))
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

	resp, err := a.client.CreateCompletion(
		context.Background(),
		&anthropic.CompletionRequest{
			Model:     anthropic.ModelClaude2,
			Prompt:    prompt,
			MaxTokens: 100,
		},
	)

	if err != nil {
		return ``, err
	}

	respText := a.SanitizeResponse(resp.Completion)
	log.Debugf("response is `%s`", respText)

	return respText, nil
}

func (a Anthropic) SanitizeResponse(s string) string {
	i := strings.LastIndex(s, ` `)

	if i == -1 {
		return s
	}

	return s[i+1:]
}
