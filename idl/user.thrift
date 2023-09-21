namespace go user
// 用户服务 (User Service)：处理用户注册、登录、身份验证和管理用户信息的服务。

include "common.thrift"

struct User{
    optional i64 uid
    string name
    string email
    string qq
}

struct UserLoginReq {
    string username (vt.min_size = "3", vt.max_size = "32")
    string password (vt.min_size = "3", vt.max_size = "32")
}

struct UserLoginResp {
    i64 uid
    string token
}

struct UserRegisterReq {
    string username (vt.min_size = "3", vt.max_size = "32")
    string nickname (vt.min_size = "3", vt.max_size = "32")
    string password (vt.min_size = "3", vt.max_size = "32")
}

struct UserRegisterResp {
    i64 uid
    string token
}

struct UserLogoutReq{
    string token
}

struct UserDeleteReq{
    string username (vt.min_size = "3", vt.max_size = "32")
    string password (vt.min_size = "3", vt.max_size = "32")
}

service UserService {
    UserLoginResp Login(UserLoginReq request) throws(common.RequestException qe) (api.post="/user/login");
    UserRegisterResp Register(UserRegisterReq request) throws(common.RequestException qe) (api.post="/user/register");
    common.EmptyStruct Logout(UserLogoutReq request) throws(common.RequestException qe) (api.delete="/user/logout");
    common.EmptyStruct Delete(UserDeleteReq request) throws(common.RequestException qe)
}