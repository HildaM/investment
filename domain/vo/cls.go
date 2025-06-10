package vo

//财联社

type ClsDepthArticle struct {
	ArticleID int    `json:"id"`
	Ctime     int    `json:"ctime"`
	SortScore int    `json:"sort_score"`
	Title     string `json:"title"`
	Brief     string `json:"brief"`
}

type ClsDepthArticleExt struct {
	URL      string `json:"url"`
	Brief    string `json:"brief"`
	Datetime string `json:"datetime"`
}

type ClsDepthArticleContent struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}
