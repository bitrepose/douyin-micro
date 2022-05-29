package main

import (
	"douyin-micro/cmd/api/router"
	"douyin-micro/cmd/api/rpc"
	"douyin-micro/pkg/constants"
	"douyin-micro/pkg/tracer"
)

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
}

func main() {
	// Init()

	err := router.SetUpRouter()
	if err != nil {
		panic(err.Error())
	}

}
