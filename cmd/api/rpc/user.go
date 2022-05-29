package rpc

import (
	"context"
	jwtutil "douyin-micro/cmd/api/jwt_util"
	"douyin-micro/kitex_gen/user"
	"douyin-micro/kitex_gen/user/userservice"
	"douyin-micro/pkg/constants"
	"douyin-micro/pkg/errno"
	"douyin-micro/pkg/middleware"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func Login(c context.Context, username string, password string) (int64, errno.ErrNo) {
	var resp *user.UserLoginResponse
	var err error
	resp, err = userClient.UserLogin(c, &user.UserLoginRequest{Username: username, Password: password})
	if err != nil {
		return 0, errno.ConvertErr(err)
	}
	if resp.StatusCode != 0 {
		return 0, errno.NewErrNo(int64(resp.StatusCode), *resp.StatusMsg)
	}
	return resp.UserId, errno.Success
}

func Register(c context.Context, username string, password string) (int64, string, errno.ErrNo) {
	var resp *user.UserRegisterResponse
	var err error
	var token string
	resp, err = userClient.UserRegister(c, &user.UserRegisterRequest{Username: username, Password: password})
	if err != nil {
		return 0, "", errno.ConvertErr(err)
	}
	if resp.StatusCode != 0 {
		return 0, "", errno.NewErrNo(int64(resp.StatusCode), *resp.StatusMsg)
	}
	//token 生成
	token, err = jwtutil.CreateToken(resp.UserId)
	if err != nil {
		return 0, "", errno.ServiceErr
	}
	return resp.UserId, token, errno.Success
}

func UserInfo(c context.Context, userId int64, toUserId int64) (user.User, errno.ErrNo) {
	klog.Info("rpc...")
	var resp *user.MUserInfoResponse
	var err error
	request := &user.MUserInfoRequest{ReqUserId: &userId, UserIds: []int64{toUserId}}
	resp, err = userClient.MUserInfo(c, request)
	if err != nil {
		return user.User{}, errno.ConvertErr(err)
	}
	if resp.StatusCode != 0 || len(resp.Users) == 0 {
		return user.User{}, errno.NewErrNo(errno.ServiceErr.ErrCode, *resp.StatusMsg)
	}
	return *resp.Users[0], errno.Success
}

func RelationAction(c context.Context, request user.RelationActionRequest) errno.ErrNo {
	var err error
	var resp *user.RelationActionResponse
	resp, err = userClient.RelationAction(c, &request)
	if err != nil {
		return errno.ConvertErr(err)
	}
	if resp.StatusCode != 0 {
		return errno.NewErrNo(int64(resp.StatusCode), *resp.StatusMsg)
	}
	return errno.Success
}

func RelationFollowList(c context.Context, request user.RelationFollowListRequest) ([]*user.User, errno.ErrNo) {
	var err error
	var resp *user.RelationFollowListResponse
	resp, err = userClient.RelationFollowList(c, &request)
	if err != nil {
		return []*user.User{}, errno.ConvertErr(err)
	}
	if resp.StatusCode != 0 {
		return []*user.User{}, errno.NewErrNo(int64(resp.StatusCode), *resp.StatusMsg)
	}
	return resp.UserList, errno.Success
}

func RelationFollowerList(c context.Context, request user.RelationFollowerListRequest) ([]*user.User, errno.ErrNo) {
	var err error
	var resp *user.RelationFollowerListResponse
	resp, err = userClient.RelationFollowerList(c, &request)
	if err != nil {
		return []*user.User{}, errno.ConvertErr(err)
	}
	if resp.StatusCode != 0 {
		return []*user.User{}, errno.NewErrNo(int64(resp.StatusCode), *resp.StatusMsg)
	}
	return resp.UserList, errno.Success
}
