package utility

import (
	"fmt"
	"os"
	"os/exec"
)

func NewCrwl(condaEnv, browserFilePath, crawlerFilePath string) *Crwl {
	return &Crwl{condaEnv: condaEnv, browserFilePath: browserFilePath, crawlerFilePath: crawlerFilePath}
}

func NewDefaultCrwl() *Crwl {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	browserFilePath := fmt.Sprintf("%s/.investment/browser.yml", dir)
	crawlerFilePath := fmt.Sprintf("%s/.investment/crawler.yml", dir)
	return NewCrwl("investment", browserFilePath, crawlerFilePath)
}

type Crwl struct {
	condaEnv        string
	browserFilePath string
	crawlerFilePath string
}

// crwl "https://www.so.com/s?ie=utf-8&fr=360sou_newhome&src=360sou_newhome&q=降息" -o markdown -B browser.yml -C crawler.yml
func (crwl *Crwl) Run(url string, formats ...string) (result string, err error) {
	format := "markdown"
	if len(formats) > 0 {
		format = formats[0]
	}

	//cmd := exec.Command("conda", "run", "-n", crwl.condaEnv, "crwl", url, "-o", format, "-B", "browser.yml", "-C", "crawler.yml")
	fmt.Println("conda", "run", "-n", crwl.condaEnv, "crwl", url, "-o", format, "-B", crwl.browserFilePath, "-C", crwl.crawlerFilePath)
	cmd := exec.Command("conda", "run", "-n", crwl.condaEnv, "crwl", url, "-o", format, "-B", crwl.browserFilePath, "-C", crwl.crawlerFilePath)
	output, err := cmd.Output()
	if err != nil {
		return
	}
	return string(output), err
}
