package db

import (
	"context"

	"github.com/Ocyss/Qblog/pkg/constants"
)

type Category struct {
	ID         uint   `gorm:"primarykey"`
	Name       string `gorm:"index"`
	Level      *constants.Role
	Sort       uint
	ShowOnHome bool
	Articles   []Article
	CategoryID *uint
	Categorys  []Category `gorm:"foreignkey:CategoryID"`
}

func CreateCategory(ctx context.Context, data *Category) error {
	return db.Create(data).Error
}
