package db

import (
	"context"
	"time"
)

type Comment struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	Aid       uint `gorm:"index"`
	CommentUser
	Show     bool
	Content  string
	ReplyID  *uint
	Reply    *Comment
	FatherID *uint
	Father   *Comment
}

type CommentUser struct {
	Uid   *uint
	Name  string
	Email string
	Qq    string
}

func CreateComment(ctx context.Context, data *Comment) error {
	return db.WithContext(ctx).Create(data).Error
}

func GetComment(ctx context.Context, aid uint) ([]*Comment, error) {
	var data []*Comment
	err := db.WithContext(ctx).Where(&Comment{Aid: aid}).Order("created_at DESC").Limit(20).Find(&data).Error
	return data, err
}
