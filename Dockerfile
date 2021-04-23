## 载入最新的 golang
FROM golang:1.14.2

MAINTAINER  dragon

## 定义环境变量
ENV GOROOT=/usr/local/go
ENV GOPATH=$GOPATH
ENV PATH=$GOPATH/bin/:$PATH

## 首先安装 beego 和 bee 工具和 mysql 扩展
RUN go get github.com/astaxie/beego
RUN go get google.golang.org/protobuf
RUN go get github.com/beego/bee
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/russross/blackfriday
RUN go get github.com/nfnt/resize
RUN go get github.com/PuerkitoBio/goquery
RUN go get github.com/aliyun/aliyun-oss-go-sdk/oss
RUN go get github.com/bitly/go-simplejson
RUN go get github.com/olivere/elastic
RUN go get golang.org/x/net

## 复制Dockerfile 同目录的所有文件到docker 中的指定位置
COPY . /www/go/src/blogserver/

## 设置项目目录
WORKDIR /www/go/src/blogserver

## 开放 8080 端口
EXPOSE 3000

## 启动执行 bee run  命令
CMD ["bee", "run"]
