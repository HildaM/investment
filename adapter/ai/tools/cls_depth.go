package tools

import (
	"context"
	"fmt"
	"investment/utility"
	"os"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
)

type ClsDepthSearchResult struct {
	Keyword string `json:"keyword" jsonschema_description:"此参数用于输入要查询的关键词，例如：“纳斯达克指数”。"`
}

type ClsDepthSearchResponse struct {
	Result string `json:"result" jsonschema_description:"返回查询内容的结果。"`
}

type ClsDepthSearch struct {
	Name string
	Desc string
}

func NewClsDepthSearch() (tool.InvokableTool, error) {
	rs := &ClsDepthSearch{
		Name: "cls_depth",
		Desc: "这是一个财联社咨询的爬虫。通过“keyword”参数输入要查询的关键词，例如“纳斯达克指数”，将会返回相关咨询。",
	}

	tools, err := utils.InferTool(rs.Name, rs.Desc, rs.GetContent)
	if err != nil {
		return nil, fmt.Errorf("failed to infer tool: %w", err)
	}
	return tools, nil
}

func (rs *ClsDepthSearch) GetContent(ctx context.Context, request *ClsDepthSearchResult) (result *ClsDepthSearchResponse, e error) {
	//https://www.cls.cn/searchPage?keyword=%E9%99%8D%E6%81%AF&type=Depth
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	browserFilePath := fmt.Sprintf("%s/.investment/browser.yml", dir)
	crawlerFilePath := fmt.Sprintf("%s/.investment/crawler_cls.yml", dir)
	urladdr := fmt.Sprintf("https://www.cls.cn/searchPage?keyword=%s&type=depth", request.Keyword)

	content, e := utility.NewCrwl("investment", browserFilePath, crawlerFilePath).Run(urladdr)
	if e != nil {
		return
	}
	result = &ClsDepthSearchResponse{Result: content}
	return
}
