namespace go notification
// 通知服务 (Notification Service)：处理博客相关的通知，如新评论、新文章、关注用户等通知的生成和传送。

include "common.thrift"

struct NotificationSendReq {
    string title
    string body
}

service NotificationService {
    common.EmptyStruct Send(NotificationSendReq request) throws(common.RequestException qe)
}