package rpc

import (
	"context"
	"douyin-micro/kitex_gen/user"
	"douyin-micro/kitex_gen/user/userservice"
	"douyin-micro/pkg/constants"
	"douyin-micro/pkg/errno"
	"time"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.UserEtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		// client.WithMiddleware(middleware.CommonMiddleware),
		// client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                    // mux
		client.WithRPCTimeout(3*time.Second),           // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond), // conn timeout
		// client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()), // tracer
		client.WithResolver(r),                          // resolver
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func MUserInfo(ctx context.Context, req *user.MUserInfoRequest) (map[int64]*user.User, error) {
	resp, err := userClient.MUserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(int64(resp.StatusCode), *resp.StatusMsg)
	}
	res := make(map[int64]*user.User)
	for _, u := range resp.Users {
		res[u.Id] = u
	}
	return res, nil
}
