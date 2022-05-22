package minio

import (
	"douyin-micro/pkg/constants"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func InitMinio() {
	client, err := minio.New(constants.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(constants.MinioAccessKeyId, constants.MinioSecretAccessKey, ""),
		Secure: constants.MinioUseSSL,
	})
	if err != nil {
		panic(err)
	}
	// fmt.Println(client)
	klog.Info("minio client init successfully")
	minioClient = client
}
