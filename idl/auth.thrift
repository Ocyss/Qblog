namespace go auth
// 身份验证与授权服务 (Authentication and Authorization Service)：负责用户身份验证和授权，以确保只有授权用户可以访问博客的敏感信息和功能。

include "common.thrift"

struct AuthTokenCheckReq{
    string token
}

service AuthService {
    common.EmptyStruct TokenCheck(AuthTokenCheckReq request) throws(common.RequestException qe)
}