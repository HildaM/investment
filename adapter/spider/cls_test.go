package spider

import (
	"encoding/json"
	"fmt"
	"investment/utility"
	"testing"
	"time"
)

// go test -timeout 1160s -run ^TestGetClsDepthList$ investment/adapter/spider -v -count=1
func TestGetClsDepthList(t *testing.T) {
	text, err := GetClsDepthList()
	if err != nil {
		panic(err)
	}

	jdata, _ := json.MarshalIndent(text, "", "")
	fmt.Println(string(jdata))
	time.Sleep(5 * time.Second)
}

func TestGetClsDepth(t *testing.T) {
	text, err := GetClsDepth(2003579)
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}

func TestGetClsQuotation(t *testing.T) {
	list, err := GetClsQuotation()
	if err != nil {
		panic(err)
	}
	t.Log(utility.Jsonout(list))
}

func TestGetSearchPage(t *testing.T) {
	t.Log(GetSearchPage("纳斯达克"))
}

func TestGetDetail(t *testing.T) {
	t.Log(GetDetail("https://www.cls.cn/detail/2053873"))
}
