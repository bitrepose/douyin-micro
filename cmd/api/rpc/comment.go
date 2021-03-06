package rpc

import (
	"context"
	"douyin-micro/kitex_gen/comment"
	"douyin-micro/kitex_gen/comment/commentservice"
	"douyin-micro/pkg/constants"
	"douyin-micro/pkg/errno"
	"douyin-micro/pkg/middleware"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var commentClient commentservice.Client

func initCommentRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := commentservice.NewClient(
		constants.CommentServiceName,
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
	commentClient = c
}

func CommentList(ctx context.Context, req *comment.CommentListRequest) (*comment.CommentListResponse, error) {
	resp, err := commentClient.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(int64(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

func CommentAction(ctx context.Context, req *comment.CommentActionRequest) error {
	resp, err := commentClient.CommentAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != errno.SuccessCode {
		return errno.NewErrNo(int64(resp.StatusCode), *resp.StatusMsg)
	}
	return nil
}
