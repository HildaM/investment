package custom

import "time"

type ClsDepthArticle struct {
	ArticleID int       `gorm:"primaryKey;column:articleId"`
	Ctime     time.Time `gorm:"column:ctime"`   // 创建时间
	Created   time.Time `gorm:"column:created"` // 创建时间
	SortScore int       `gorm:"column:sortScore"`
	Title     string    `gorm:"column:title"`
	Brief     string    `gorm:"column:brief"`
	Content   string    `gorm:"column:content"`
}
