package db

type Tag struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"uniqueIndex"`
}
