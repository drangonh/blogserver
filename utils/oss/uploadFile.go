package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"strings"
)

func UploadFile(fileName string, path string, fileType int) (err error) {
	// 获取存储空间。
	bucket, err := Client.Bucket(BucketName)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	// 指定存储类型为标准存储，缺省也为标准存储。
	storageType := oss.ObjectStorageClass(oss.StorageStandard)

	// 指定访问权限为私有，缺省为继承bucket的权限。
	objectAcl := oss.ObjectACL(oss.ACLPrivate)

	// 上传字符串。
	if fileType == 1 {
		err = bucket.PutObject(PhotoPath+fileName, strings.NewReader(path), storageType, objectAcl)
	} else if fileType == 2 {
		err = bucket.PutObject(LanguagePath+fileName, strings.NewReader(path), storageType, objectAcl)
	} else if fileType == 3 {
		err = bucket.PutObject(ArticlePath+fileName, strings.NewReader(path), storageType, objectAcl)
	}
	return err
}
