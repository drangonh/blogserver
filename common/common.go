package common

import (
	"strings"

	"github.com/astaxie/beego"
)

// session
const SessionName = "blog_server"

// cookie过期时间设置
const CookieMastLiftTime = 3600

// app_key
func AppKey() string {
	return beego.AppConfig.DefaultString("app_key", "blog_secret")
}

//默认头像
func DefaultAvatar() string {
	return beego.AppConfig.DefaultString("avatar", "/static/images/headimgurl.jpg")
}

//获取文件类型
func getFileExt() []string {
	ext := beego.AppConfig.DefaultString("upload_file_ext", "png|jpg|jpeg|gif|txt|doc|docx|pdf")
	temp := strings.Split(ext, "|")
	exts := make([]string, len(temp))

	i := 0
	for _, item := range temp {
		if item != "" {
			exts[i] = item
			i++
		}
	}
	return exts
}

//是否允许该类文件类型
func IsAllowedFileExt(ext string) bool {

	if strings.HasPrefix(ext, ".") {
		ext = string(ext[1:])
	}
	exts := getFileExt()

	for _, item := range exts {
		if strings.EqualFold(item, ext) {
			return true
		}
	}
	return false
}
