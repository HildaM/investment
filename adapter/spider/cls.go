package spider

import (
	"encoding/json"
	"fmt"
	"investment/domain/vo"
	"investment/utility"
	"strings"
	"sync"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

var scrollToBottomJS = `
() => {
    const height = Math.max(
        document.body.scrollHeight,
        document.documentElement.scrollHeight,
        document.body.offsetHeight,
        document.documentElement.offsetHeight
    );
    window.scrollTo(0, height);
    return height;
}
`

func autoScroll(page *rod.Page, maxCount int) {
	lastHeight := 0
	retry := 0
	for i := 0; i < maxCount; i++ {
		// 执行滚动并获取最新高度
		newHeight := page.MustEval(scrollToBottomJS).Int()

		// 高度未变化时退出
		if newHeight == lastHeight {
			retry++
			if retry > 2 { // 允许最多重试2次
				break
			}
		} else {
			retry = 0
			lastHeight = newHeight
		}

		// 等待内容加载
		//page.WaitStable(time.Second) // 更智能的等待
		page.WaitLoad()
		time.Sleep(4 * time.Second)
	}
}

// GetClsNews 获取财联社首页电报数据数据
func GetClsNews() (result string, e error) {
	browser := NewBrowser()
	page := browser.NewMobilePage()
	defer page.Close()
	defer browser.Close()

	page.Navigate("https://m.cls.cn/")
	time.Sleep(3 * time.Second)
	autoScroll(page.Page, 8)
	time.Sleep(3 * time.Second)
	htmldata := page.MustElementX(`//*[@id="__next"]/div/div/div[3]/div/div`).MustHTML()

	// 删除特定的图片链接
	pattern := `!\[\]\(https://cdnjs\.cailianpress\.com/images/msite/(comment_btn\.png|index_calendar\.png)\)[0-9]*`
	return utility.ConvertMarkDown(htmldata, []string{pattern})
}

// https://m.cls.cn/depth 获取财联社深度列表
func GetClsDepthList() (result []vo.ClsDepthArticle, e error) {
	browser := NewBrowser()
	page := browser.NewPage()
	page.Page.MustSetExtraHeaders("referer", "https://www.cls.cn")
	defer page.Close()
	defer browser.Close()

	req := page.HijackRequests()
	var mutex sync.Mutex
	req.Add("https://www.cls.cn/v3/depth/home/assembled/1000?*", proto.NetworkResourceTypeXHR, func(ctx *rod.Hijack) {
		var data struct {
			Errno int `json:"errno"`
			Data  struct {
				DepthList  []vo.ClsDepthArticle `json:"depth_list"`
				TopArticle []vo.ClsDepthArticle `json:"top_article"`
			} `json:"data"`
		}
		ctx.MustLoadResponse()
		body := []byte(ctx.Response.Body())
		if err := json.Unmarshal(body, &data); err != nil {
			return
		}

		mutex.Lock()
		result = append(result, data.Data.TopArticle...)
		result = append(result, data.Data.DepthList...)
		mutex.Unlock()
	})
	req.Add("https://www.cls.cn/v3/depth/list/1000?*", proto.NetworkResourceTypeXHR, func(ctx *rod.Hijack) {
		var data struct {
			Errno int                  `json:"errno"`
			Data  []vo.ClsDepthArticle `json:"data"`
		}
		ctx.MustLoadResponse()
		body := []byte(ctx.Response.Body())
		if err := json.Unmarshal(body, &data); err != nil {
			return
		}

		mutex.Lock()
		result = append(result, data.Data...)
		mutex.Unlock()
	})
	go req.Run()

	page.Navigate("https://www.cls.cn/depth?id=1000")
	page.MustWaitLoad()
	time.Sleep(2 * time.Second)

	for i := 0; i < 2; i++ {
		fmt.Println("Cls depth scroll.")
		autoScroll(page.Page, 2)
		next := page.MustElementX(`//*[@id="__next"]/div/div[2]/div[2]/div[1]/div[3]/div[36]`)
		next.MustClick()
		time.Sleep(2 * time.Second)
	}
	time.Sleep(5 * time.Second)
	return
}

// https://m.cls.cn/depth/2002105 获取财联社深度详情
func GetClsDepth(articleId int) (result vo.ClsDepthArticleContent, e error) {
	browser := NewBrowser()
	page := browser.NewMobilePage()
	defer page.Close()
	defer browser.Close()

	page.Page.MustSetExtraHeaders("referer", "https://m.cls.cn/depth")
	page.Navigate(fmt.Sprintf("https://m.cls.cn/depth/%d", articleId))
	page.MustWaitLoad()
	htmldata := page.MustElement(`#__next > div > div > div.article`).MustHTML()
	// 删除特定的图片链接
	pattern := `!\[.*?\]\(.*?\)|<img.*?>`
	time.Sleep(1 * time.Second)

	content, e := utility.ConvertMarkDown(htmldata, []string{pattern})
	if e != nil {
		return
	}
	return vo.ClsDepthArticleContent{
		ID:      articleId,
		Content: content,
	}, nil
}

// GetClsQuotation
// https://www.cls.cn/quotation 获取财联社指数
func GetClsQuotation() (result []any, e error) {
	browser := NewBrowser()
	page := browser.NewPage()
	defer page.Close()
	defer browser.Close()

	req := page.HijackRequests()
	var mutex sync.Mutex
	ok := false
	req.Add("https://x-quote.cls.cn/quote/index/home?*", proto.NetworkResourceTypeXHR, func(ctx *rod.Hijack) {
		var data struct {
			Code int `json:"code"`
			Data struct {
				List []any `json:"index_quote"`
			} `json:"data"`
		}
		ctx.MustLoadResponse()
		body := []byte(ctx.Response.Body())
		if err := json.Unmarshal(body, &data); err != nil {
			return
		}

		mutex.Lock()
		result = data.Data.List
		ok = true
		mutex.Unlock()
	})
	go req.Run()

	page.Page.MustSetExtraHeaders("referer", "https://www.cls.cn/finance")
	page.Navigate("https://www.cls.cn/quotation")
	page.MustWaitLoad()

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		over := false
		mutex.Lock()
		over = ok
		mutex.Unlock()

		if over {
			break
		}
	}
	return
}

// GetSearchPage
// https://www.cls.cn/searchPage?keyword=%E7%BE%8E%E8%82%A1&type=telegram
func GetSearchPage(keyword string) (result string, e error) {
	browser := NewBrowser()
	defer browser.Close()
	page := browser.NewPage()
	defer page.Close()

	page.Navigate(fmt.Sprintf("https://www.cls.cn/searchPage?keyword=%s&type=telegram", keyword))
	fmt.Println("搜索进行中...")
	time.Sleep(5 * time.Second)
	body, e := page.MustElementX(`//*[@id="__next"]/div/div[2]/div[1]/div[3]/div[1]`).HTML()
	if e != nil {
		return
	}

	// 删除特定的图片链接
	pattern := `!\[.*?\]\(.*?\)|<img.*?>`
	data, e := utility.ConvertMarkDown(body, []string{pattern})
	if e != nil {
		return
	}
	result = strings.ReplaceAll(data, "(/detail/", "(https://www.cls.cn/detail/")
	return
}

// GetDetail
func GetDetail(url string) (result string, e error) {
	browser := NewBrowser()
	defer browser.Close()
	page := browser.NewPage()
	defer page.Close()

	page.Navigate(url)
	fmt.Println("查看详情进行中...")
	time.Sleep(5 * time.Second)

	body, e := page.MustElementX(`//*[@id="__next"]/div/div[2]/div[2]/div[1]`).HTML()
	if e != nil {
		return
	}

	// 删除特定的图片链接
	pattern := `!\[.*?\]\(.*?\)|<img.*?>`
	time.Sleep(1 * time.Second)
	data, e := utility.ConvertMarkDown(body, []string{pattern})
	if e != nil {
		return
	}
	result = strings.ReplaceAll(data, "(/detail/", "(https://www.cls.cn/detail/")
	return
}
