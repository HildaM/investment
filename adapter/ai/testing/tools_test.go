package testing

import (
	"context"
	"encoding/json"
	"investment/adapter/ai/tools"
	"testing"
)

// go test -timeout 60s -run ^TestPage$ investment/ai/testing -v -count=1
func TestPage(t *testing.T) {
	page, err := tools.NewPage()
	if err != nil {
		panic(err)
	}

	request := &tools.PageResult{
		URL: "https://www.cls.cn",
	}
	jdata, _ := json.Marshal(request)
	output, err := page.InvokableRun(context.Background(), string(jdata))
	if err != nil {
		panic(err)
	}
	t.Log(output)
}

// go test -timeout 60s -run ^TestGoogle$ investment/ai/testing -v -count=1
func TestGoogle(t *testing.T) {
	page, err := tools.NewGooglesSearch()
	if err != nil {
		panic(err)
	}

	request := &tools.GooglesSearchResult{
		Query: "上证指数",
	}
	jdata, _ := json.Marshal(request)
	output, err := page.InvokableRun(context.Background(), string(jdata))
	if err != nil {
		panic(err)
	}
	t.Log(output)
}

// go test -timeout 160s -run ^TestReuters$ investment/ai/testing -v -count=1
func TestReuters(t *testing.T) {
	page, err := tools.NewReutersSearch()
	if err != nil {
		panic(err)
	}

	request := &tools.ReutersSearchResult{
		Query: "纳斯达克",
	}
	jdata, _ := json.Marshal(request)
	output, err := page.InvokableRun(context.Background(), string(jdata))
	if err != nil {
		panic(err)
	}
	t.Log(output)
}

func TestClsTelegramSearch(t *testing.T) {
	page, err := tools.NewClsTelegramSearch()
	if err != nil {
		panic(err)
	}

	request := &tools.ClsTelegramSearchResult{
		Keyword: "纳斯达克",
	}
	jdata, _ := json.Marshal(request)
	output, err := page.InvokableRun(context.Background(), string(jdata))
	if err != nil {
		panic(err)
	}
	t.Log(output)
}

func TestClsDepthSearch(t *testing.T) {
	page, err := tools.NewClsDepthSearch()
	if err != nil {
		panic(err)
	}

	request := &tools.ClsDepthSearchResult{
		Keyword: "纳斯达克",
	}
	jdata, _ := json.Marshal(request)
	output, err := page.InvokableRun(context.Background(), string(jdata))
	if err != nil {
		panic(err)
	}
	t.Log(output)
}

func TestClsDetailSearch(t *testing.T) {
	page, err := tools.NewClsDetail()
	if err != nil {
		panic(err)
	}

	request := &tools.ClsDetailResult{
		URL: "https://www.cls.cn/detail/2004091",
	}
	jdata, _ := json.Marshal(request)
	output, err := page.InvokableRun(context.Background(), string(jdata))
	if err != nil {
		panic(err)
	}
	t.Log(output)
}
