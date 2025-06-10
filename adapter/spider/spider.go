package spider

import (
	"investment/adapter/spider/stealth"
	"time"

	"github.com/8treenet/freedom"
	"github.com/go-rod/rod"

	//"github.com/go-rod/rod/lib/defaults"
	"github.com/go-rod/rod/lib/devices"
	"github.com/go-rod/rod/lib/proto"
)

func init() {

}

func NewBrowser() *Browser {
	//defaults.Show = true //关闭无头模式
	browser := rod.New().MustConnect()
	result := &Browser{
		beginTime: time.Now().Unix(),
		browser:   browser,
	}
	return result
}

type Browser struct {
	beginTime int64
	browser   *rod.Browser
	imageDir  string
	proxyAddr string
	userDir   string
}

func (browser *Browser) Close() {
	freedom.Logger().Infof("Browser close  userDir %v, proxyAddr %v", browser.userDir, browser.proxyAddr)
	browser.browser.MustClose()
}

func (browser *Browser) HijackRequests() {
	req := browser.browser.HijackRequests()
	req.MustAdd("*", func(ctx *rod.Hijack) {
		if ctx.Request.Type() == proto.NetworkResourceTypeImage {
			ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
			return
		}
		ctx.ContinueRequest(&proto.FetchContinueRequest{})
	})
	go func() {
		req.Run()
		freedom.Logger().Debug("Browser HijackRequests 结束")
	}()
}

func (browser *Browser) NewMobilePage() *Page {
	page := stealth.MustStealtPage(browser.browser).Timeout(time.Minute * 10)
	page.MustEmulate(devices.IPhoneX)
	return &Page{
		Page:    page,
		browser: browser,
	}
}

func (browser *Browser) NewPage() *Page {
	page := stealth.MustStealtPage(browser.browser).Timeout(time.Minute * 10)
	return &Page{
		Page:    page,
		browser: browser,
	}
}

type Page struct {
	*rod.Page
	browser *Browser
}
