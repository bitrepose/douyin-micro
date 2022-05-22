namespace go video

include "user.thrift"

struct Video {
    1: required i64        id;               // 视频唯一标识
    2: required user.User  author;           // 视频作者信息
    3: required string     play_url;         // 视频播放地址
    4: required string     cover_url;        // 视频封面地址
    5: required i64        favorite_count;   // 视频的点赞总数
    6: required i64        comment_count;    // 视频的评论总数
    7: required string     title;            // 视频标题
}

struct FeedRequset {
    1: optional i64     latest_time;   // 限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
}

struct FeedResponse {
    1: required i32          status_code;   // 状态码，0-成功，其他值-失败
    2: optional string       status_msg;    // 返回状态描述
    3: list<Video>           video_list;    // 视频列表
    4: optional i64          next_time;     // 本次返回的视频中，发布最早的时间，作为下次请求是的latest_time
}

struct PublishActionRequest {
    1: required i64         user_id;  // 用户id
    2: required list<byte>  data;     // 视频数据
    3: required string      title;    // 视频标题
}

struct PublishActionResponse {
    1: required i32     status_code;   // 状态码，0-成功，其他值-失败
    2: optional string  status_msg;    // 返回状态描述
}

struct PublishListRequest {
    1: required i64     user_id;   // 用户id
}

struct PublishListResponse {
    1: required i32          status_code;   // 状态码，0-成功，其他值-失败
    2: optional string       status_msg;    // 返回状态描述
    3: list<Video>           video_list;    // 用户发布的视频列表
}

struct FavoriteActionRequest {
    1: required i64     user_id;       // 用户id
    3: required i64     video_id;      // 视频id
    4: required i32     action_type;   // 1-点赞，2-取消点赞
}

struct FavoriteActionResponse {
    1: required i32     status_code;   // 状态码，0-成功，其他值-失败
    2: optional string  status_msg;    // 返回状态描述
}

struct FavoriteListRequest {
    1: required i64     user_id;       // 用户id
}

struct FavoriteListResponse {
    1: required i32          status_code;   // 状态码，0-成功，其他值-失败
    2: optional string       status_msg;    // 返回状态描述
    3: required list<Video>  video_list;    // 用户点赞视频列表
}

service VideoService {
    FeedResponse Feed(1: FeedRequset req)
    PublishActionResponse PublishAction(1: PublishActionRequest req)
    PublishListResponse PublishList(1: PublishListRequest req)
    FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest req)
    FavoriteListResponse FavoriteList(1: FavoriteListRequest req)
}
