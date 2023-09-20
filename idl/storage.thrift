namespace go storage
// 存储服务 (Storage Service)：负责存储和管理博客文章中的图片和多媒体文件。

include "common.thrift"

struct StorageUploadReq{
    binary data
    2: string type
}

struct StorageUploadResp{
    string url
}

struct StorageDeleteReq{

}
struct StorageDeleteResp{

}
struct StorageManageReq{

}
struct StorageManageResp{

}

service StorageService {
    StorageUploadResp Upload(StorageUploadReq request) throws(common.RequestException qe) (api.post="storage/upload")
}