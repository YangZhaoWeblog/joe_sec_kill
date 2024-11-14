//go:build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/YangZhaoWeblog/swift_sec_kill/internal/biz"
	"github.com/YangZhaoWeblog/swift_sec_kill/internal/conf"
	"github.com/YangZhaoWeblog/swift_sec_kill/internal/data"
	"github.com/YangZhaoWeblog/swift_sec_kill/internal/interfaces"
	"github.com/YangZhaoWeblog/swift_sec_kill/internal/server"
	"github.com/YangZhaoWeblog/swift_sec_kill/internal/service"
	"github.com/YangZhaoWeblog/swift_sec_kill/internal/task"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp
//
//	@Description: wireApp init kratos application.
//	@param *conf.Server
//	@param *conf.Data
//	@return *kratos.App
//	@return func()
//	@return error
func wireApp(*conf.Server, *conf.Data) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		interfaces.ProviderSet,
		task.ProviderSet,
		newApp))
}
