namespace go search
// 搜索服务 (Search Service)：提供全文搜索功能，使用户可以搜索博客文章、标签和评论等内容。

include "common.thrift"
include "article.thrift"

struct SearchArticleReq {
    string content
}

struct SearchArticleResp{
    list<article.article> articles
}

service SearchService {
    SearchArticleResp FindArticle(SearchArticleReq request) throws(common.RequestException qe)
}