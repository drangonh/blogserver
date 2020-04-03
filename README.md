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