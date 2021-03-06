namespace go user

struct User {
    1: required i64     id;               // 用户id
    2: required string  name;             // 用户名称
    3: optional i64     follow_count;     // 关注总数
    4: optional i64     follower_count;   // 粉丝总数
    5: required bool    is_follow;        // true-已关注，false-未关注
}

struct UserRegisterRequest {
    1: required string  username;   // 注册用户名，最长32个字符
    2: required string  password;   // 密码，最长32个字符
}

struct UserRegisterResponse {
    1: required i32     status_code;   // 状态码，0-成功，其他值-失败
    2: optional string  status_msg;    // 返回状态描述
    3: required i64     user_id;       // 用户id
}

struct UserLoginRequest {
    1: required string  username;   // 登陆用户名
    2: required string  password;   // 登陆密码
}

struct UserLoginResponse {
    1: required i32     status_code;   // 状态码，0-成功，其他值-失败
    2: optional string  status_msg;    // 返回状态描述
    3: required i64     user_id;       // 用户id
}

struct MUserInfoRequest {
    1: required list<i64>     user_ids;      // 用户id
    2: optional i64           req_user_id;   // 发起请求的用户id，考虑到游客，设为可选
}

struct MUserInfoResponse {
    1: required i32           status_code;   // 状态码，0-成功，其他值-失败
    2: optional string        status_msg;    // 返回状态描述
    3: required list<User>    users;         // 用户信息
}

struct RelationActionRequest {
    1: required i64     user_id;       // 用户id
    3: required i64     to_user_id;    // 对方用户id
    4: required i32     action_type;   // 1-关注，2-取消关注 
}

struct RelationActionResponse {
    1: required i32     status_code;   // 状态码，0-成功，其他值-失败
    2: optional string  status_msg;    // 返回状态描述
}

struct RelationFollowListRequest {
    1: required i64     user_id;       // 用户id
    2: required i64     req_user_id;   // 发起请求的用户id
}

struct RelationFollowListResponse {
    1: required i32         status_code;   // 状态码，0-成功，其他值-失败
    2: optional string      status_msg;    // 返回状态描述
    3: required list<User>  user_list;     // 用户信息列表
}

struct RelationFollowerListRequest {
    1: required i64     user_id;       // 用户id
    2: required i64     req_user_id;   // 发起请求的用户id
}

struct RelationFollowerListResponse {
    1: required i32         status_code;   // 状态码，0-成功，其他值-失败
    2: optional string      status_msg;    // 返回状态描述
    3: required list<User>  user_list;     // 用户信息列表
}

service UserService {
    UserRegisterResponse UserRegister(1: UserRegisterRequest req)
    UserLoginResponse UserLogin(1: UserLoginRequest req)
    MUserInfoResponse MUserInfo(1: MUserInfoRequest req)
    RelationActionResponse RelationAction(1: RelationActionRequest req)
    RelationFollowListResponse RelationFollowList(1: RelationFollowListRequest req)
    RelationFollowerListResponse RelationFollowerList(1: RelationFollowerListRequest req)
}
