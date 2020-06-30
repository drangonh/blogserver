**redis docker 安装**
* docker pull redis,拉取最新的redis版本
* redis 运行  docker run --name some-redis1 -p 6380:6379  -d redis --requirepass "123456"
* docker exec -it 0b63b7eacc29 redis-cli，redis连接客户端
* redis 登录：auth 123456
* redis set :set name hl
* redis get:get name