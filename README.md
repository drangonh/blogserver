# blogserver
个人博客服务端
### 包含功能
* 登录注册
* 文章编辑，发布，列表
* 简历介绍 

### 第三方库安装
* go get github.com/astaxie/beego
* go get github.com/beego/bee

### 搭建项目
* 安装beego
* 安装bee
* bee new blogserver
* bee在go mod模式中还没法使用

### 如何把.sql文件倒入数据库
* 目的：通过.sql文件创建表
* mysql -u root -p,输入数据库密码就可以登录数据库
* create database 数据库名字，创建数据库
* source .sql文件的目录，例如source /Users/liudong/Desktop/git/go/src/ziyoubiancheng/mbook/mbook.sql
这个过程可以创建表

### 项目运行
* bee run main.go，使用go run无法还在beego中的默认配置(app.conf)

### API文档生成命令
* bee run -gendoc=true -downdoc=true，这会生成一个swagger目录，只有在测试环境中才会生成该文档

### API文档地址
* http://localhost:8888/swagger/#!

### 阿里云
* go get github.com/aliyun/aliyun-oss-go-sdk/oss

### 时区导致数据库中的时间比本地时间慢8个小时
* 解决办法在beego的注册数据库时加上&loc=Local，如下所示 
* ```dataSource := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&loc=Local"```