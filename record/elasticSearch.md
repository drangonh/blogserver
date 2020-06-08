**安装**
* docker pull docker.elastic.co/elasticsearch/elasticsearch:7.7.1

**启动**
* docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.7.1

**测试elasticSearch**
* goland-->Tools-->Http Client-->Test Restful Web Service
* postman

**全文搜索**
* http://localhost:9200/blogserver/user/_search?q=hl， 其中_search表示全文搜索。q=hl表示搜索关键字"hl"
。其中index/type/是可以省略的，也就是（blogserver/user/）
* http://localhost:9200/blogserver/user/_mapping  可以查看表结构
* elastic search在6.x版本调整了， 一个index只能存储一种type。

**请求方法**
* Put/Post 都可以创建/修改 数据，但是Post新增数据可以省略id，elasticsearch会自动新增id
* get 获取数据，全文搜索 index/type/_search?q= 