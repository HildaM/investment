package agent

import (
	"context"
	"fmt"

	"investment/utility"
	"time"

	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
)

// 雅虎社分析师
type YahooAnalyst struct {
	inputMessage []*schema.Message
	agent        *react.Agent
}

func NewYahooAnalyst(ctx context.Context, roles []*schema.Message, inputMsg *schema.Message) (result *YahooAnalyst, e error) {
	result = &YahooAnalyst{}
	result.inputMessage = make([]*schema.Message, 0)
	for _, v := range roles {
		result.inputMessage = append(result.inputMessage, v)
	}
	result.inputMessage = append(result.inputMessage, schema.SystemMessage(fmt.Sprintf("Current time is %s", time.Now())))
	result.inputMessage = append(result.inputMessage, inputMsg)

	result.agent, e = react.NewAgent(ctx, &react.AgentConfig{
		Model:           utility.GetTTDeepSeekChatModel(),
		MessageModifier: nil,
		MaxStep:         30,
	})
	if e != nil {
		return
	}
	return result, nil
}

func (ra *YahooAnalyst) GetInputMessage() (result []*schema.Message) {
	return ra.inputMessage
}

func (ra *YahooAnalyst) Run(ctx context.Context) (result *schema.Message, e error) {
	return ra.agent.Generate(ctx, ra.inputMessage)
}
