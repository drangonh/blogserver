/**
 * @author dragonH
 * @description 我的的oss账号信息
 * @date 2020/5/23 0023 下午 12:29
 */

package models

type OSS struct {
	Id              int    `orm:"pk;auto;column(id)" json:"id"`
	Uid             int    `orm:"column(uid);unique" json:"uid"`
	Endpoint        string `orm:"column(endPoint)" json:"endpoint"`
	AccessKeyId     string `orm:"column(accessKeyId)" json:"accessKeyId"`
	AccessKeySecret string `orm:"column(accessKeySecret)"  json:"accessKeySecret"`
	BucketName      string `orm:"column(bucketName)" json:"bucketName"`
}

func (o *OSS) TableName() string {
	return TNOSSStore()
}

func NewOSS() *OSS {
	return &OSS{}
}

/**
 * @Author dragonH
 * @Description 获取用户的oss信息
 * @Param nil
 * @return (oss *OSS, err error)
 **/
func (o *OSS) GetOSSInfo() (oss *OSS, err error) {
	orm := GetOrm("w")

	ossList := make([]*OSS, 2)
	num, err := orm.QueryTable(o.TableName()).All(ossList)
	if num > 0 && err == nil {
		return ossList[0], err
	}
	return
}

/**
 * @Author dragonH
 * @Description 更新或者插入oss信息//TODO
 * @Param
 * @return
 **/
func (o *OSS) Edit() (err error) {
	ormer := GetOrm("w")
	if o.Id > 0 {
		_, err = ormer.Update(o)
	} else {
		_, err = ormer.Insert(o)
	}
	return
}
