package agent

import (
	"context"
	"fmt"

	"investment/adapter/ai/tools"
	"investment/utility"
	"time"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
)

// 财联社分析师
type ClsAnalyst struct {
	inputMessage []*schema.Message
	agent        *react.Agent
}

func NewClsAnalyst(ctx context.Context, roles []*schema.Message, inputMsg *schema.Message) (result *ClsAnalyst, e error) {
	result = &ClsAnalyst{}
	result.inputMessage = make([]*schema.Message, 0)
	for _, v := range roles {
		result.inputMessage = append(result.inputMessage, v)
	}
	result.inputMessage = append(result.inputMessage, schema.SystemMessage(fmt.Sprintf("Current time is %s", time.Now())))
	result.inputMessage = append(result.inputMessage, inputMsg)

	clsTelegramSearch, e := tools.NewClsTelegramSearch()
	if e != nil {
		return
	}
	clsDepthSearch, e := tools.NewClsDepthSearch()
	if e != nil {
		return
	}
	clsDetail, e := tools.NewClsDetail()

	tools := []tool.BaseTool{
		clsTelegramSearch,
		clsDepthSearch,
		clsDetail,
	}

	result.agent, e = react.NewAgent(ctx, &react.AgentConfig{
		Model:           utility.GetQwenChatModel(),
		ToolsConfig:     compose.ToolsNodeConfig{Tools: tools},
		MessageModifier: nil,
		MaxStep:         66,
	})
	if e != nil {
		return
	}
	return result, nil
}

func (ra *ClsAnalyst) GetInputMessage() (result []*schema.Message) {
	return ra.inputMessage
}

func (ra *ClsAnalyst) Run(ctx context.Context) (result *schema.Message, e error) {
	return ra.agent.Generate(ctx, ra.inputMessage)
}
