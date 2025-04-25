package entity

import "time"

type Article struct {
	BaseEntity
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	Content     string `gorm:"column:content"`
	Tags        []*Tag `gorm:"-"`
}

func (Article) TableName() string {
	return "article"
}

func (a *Article) Edit(modified *Article) {
	a.Title = modified.Title
	a.Description = modified.Description
	a.Content = modified.Content
	a.Tags = modified.Tags
	a.UpdatedAt = time.Now()

}
