package service

import (
	"github.com/YangZhaoWeblog/swift_sec_kill/internal/biz"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserXService)

type UserXService struct {
	uc *biz.UserXUseCase
}

// NewUserXService
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@param uc
//	@return *UserXService
func NewUserXService(uc *biz.UserXUseCase) *UserXService {
	return &UserXService{uc: uc}
}
