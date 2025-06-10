package tools

import (
	"context"
	"fmt"
	"investment/adapter/spider"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
)

type ClsTelegramSearchResult struct {
	Keyword string `json:"keyword" jsonschema_description:"此参数用于输入要查询的关键词，例如：“纳斯达克指数”。"`
}

type ClsTelegramSearchResponse struct {
	Result string `json:"result" jsonschema_description:"返回查询内容的结果。"`
}

type ClsTelegramSearch struct {
	Name string
	Desc string
}

func NewClsTelegramSearch() (tool.InvokableTool, error) {
	rs := &ClsTelegramSearch{
		Name: "cls_telegram",
		Desc: "这是一个财联社电报的爬虫。通过“keyword”参数输入要查询的关键词，例如“纳斯达克指数”，将会返回相关电报。",
	}

	tools, err := utils.InferTool(rs.Name, rs.Desc, rs.GetContent)
	if err != nil {
		return nil, fmt.Errorf("failed to infer tool: %w", err)
	}
	return tools, nil
}

func (rs *ClsTelegramSearch) GetContent(ctx context.Context, request *ClsTelegramSearchResult) (result *ClsTelegramSearchResponse, e error) {
	content, e := spider.GetSearchPage(request.Keyword)
	if e != nil {
		return
	}
	result = &ClsTelegramSearchResponse{Result: content}
	return
}
