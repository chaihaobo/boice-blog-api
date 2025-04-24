package entity

type Tag struct {
	BaseEntity
	Name string `gorm:"column:name"`
}

func (Tag) TableName() string {
	return "tag"
}
