package service

import (
	"context"
	"douyin-micro/cmd/user/dal/db"
	"douyin-micro/cmd/user/pack"
	"douyin-micro/kitex_gen/user"
)

type MUserInfoService struct {
	ctx context.Context
}

func NewMUserInfoService(ctx context.Context) *MUserInfoService {
	return &MUserInfoService{ctx: ctx}
}

// MGetUser multiple get list of user info
func (s *MUserInfoService) MUserInfo(req *user.MUserInfoRequest) ([]*user.User, error) {
	modelUsers, err := db.FindUsersByIds(s.ctx, &req.UserIds)
	if err != nil {
		return nil, err
	}
	var users = []*user.User{}
	for _, tempUser := range *modelUsers {
		if req.ReqUserId != nil {
			tempUser.IsFollow, err = db.Followed(s.ctx, *req.ReqUserId, tempUser.ID)
			//if err!=nil && !errors.As(err,&gorm.ErrRecordNotFound) {
			//	return nil,err
			//}
			// ?2. 如果有异常,说明没有关注 就让他为false
			if err != nil {
				tempUser.IsFollow = false
			}
		}

		users = append(users, pack.ConvUser(tempUser))
	}
	return users, nil
}
