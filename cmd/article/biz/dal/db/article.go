package db

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/Ocyss/Qblog/pkg/errno"

	"github.com/gosimple/slug"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	UserId     uint   `gorm:"index"`
	Uri        string `gorm:"uniqueIndex"`
	Title      string
	Introduce  string
	Content    string
	Image      string
	Tags       []Tag `gorm:"many2many:article_tags;constraint:OnDelete:CASCADE;"`
	CategoryID *uint `gorm:"index"`
	Show       bool
}

func (a *Article) AfterCreate(tx *gorm.DB) (err error) {
	if a.Uri == "" {
		var buf strings.Builder
		buf.WriteString(strconv.Itoa(int(a.ID)))
		buf.WriteByte('-')
		buf.WriteString(slug.Make(a.Title))
		tx.Model(a).Update("uri", buf.String())
	}
	return
}

func CreateArticle(ctx context.Context, data *Article, tags []string) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := toStrTags(data, tags); err != nil {
			return err
		}
		return tx.Create(data).Error
	})
}

func UpdateArticle(ctx context.Context, aid uint, uri string, data Article, tags []string) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where(&Article{Model: id(aid), Uri: uri}).Select("id").Take(&data).Error; err != nil {
			return err
		}
		if err := toStrTags(&data, tags); err != nil {
			return err
		}
		if err := tx.Model(&Article{Model: id(data.ID)}).Association("Tags").Clear(); err != nil {
			return err
		}
		return tx.Updates(&data).Error
	})
}

func DeleteArticle(ctx context.Context, aid uint, uri string, delete bool) error {
	tx := db.WithContext(ctx).Where(&Article{Model: id(aid), Uri: uri})
	if delete {
		tx.Unscoped()
	}
	return tx.Delete(&Article{}).Error
}

func GetArticle(ctx context.Context, aid uint, uri string) (data *Article, err error) {
	data = &Article{}
	err = db.WithContext(ctx).Where(&Article{Model: id(aid), Uri: uri}).Take(data).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errno.ArticleNotExist
	}
	return
}

func GetArticleList(ctx context.Context, pageSize, pageNum, categoryID, tagID int) (data []*Article, err error) {
	data = make([]*Article, 0, pageSize)
	offset := (pageNum - 1) * pageSize
	tx := db.WithContext(ctx).Limit(pageSize).Offset(offset)
	if categoryID != 0 {
		tx = tx.Where("category_id = ?", categoryID)
	}
	if tagID != 0 {
		tx = tx.Joins("INNER JOIN article_tags ON articles.id = article_tags.article_id").
			Where("article_tags.tag_id = ?", tagID)
	}
	err = tx.Order("created_at DESC").Find(&data).Error
	return
}

func toStrTags(data *Article, tags []string) error {
	if n := len(tags); n > 0 {
		tagChan := make(chan Tag)
		errChan := make(chan error)
		data.Tags = make([]Tag, 0, len(tags))

		for i := 0; i < n; i++ {
			go func(index int, tagName string) {
				tag := Tag{Name: tagName}
				if err := db.Where("name = ?", tagName).FirstOrCreate(&tag).Error; err != nil {
					errChan <- err
					return
				}
				tagChan <- tag
			}(i, tags[i])
		}
		for i := 0; i < n; i++ {
			select {
			case tag := <-tagChan:
				data.Tags = append(data.Tags, tag)
			case err := <-errChan:
				return err
			}
		}
	}
	return nil
}

func GetIdUri(ctx context.Context, aid uint, uri string) (uint, string, error) {
	var data Article
	err := db.WithContext(ctx).Where(&Article{Model: id(aid), Uri: uri}).Take(&data).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errno.ArticleNotExist
	}
	return data.ID, data.Uri, err
}
