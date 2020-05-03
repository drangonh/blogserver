## 载入最新的 golang
FROM golang:1.14.2
## 定义环境变量
ENV GOROOT=/usr/local/go
ENV GOPATH=/Users/liudong/Desktop/git/go
ENV PATH=$GOPATH/bin/:$PATH
## 首先安装 beego 和 bee 工具和 mysql 扩展
RUN go get github.com/astaxie/beego && go get github.com/beego/bee && go get github.com/go-sql-driver/mysql
## 复制Dockerfile 同目录的所有文件到docker 中的指定位置
COPY . /www/go/src/goweb/
## 设置项目目录
WORKDIR /www/go/src/goweb
## 开放 8080 端口
EXPOSE 8080
## 启动执行 bee run  命令
CMD ["bee", "run"]
