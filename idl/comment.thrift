namespace go comment
// 评论服务 (Comment Service)：处理博客文章的评论，包括评论的创建以及评论的管理。

include "common.thrift"
include "user.thrift"

struct Comment {
    user.User user
    string content
    optional i64 reply
    optional i64 father
}

struct CommentCreateReq{
    i64 aid
    Comment data
}

struct CommentListReq{
    i64 aid
}

struct CommentListResp{
    list<Comment> data
}

service CommentService {
    common.EmptyStruct Create(CommentCreateReq request) throws(common.RequestException qe) (api.get="/comment/create");
    CommentListResp Get(CommentListReq request) throws(common.RequestException qe) (api.get="/comment/get");
}