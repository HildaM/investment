package spider

import (
	"errors"
	"investment/utility"
	"time"
)

// GetClsNews 获取财联社首页电报数据数据
func GetYahooNews() (result string, e error) {
	browser := NewBrowser()
	page := browser.NewMobilePage()
	defer page.Close()
	defer browser.Close()

	page.Navigate("https://finance.yahoo.com/")
	page.MustWaitLoad()
	time.Sleep(1 * time.Second)
	autoScroll(page.Page, 8)
	page.WaitLoad()
	time.Sleep(10 * time.Second)
	quoteData := page.MustElementX(`//*[@id="nimbus-app"]/section/section/section/article/div[1]/section[1]/ul`).MustHTML()
	bodydata, err := page.MustElementX(`//*[@id="nimbus-app"]/section/section/section/article/section[2]/div`).Text()
	if err != nil {
		e = err
		return
	}
	imgPattern := `!\[.*?\]\(.*?\)|<img.*?>`
	linkPattern := `\(https:\/\/.*?\)`
	if len(bodydata) < 500 {
		e = errors.New("抓取数据长度小于 500")
		return
	}

	qdata, err := utility.ConvertMarkDown(quoteData, []string{imgPattern, linkPattern})
	if err != nil {
		e = err
		return
	}
	result = qdata + "\n" + bodydata
	return
}
