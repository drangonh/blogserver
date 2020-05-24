package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	endpoint        string
	accessKeyId     string
	accessKeySecret string
	Client          *oss.Client
	BucketName      string
	PhotoPath       string
	LanguagePath    string
	ArticlePath     string
)

func init() {
	newClient()
}

/*
* 初始化oss
 */
func newClient() {
	// Endpoint以杭州为例，其它Region请按实际情况填写。https://oss-cn-shanghai.aliyuncs.com
	endpoint = ""

	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，
	//请登录 https://ram.console.aliyun.com 创建RAM账号。

	accessKeyId = ""

	accessKeySecret = ""

	BucketName = "blogdrag"

	// 设置文件地址
	// 用户头像上传地址
	PhotoPath = "photo/"

	//文章分类图片上传地址
	LanguagePath = "language/"

	// 文章图片上传地址
	ArticlePath = "article/"

	// 创建OSSClient实例。
	var err error
	Client, err = oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}

	// 初始化请求者付费模式。
	reqPayConf := oss.RequestPaymentConfiguration{
		Payer: "Requester",
	}

	// 设置bucket请求者付费模式。
	err = Client.SetBucketRequestPayment(BucketName, reqPayConf)
	if err != nil {
		handleError(err)
	}

	fmt.Println("初始化oss", Client)
}

func handleError(err error) error {
	fmt.Println("Error:", err)
	//os.Exit(-1)
	return err
}
