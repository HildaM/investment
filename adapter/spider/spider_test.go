package spider

import (
	"fmt"
	"testing"
	"time"
)

func TestGetClsIndex(t *testing.T) {
	text, err := GetClsNews()
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
	time.Sleep(5 * time.Second)
}

// go test -timeout 1160s -run ^TestGetYahooNews$ investment/adapter/spider -v -count=1
func TestGetYahooNews(t *testing.T) {
	text, err := GetYahooNews()
	if err != nil {
		panic(err)
	}
	t.Log(text)
}
