namespace go comment
// 评论服务 (Comment Service)：处理博客文章的评论，包括评论的创建以及评论的管理。

include "common.thrift"

struct CommentCreateReq{
    i64 aid
    string uri
    string content
    optional i64 reply
}

service CommentService {
    common.EmptyStruct Create(CommentCreateReq request) throws(common.RequestException qe) (api.get="/comment/create");
}