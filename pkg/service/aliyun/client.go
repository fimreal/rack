package aliyun

import (
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/fimreal/goutils/ezap"
)

type AliyunAccessKey struct {
	ACCESS_KEY_ID     string
	ACCESS_KEY_SECRET string
	REGION_ID         string
}

var ak = &AliyunAccessKey{
	REGION_ID:         os.Getenv("REGION_ID"),
	ACCESS_KEY_ID:     os.Getenv("ACCESS_KEY_ID"),
	ACCESS_KEY_SECRET: os.Getenv("ACCESS_KEY_SECRET"),
}

func NewClient() *ecs.Client {
	// ezap.Debugf("获取 阿里云 ecs 连接配置，ACCESS_KEY_ID: %s，ACCESS_KEY_SECRET: %s，REGION_ID: %s", ak.ACCESS_KEY_ID, ak.ACCESS_KEY_SECRET, ak.REGION_ID)
	ezap.Debugf("获取 阿里云 ecs 连接配置，%v", ak)
	client, err := ecs.NewClientWithAccessKey(ak.REGION_ID, ak.ACCESS_KEY_ID, ak.ACCESS_KEY_SECRET)
	if err != nil {
		ezap.Fatal(err)
	}
	return client
}
