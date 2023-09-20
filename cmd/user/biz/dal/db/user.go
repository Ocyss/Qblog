package db

import (
	"context"
	"errors"
	"time"

	"github.com/Ocyss/Qblog/pkg/constants"
	"github.com/Ocyss/Qblog/pkg/errno"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"uniqueIndex"`
	Password  string
	Nickname  string
	Role      constants.Role
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	if u.ID == 1 {
		tx.Model(u).Update("role", constants.SuperAdmin)
	}
	return
}

func CreateUser(ctx context.Context, user *User) (uint, error) {
	err := db.WithContext(ctx).Create(user).Error
	return user.ID, err
}

func GetUserByUserName(ctx context.Context, username string) (data *User, err error) {
	data = &User{Username: username}
	err = db.WithContext(ctx).Where("username = ?", username).Take(data).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errno.UserNotExist
	}
	return
}

func UpdatePasswordHash(ctx context.Context, user *User, hash string) error {
	return db.WithContext(ctx).Where(user).Update("password", hash).Error
}

func DeleteUser(ctx context.Context, username string, delete bool) error {
	tx := db.WithContext(ctx).Where("username = ?", username)
	if delete {
		tx.Unscoped()
	}
	return tx.Delete(&User{}).Error
}
