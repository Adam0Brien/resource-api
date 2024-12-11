// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tonytheleg/resource-api/internal/biz"
	"github.com/tonytheleg/resource-api/internal/conf"
	"github.com/tonytheleg/resource-api/internal/data"
	"github.com/tonytheleg/resource-api/internal/server"
	"github.com/tonytheleg/resource-api/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
// wireApp initializes the Kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataLayer, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}

	// Initialize repositories and use cases
	resourceRepo := data.NewResourceRepo(dataLayer, logger)
	relationshipRepo := data.NewRelationshipRepo(dataLayer, logger)

	resourceUsecase := biz.NewResourceUsecase(resourceRepo, logger)
	relationshipUsecase := biz.NewRelationshipUsecase(relationshipRepo, logger)

	// Initialize services
	resourceService := service.NewKesselResourceServiceService(resourceUsecase)
	relationshipService := service.NewKesselRelationshipServiceService(relationshipUsecase)

	// Initialize servers
	grpcServer := server.NewGRPCServer(confServer, resourceService, relationshipService, logger)
	httpServer := server.NewHTTPServer(confServer, resourceService, relationshipService, logger)

	// Create the application
	app := newApp(logger, grpcServer, httpServer)

	return app, func() {
		cleanup()
	}, nil
}

