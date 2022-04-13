// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/configuration"
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/handler"
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/server"
)

// Injectors from inject.go:

func initialize() server.GrpcServer {
	config := configuration.Get()
	echoProvider := handler.NewEchoProvider()
	bazelGrpcGatewayPracticeServiceV1 := handler.NewBazelGrpcGatewayPracticeServiceV1(echoProvider)
	grpcServer := server.NewGrpcServer(config, bazelGrpcGatewayPracticeServiceV1)
	return grpcServer
}
