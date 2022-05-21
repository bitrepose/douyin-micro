package main

import (
	"context"
	"douyin-micro/kitex_gen/user"
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
	return
}

// RelationFollowList implements the UserServiceImpl interface.
func (s *UserServiceImpl) RelationFollowList(ctx context.Context, req *user.RelationFollowListRequest) (resp *user.RelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowerList implements the UserServiceImpl interface.
func (s *UserServiceImpl) RelationFollowerList(ctx context.Context, req *user.RelationFollowerListRequest) (resp *user.RelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}
