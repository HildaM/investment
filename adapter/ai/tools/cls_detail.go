package tools

import (
	"context"
	"fmt"
	"investment/adapter/spider"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
)

type ClsDetailResult struct {
	URL string `json:"url" jsonschema_description:"该参数可以输入财联社相关文章的URL地址,如:https://www.cls.cn/detail/2004091."`
}

type ClsDetailResponse struct {
	Result string `json:"result" jsonschema_description:"The content of the website."`
}

type ClsDetail struct {
	Name string
	Desc string
}

func NewClsDetail() (tool.InvokableTool, error) {
	page := &ClsDetail{
		Name: "cls_detail",
		Desc: `这是一个财联社详情页的爬虫，输入详情页的URL。返回该详情页的内容。
Examples: 
Fill in the URL parameter with "https://www.cls.cn/detail/1979335", and the content of this article will be returned.`,
	}

	tools, err := utils.InferTool(page.Name, page.Desc, page.GetContent)
	if err != nil {
		return nil, fmt.Errorf("failed to infer tool: %w", err)
	}
	return tools, nil
}

func (page *ClsDetail) GetContent(ctx context.Context, request *ClsDetailResult) (result *ClsDetailResponse, e error) {
	content, e := spider.GetDetail(request.URL)
	if e != nil {
		return
	}
	result = &ClsDetailResponse{Result: content}
	return
}
