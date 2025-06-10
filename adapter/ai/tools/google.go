package tools

import (
	"context"
	"fmt"
	"investment/utility"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
)

type GooglesSearchResult struct {
	Query string `json:"query" jsonschema_description:"This parameter is used to enter the keywords to be queried, such as: 'Nasdaq Index'."`
}

type GooglesSearchResponse struct {
	Result string `json:"result" jsonschema_description:"Return the results of the queried content."`
}

type GooglesSearch struct {
	Name string
	Desc string
}

func NewGooglesSearch() (tool.InvokableTool, error) {
	rs := &GooglesSearch{
		Name: "google_search",
		Desc: "This is a Google search tool. Enter the information to be queried through the 'query' parameter. For example, if you enter 'Nasdaq Index', relevant results will be returned.",
	}

	tools, err := utils.InferTool(rs.Name, rs.Desc, rs.GetContent)
	if err != nil {
		return nil, fmt.Errorf("failed to infer tool: %w", err)
	}
	return tools, nil
}

func (rs *GooglesSearch) GetContent(ctx context.Context, request *GooglesSearchResult) (result *GooglesSearchResponse, e error) {
	//https://www.google.com/search?q=%E7%BA%B3%E6%96%AF%E8%BE%BE%E5%85%8B%E6%8C%87%E6%95%B0
	urladdr := fmt.Sprintf("https://www.google.com/search?q=%s", request.Query)
	content, e := utility.NewDefaultCrwl().Run(urladdr)
	if e != nil {
		return
	}
	result = &GooglesSearchResponse{Result: content}
	return
}
