// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"pigeon2/internal/biz"
	"pigeon2/internal/conf"
	"pigeon2/internal/data"
	"pigeon2/internal/server"
	"pigeon2/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	grpcServer := server.NewGRPCServer(confServer, greeterService, logger)
	workFlowRepo := data.NewWorkFlowRepo(dataData, logger)
	temproalRepo := data.NewTemporalRepo(confData, logger)
	workFlowUsecase := biz.NewWorkFlowUsecase(workFlowRepo, temproalRepo, logger)
	workFlowService := service.NewWorkService(workFlowUsecase)
	httpServer := server.NewHTTPServer(confServer, greeterService, workFlowService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
