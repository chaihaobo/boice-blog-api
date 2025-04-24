package entity

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
