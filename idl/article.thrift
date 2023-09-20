namespace go article
// 文章服务 (Article Service)：负责管理博客文章的创建、编辑、发布和删除，以及文章的相关操作。

include "common.thrift"

struct articleEdit{
    string title
    string introduce
    string content
    string image
    list<string> tags
    optional i32 category
    bool show
}

struct article{
   i32 aid
   string uri
   string title
   string introduce
   string content
   string image
   list<string> tags
   optional i32 category
   bool show
   i32 uv
   i32 pv
   i32 createTime
   i32 updateTime
}

struct ArticleCreateReq{
    i32 userid
    articleEdit data
}

struct ArticleCreateResp{
    i32 aid
    string uri
}

struct ArticleEditReq{
    i32 aid
    string uri
    articleEdit data
}

struct ArticleDeleteReq{
    i32 aid
    string uri
    optional bool isDelete
}

struct ArticleReq{
    i32 aid
    string uri
}

struct ArticleResp{
    article article
}

struct ArticlesReq{
    i32 pageSize (vt.le = "20")
    i32 pageNum
    i32 category
    i32 tag
}

struct ArticlesResp{
    list<article> articles
}

struct Category {
    i32 id
    string name
    bool showOnHome
    list<Category> categorys
}

struct CategorysResp {
    list<Category> categorys
}

struct ArticleIdUriReq{
    i32 id
    string uri
}

struct ArticleIdUriResp{
    i32 id
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