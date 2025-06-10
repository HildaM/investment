package agent

import (
	"context"
	"fmt"
	"investment/adapter/ai/prompts"
	"investment/utility"

	"time"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/schema"
)

// 选eth分析师
type SelectAnalyst struct {
	inputMessage []*schema.Message
	agent        *openai.ChatModel
}

func NewSelectAnalyst(ctx context.Context, inputMsg *schema.Message) (result *SelectAnalyst, e error) {
	result = &SelectAnalyst{}
	result.inputMessage = make([]*schema.Message, 0)
	result.inputMessage = append(result.inputMessage, schema.SystemMessage(prompts.AgentSelectRole))
	result.inputMessage = append(result.inputMessage, schema.SystemMessage(fmt.Sprintf("现在时间是 %s", time.Now())))
	result.inputMessage = append(result.inputMessage, inputMsg)
	result.agent = utility.GetTTDeepSeekChatModel()
	return
}

func (sa *SelectAnalyst) GetInputMessage() (result []*schema.Message) {
	return sa.inputMessage
}

func (sa *SelectAnalyst) Run(ctx context.Context) (result *schema.Message, e error) {
	return sa.agent.Generate(ctx, sa.inputMessage)
}
