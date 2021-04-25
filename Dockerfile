## 载入最新的 golang
FROM golang:1.15.2

#MAINTAINER  dragon

## 定义环境变量
#ENV GOROOT=/usr/local/go
#ENV GOPATH=$GOPATH
#ENV PATH=$GOPATH/bin/:$PATH

#//正常这2句不需要，上网   逼得
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

## 首先安装 beego 和 bee 工具和 mysql 扩展
RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/mailru/easyjson
RUN go get github.com/russross/blackfriday
RUN go get github.com/nfnt/resize
RUN go get github.com/PuerkitoBio/goquery
RUN go get github.com/aliyun/aliyun-oss-go-sdk/oss
RUN go get github.com/bitly/go-simplejson
RUN go get github.com/olivere/elastic
RUN go get golang.org/x/time

## 复制Dockerfile 同目录的所有文件到docker 中的指定位置
COPY . /go/src/blogserver/

## 设置项目目录
WORKDIR /go/src/blogserver

## 开放 3000 端口
EXPOSE 3000

## 启动执行 bee run  命令
CMD ["bee", "run"]

#docker run -it --name blog-instance -p 3000:3000  -v  /Users/mac/Desktop/git/drag/go/src/blogserver:/go/src/blogserver -w /go/src/blogserver blogserver
