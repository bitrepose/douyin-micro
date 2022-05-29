package jwtutil

import (
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
)

func TestJwt(t *testing.T) {
	tokenStr, err := CreateToken(1)
	if err != nil {
		klog.Error(err)
	}
	klog.Info(tokenStr)
	// time.Sleep(time.Second * 20)
	uid, iss, err := ParseToken(tokenStr)
	if err != nil {
		klog.Error(err)
	}
	klog.Info(uid, iss)
}
