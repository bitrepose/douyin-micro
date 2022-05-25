namespace go comment

include "user.thrift"

struct Comment {
    1: required i64 id;               // 视频评论id
    2: required user.User user;       // 评论用户信息
    3: required string content;       // 评论内容
    4: required string create_date;   // 评论发布日期，格式mm-dd
}

struct CommentActionRequest {
    1: required i64     user_id;       // 用户id
    2: required i64     video_id;      // 视频id
    3: required i32     action_type;   // 1-发布评论，2-删除评论
    4: optional string  comment_text;  // 用户填写的评论内容，action_type为1时使用
    5: optional string  comment_id;    // 要删除的评论id，action_type为2时使用
}

struct CommentActionResponse {
    1: required i32     status_code;   // 状态码，0-成功，其他值-失败
    2: optional string  status_msg;    // 返回状态描述
    3: optional Comment comment;       // 评论成功返回评论内容，不需要重新拉取整个列表
}

struct CommentListRequest {
    1: required i64     user_id;       // 用户id
    2: required i64     req_user_id;   // 用户鉴权token
    3: required i64     video_id;      // 视频id
}

struct CommentListResponse {
    1: required i32          status_code;      // 状态码，0-成功，其他值-失败
    2: optional string       status_msg;       // 返回状态描述
    3: required list<Comment>  comment_list;   // 评论列表
}

service CommentService {
    CommentActionResponse CommentAction(1: CommentActionRequest req)
    CommentListResponse CommentList(1: CommentListRequest req)
}
