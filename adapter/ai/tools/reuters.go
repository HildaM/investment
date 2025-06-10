package tools

import (
	"context"
	"fmt"
	"investment/utility"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
)

type ReutersSearchResult struct {
	Query string `json:"query" jsonschema_description:"This parameter is used to enter the keywords to be queried, such as: 'Nasdaq Index'."`
}

type ReutersSearchResponse struct {
	Result string `json:"result" jsonschema_description:"Return the results of the queried content."`
}

type ReutersSearch struct {
	Name string
	Desc string
}

func NewReutersSearch() (tool.InvokableTool, error) {
	rs := &ReutersSearch{
		Name: "cls_crawler",
		Desc: "这是一个财联社的爬虫。通过“query”参数输入要查询的信息，例如“纳斯达克指数”，将会返回相关咨询。",
	}

	tools, err := utils.InferTool(rs.Name, rs.Desc, rs.GetContent)
	if err != nil {
		return nil, fmt.Errorf("failed to infer tool: %w", err)
	}
	return tools, nil
}

func (rs *ReutersSearch) GetContent(ctx context.Context, request *ReutersSearchResult) (result *ReutersSearchResponse, e error) {
	//https://www.reuters.com/site-search/?query=NASDAQ&offset=20
	urladdr := fmt.Sprintf("https://www.cls.cn/searchPage?keyword=%s&type=depth", request.Query)
	content, e := utility.NewDefaultCrwl().Run(urladdr)
	if e != nil {
		return
	}
	result = &ReutersSearchResponse{Result: content}
	return
}
