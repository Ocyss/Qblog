namespace go article
// 文章服务 (Article Service)：负责管理博客文章的创建、编辑、发布和删除，以及文章的相关操作。

include "common.thrift"

struct articleEdit{
    string title
    string introduce
    string content
    string image
    list<string> tags
    optional i64 category
    bool show
}

struct article{
   i64 aid
   string uri
   string title
   string introduce
   string content
   string image
   list<string> tags
   optional i64 category
   bool show
   i64 uv
   i64 pv
   i64 createTime
   i64 updateTime
}

struct ArticleCreateReq{
    i64 userid
    articleEdit data
}

struct ArticleCreateResp{
    i64 aid
    string uri
}

struct ArticleEditReq{
    i64 aid
    string uri
    articleEdit data
}

struct ArticleDeleteReq{
    i64 aid
    string uri
    optional bool isDelete
}

struct ArticleReq{
    i64 aid
    string uri
}

struct ArticleResp{
    article article
}

struct ArticlesReq{
    i64 pageSize (vt.le = "20")
    i64 pageNum
    i64 category
    i64 tag
}

struct ArticlesResp{
    list<article> articles
}

struct Category {
    i64 id
    string name
    bool showOnHome
    list<Category> categorys
}

struct CategorysResp {
    list<Category> categorys
}

struct ArticleIdUriReq{
    i64 id
    string uri
}

struct ArticleIdUriResp{
    i64 id
    string uri
}

service ArticleService {
    ArticleCreateResp Create(ArticleCreateReq request) throws(common.RequestException qe) (api.post="/article/create");
    common.EmptyStruct Edit(ArticleEditReq request) throws(common.RequestException qe) (api.put="/article/edit");
    common.EmptyStruct Delete(ArticleDeleteReq request) throws(common.RequestException qe) (api.delete="/article/delete");
    ArticleResp  Get(ArticleReq request) throws(common.RequestException qe) (api.get="/article/get");
    ArticlesResp GetList(ArticlesReq request) throws(common.RequestException qe) (api.get="/article/getlist");
    CategorysResp GetCategoryList() throws(common.RequestException qe) (api.get="/article/categorys");
    ArticleIdUriResp GetIdUri(ArticleIdUriReq request) throws(common.RequestException qe)
}