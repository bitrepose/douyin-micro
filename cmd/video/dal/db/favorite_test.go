package db

import (
	"context"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
)

func TestFavorite(t *testing.T) {
	Init()
	err := Favorite(context.Background(), 3, 2)
	klog.Error(err)
}

func TestDisFavorite(t *testing.T) {

}

func TestFavoriteList(t *testing.T) {
	Init()
	res, err := FavoriteList(context.Background(), 2)
	if err != nil {
		klog.Error(err)
	}
	for _, v := range res {
		klog.Info(v)
	}

}

func TestFavoriteIdList(t *testing.T) {
	Init()
	res, err := FavoriteIdList(context.Background(), 2)
	if err != nil {
		klog.Error(err)
	}
	for _, v := range res {
		klog.Info(v)
	}

}
