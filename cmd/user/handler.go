package main

import (
	"context"
	"douyin-micro/cmd/user/service"
	"douyin-micro/kitex_gen/user"
	"douyin-micro/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	resp = new(user.UserRegisterResponse)

	if len(req.UserName) == 0 || len(req.Password) == 0 || len(req.UserName) > 32 || len(req.Password) > 32{
		e := errno.ParamErr
		convUserRegisterRequest(resp, int32(e.ErrCode), e.ErrMsg)
		return resp, nil
	}

	err = service.NewUserRegisterService(ctx).UserRegister(req)
	if err != nil {
		e := errno.ServiceErr
		convUserRegisterRequest(resp, int32(e.ErrCode), e.ErrMsg)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

//上个方法的自用函数
func convUserRegisterRequest(resp *user.UserRegisterRequest, code int32, msg string) {
	resp.StatusCode = code
	resp.StatusMsg = &msg
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	resp = new(user.UserLoginResponse)
	
	if len(req.UserName) == 0 || len(req.Password) == 0 || len(req.UserName) > 32 || len(req.Password) > 32{
		e := errno.ParamErr
		convUserLoginResponse(resp, int32(e.ErrCode), e.ErrMsg)
		return resp, nil
	}

	uid, err := service.NewUserLoginService(ctx).UserLogin(req)
	if err != nil {
		e := errno.ServiceErr
		convUserLoginResponse(resp, int32(e.ErrCode), e.ErrMsg)
		return resp, nil
	}
	resp.UserId = uid
	e := errno.Success
	convUserLoginResponse(resp, int32(e.ErrCode), e.ErrMsg)
	return resp, nil
}

//上个方法的自用函数
func convUserLoginResponse(resp *user.UserLoginResponse, code int32, msg string) {
	resp.StatusCode = code
	resp.StatusMsg = &msg
}


// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) MUserInfo(ctx context.Context, req *user.MUserInfoRequest) (resp *user.MUserInfoResponse, err error) {
	resp = new(user.MUserInfoResponse)

	if len(req.UserIds) == 0 {
		e := errno.ParamErr
		convMUserInfoRequest(resp, int32(e.ErrCode), e.ErrMsg)
		return resp, nil
	}

	users, err := service.NewMUserInfoService(ctx).MUserInfo(req)
	if err != nil {
		e := errno.ServiceErr
		convMUserInfoRequest(resp, int32(e.ErrCode), e.ErrMsg)
		return resp, nil
	}
	e := errno.Success
	resp.Users = users
	convMUserInfoRequest(resp, int32(e.ErrCode), e.ErrMsg)
	return resp, nil
}

//上个方法的自用函数
func convMUserInfoRequest(resp *user.MUserInfoRequest, code int32, msg string) {
	resp.StatusCode = code
	resp.StatusMsg = &msg
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
