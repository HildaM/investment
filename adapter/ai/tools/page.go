package tools

import (
	"context"
	"fmt"
	"investment/utility"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
)

type PageResult struct {
	URL string `json:"url" jsonschema_description:"Website address; URL (Uniform Resource Locator)."`
}

type PageResponse struct {
	Result string `json:"result" jsonschema_description:"The content of the website."`
}

type Page struct {
	Name string
	Desc string
}

func NewPage() (tool.InvokableTool, error) {
	page := &Page{
		Name: "crawler",
		Desc: `This tool supports entering any https URL and returning the content of the site.
Examples: 
1. Enter the parameter https://www.reuters.com, and the content of the Reuters homepage will be returned.
2. For https://www.cls.cn/detail/1993685, the content of a certain article on Caixin will be returned.
Note that any URL is supported.`,
	}

	tools, err := utils.InferTool(page.Name, page.Desc, page.GetContent)
	if err != nil {
		return nil, fmt.Errorf("failed to infer tool: %w", err)
	}
	return tools, nil
}

func (page *Page) GetContent(ctx context.Context, request *PageResult) (result *PageResponse, e error) {
	content, e := utility.NewDefaultCrwl().Run(request.URL)
	if e != nil {
		return
	}
	result = &PageResponse{Result: content}
	return
}
