namespace go statistics
// 统计服务 (Statistics Service)：收集和分析博客的访问量、用户行为和其他统计信息，用于监控和优化博客性能。

include "common.thrift"

struct StatisticsArticleReq{
    i32 id
    string uri
    optional string ua
}

struct StatisticsArticleResp{
    i32 pv
    i32 uv
}

service StatisticsService {
    StatisticsArticleResp Article(1: StatisticsArticleReq request) throws(common.RequestException qe) (api.get="/statistics/article");
}