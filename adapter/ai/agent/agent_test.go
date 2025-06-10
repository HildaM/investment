package agent

import (
	"context"
	"fmt"
	"investment/adapter/ai/prompts"
	"investment/adapter/spider"
	"investment/utility"
	"testing"

	"github.com/cloudwego/eino/schema"
)

// go test -timeout 1160s -run ^TestNewClsAnalyst$ investment/adapter/ai/agent -v -count=1
func TestNewClsAnalyst(t *testing.T) {
	var inputMsg = `
分析下今晚纳斯达克的走势，你因该用更宽阔的视野和思维去思考，不因该拘泥于搜索到的数据。然后给我一份报告，报告里需明确涨跌趋势以及逻辑。当然如果没有搜到有用信息也可以放弃。
`
	role := schema.SystemMessage(prompts.AgentFinancialData)
	input := schema.UserMessage(inputMsg)
	agent, err := NewClsAnalyst(context.Background(), []*schema.Message{role, schema.SystemMessage(prompts.ClsSysMsg)}, input)
	if err != nil {
		panic(err)
	}
	output, err := agent.Run(context.Background())
	if err != nil {
		panic(err)
	}
	t.Log(output.Content)
}

// go test -timeout 1160s -run ^TestSelectAnalyst$ investment/ai/agent -v -count=1
func TestSelectAnalyst(t *testing.T) {
	var SelectData = `
	通过我提供的新闻，分析ETF的买入信号
	以下是新闻内容:
	%s
	`
	clsContent, err := spider.GetClsNews()
	if err != nil {
		panic(err)
	}
	content := fmt.Sprintf(SelectData, clsContent)
	agent, err := NewSelectAnalyst(context.Background(), schema.UserMessage(content))
	if err != nil {
		panic(err)
	}
	msg, err := agent.Run(context.Background())
	if err != nil {
		panic(err)
	}
	htmldata, _ := utility.ConvertHtml(msg.Content)
	//fmt.Println(htmldata)
	//utility.SendHtmlMail()
	utility.SendHtmlMail("investment ai", htmldata)
}
