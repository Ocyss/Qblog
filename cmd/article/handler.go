package main

import (
	"context"

	"github.com/Ocyss/Qblog/kitex_gen/article"
	"github.com/Ocyss/Qblog/kitex_gen/common"

	"github.com/Ocyss/Qblog/cmd/article/biz/service"
)

// ArticleServiceImpl implements the last service interface defined in the IDL.
type ArticleServiceImpl struct{}

// Create implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) Create(ctx context.Context, request *article.ArticleCreateReq) (resp *article.ArticleCreateResp, err error) {
	articleService := service.NewArticleService(ctx)
	resp, err = articleService.Create(request)
	return
}

// Edit implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) Edit(ctx context.Context, request *article.ArticleEditReq) (resp *common.EmptyStruct, err error) {
	articleService := service.NewArticleService(ctx)
	resp, err = articleService.Edit(request)
	return
}

// Delete implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) Delete(ctx context.Context, request *article.ArticleDeleteReq) (resp *common.EmptyStruct, err error) {
	articleService := service.NewArticleService(ctx)
	resp, err = articleService.Delete(request)
	return
}

// Get implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) Get(ctx context.Context, request *article.ArticleReq) (resp *article.ArticleResp, err error) {
	articleService := service.NewArticleService(ctx)
	resp, err = articleService.Get(request)
	return
}

// GetList implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) GetList(ctx context.Context, request *article.ArticlesReq) (resp *article.ArticlesResp, err error) {
	articleService := service.NewArticleService(ctx)
	resp, err = articleService.GetList(request)
	return
}

// GetCategoryList implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) GetCategoryList(ctx context.Context) (resp *article.CategorysResp, err error) {
	articleService := service.NewArticleService(ctx)
	resp, err = articleService.GetCategoryList()
	return
}

// GetIdUri implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) GetIdUri(ctx context.Context, request *article.ArticleIdUriReq) (resp *article.ArticleIdUriResp, err error) {
	articleService := service.NewArticleService(ctx)
	resp, err = articleService.GetIdUri(request)
	return
}

// CheckIdUri implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) CheckIdUri(ctx context.Context, request *article.ArticleIdUriReq) (err error) {
	articleService := service.NewArticleService(ctx)
	err = articleService.CheckIdUri(request)
	return
}
