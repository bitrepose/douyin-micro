package main

import (
	"context"
	"douyin-micro/cmd/user/service"
	"douyin-micro/kitex_gen/user"
	"douyin-micro/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationAction implements the UserServiceImpl interface.
func (s *UserServiceImpl) RelationAction(ctx context.Context, req *user.RelationActionRequest) (resp *user.RelationActionResponse, err error) {
	// TODO: Your code here...
	resp = new(user.RelationActionResponse)
	if req.ActionType < 1 || req.ActionType > 2 {
		e := errno.ParamErr
		convRelationActionResp(resp, int32(e.ErrCode), e.ErrMsg)
		return resp, nil
	}
	err = service.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		e := errno.ServiceErr
		convRelationActionResp(resp, int32(e.ErrCode), e.ErrMsg)
		return resp, nil
	}
	e := errno.Success
	convRelationActionResp(resp, int32(e.ErrCode), e.ErrMsg)
	return resp, nil
}

//上个方法的自用函数
func convRelationActionResp(resp *user.RelationActionResponse, code int32, msg string) {
	resp.StatusCode = code
	resp.StatusMsg = &msg
}

// RelationFollowList implements the UserServiceImpl interface.
func (s *UserServiceImpl) RelationFollowList(ctx context.Context, req *user.RelationFollowListRequest) (resp *user.RelationFollowListResponse, err error) {
	// TODO: Your code here...
	resp = new(user.RelationFollowListResponse)
	users, err := service.NewRelationFollowListService(ctx).RelationFollowList(req)
	if err != nil {
		e := errno.ServiceErr
		convRelationFollowListResponse(resp, int32(e.ErrCode), e.ErrMsg)
		resp.UserList = []*user.User{}
		return resp, nil
	}
	e := errno.Success
	convRelationFollowListResponse(resp, int32(e.ErrCode), e.ErrMsg)
	resp.UserList = users
	return resp, nil
}

//上个方法的自用函数
func convRelationFollowListResponse(resp *user.RelationFollowListResponse, code int32, msg string) {
	resp.StatusCode = code
	resp.StatusMsg = &msg
}

// RelationFollowerList implements the UserServiceImpl interface.
func (s *UserServiceImpl) RelationFollowerList(ctx context.Context, req *user.RelationFollowerListRequest) (resp *user.RelationFollowerListResponse, err error) {
	// TODO: Your code here...
	resp = new(user.RelationFollowerListResponse)
	users, err := service.NewRelationFollowerListService(ctx).RelationFollowerList(req)
	if err != nil {
		e := errno.ServiceErr
		convRelationFollowerListResponse(resp, int32(e.ErrCode), e.ErrMsg)
		return resp, nil
	}
	e := errno.Success
	convRelationFollowerListResponse(resp, int32(e.ErrCode), e.ErrMsg)
	resp.UserList = users
	return resp, nil
}

//上一个函数自用
func convRelationFollowerListResponse(resp *user.RelationFollowerListResponse, code int32, msg string) {
	resp.StatusCode = code
	resp.StatusMsg = &msg
}
