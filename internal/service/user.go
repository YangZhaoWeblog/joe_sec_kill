package service

import (
	"context"
	v1 "github.com/YangZhaoWeblog/swift_sec_kill/api/my_sec_kill/v1"
	"github.com/YangZhaoWeblog/swift_sec_kill/internal/biz"
)

// CreateUser
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver s
//	@param ctx
//	@param req
//	@return *v1.CreateUserReply
//	@return error
func (s *UserXService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	_, err := s.uc.CreateUser(ctx, &biz.User{
		UserName: req.UserName,
		Pwd:      req.Pwd,
		Sex:      int(req.Sex),
		Age:      int(req.Age),
		Email:    req.Email,
		Contact:  req.Contact,
		Mobile:   req.Mobile,
		IdCard:   req.IdCard,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserReply{Message: "trytest"}, nil
}

func (s *UserXService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	return &v1.UpdateUserReply{}, nil
}

// DeleteUser
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver s
//	@param ctx
//	@param req
//	@return *v1.DeleteUserReply
//	@return error
func (s *UserXService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	return &v1.DeleteUserReply{}, nil
}

// GetUser
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver s
//	@param ctx
//	@param req
//	@return *v1.GetUserReply
//	@return error
func (s *UserXService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	userInfo, err := s.uc.GetUserInfo(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserReply{
		Code:    0,
		Message: "success",
		Data: &v1.GetUserReplyData{
			UserName: userInfo.UserName,
			Pwd:      userInfo.Pwd,
			Sex:      int32(userInfo.Sex),
			Age:      int32(userInfo.Age),
			Email:    userInfo.Email,
			Contact:  userInfo.Contact,
			Mobile:   userInfo.Mobile,
			IdCard:   userInfo.IdCard,
		},
	}, nil
}

// GetUserByName
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver s
//	@param ctx
//	@param req
//	@return *v1.GetUserByNameReply
//	@return error
func (s *UserXService) GetUserByName(ctx context.Context, req *v1.GetUserByNameRequest) (*v1.GetUserByNameReply, error) {
	userInfo, err := s.uc.GetUserInfoByName(ctx, req.UserName)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserByNameReply{
		Code:    0,
		Message: "success",
		Data: &v1.GetUserReplyData{
			Id:       int64(userInfo.UserID),
			UserName: userInfo.UserName,
			Pwd:      userInfo.Pwd,
			Sex:      int32(userInfo.Sex),
			Age:      int32(userInfo.Age),
			Email:    userInfo.Email,
			Contact:  userInfo.Contact,
			Mobile:   userInfo.Mobile,
			IdCard:   userInfo.IdCard,
		},
	}, nil
}

// ListUser
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver s
//	@param ctx
//	@param req
//	@return *v1.ListUserReply
//	@return error
func (s *UserXService) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error) {
	return &v1.ListUserReply{}, nil
}

// 执行定时任务
func (u *UserXService) Cronjob(ctx context.Context, req *v1.CreateUserRequest) {
	u.uc.CreateUser(ctx, &biz.User{
		UserName: req.UserName,
		Pwd:      req.Pwd,
		Sex:      int(req.Sex),
		Age:      int(req.Age),
		Email:    req.Email,
		Contact:  req.Contact,
		Mobile:   req.Mobile,
		IdCard:   req.IdCard,
	})
}
