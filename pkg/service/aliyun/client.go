package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/fimreal/goutils/ezap"
	"github.com/ory/viper"
)

type AliyunAccessKey struct {
	ACCESS_KEY_ID     string
	ACCESS_KEY_SECRET string
	REGION_ID         string
}

var ak = &AliyunAccessKey{}

func NewClient() *ecs.Client {
	ak = &AliyunAccessKey{
		REGION_ID:         viper.GetString("aliyun.regionid"),
		ACCESS_KEY_ID:     viper.GetString("aliyun.akid"),
		ACCESS_KEY_SECRET: viper.GetString("aliyun.aksecret"),
	}
	ezap.Debugf("获取 阿里云 ecs 连接配置，ACCESS_KEY_ID: %s，ACCESS_KEY_SECRET: %s，REGION_ID: %s", ak.ACCESS_KEY_ID, ak.ACCESS_KEY_SECRET, ak.REGION_ID)
	client, err := ecs.NewClientWithAccessKey(ak.REGION_ID, ak.ACCESS_KEY_ID, ak.ACCESS_KEY_SECRET)
	if err != nil {
		ezap.Fatal(err)
	}
	if viper.GetBool("aliyun.insecureskipverify") {
		client.SetHTTPSInsecure(true) // 跳过证书验证，阿里云 sdk 似乎不支持使用 /etc/ssl/certs 中 ca 配置
	}
	return client
}
