package utility

import (
	"bytes"
	"regexp"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/yuin/goldmark"
)

func ConvertMarkDown(html string, removePatterns []string) (string, error) {
	markdown, err := htmltomarkdown.ConvertString(html)
	if err != nil {
		return "", err
	}
	for _, v := range removePatterns {
		re := regexp.MustCompile(v)
		markdown = re.ReplaceAllString(markdown, "")
	}

	return markdown, nil
}

func ConvertHtml(markDown string) (string, error) {
	var buf bytes.Buffer
	err := goldmark.Convert([]byte(markDown), &buf)
	if err != nil {
		panic(err)
	}
	return buf.String(), nil
}
