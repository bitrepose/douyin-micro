package rpc

import (
	"douyin-micro/kitex_gen/user/userservice"
	"douyin-micro/pkg/constants"
	"errors"

	etcd "github.com/kitex-contrib/registry-etcd"
)

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.UserEtcdAddress})
	if err != nil{
		panic(errors.New("can't connect to the User Service"))
	}
	c,err:=userservice.NewClient(
		cons
	)
}