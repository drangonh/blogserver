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

**全文搜索**
* 在body里面可以传递搜索条件，按照application/json传递。
* must,全文搜索必须满足address、city的要求
```
go
{
	"query":{
		"bool":{
			"must":[
					{"match":{"address":"880Holmes"}},
					{"match":{"city":"Brogan1"}}	
			]
		}
	}
}
```
* should,全文搜索，只要满足address、city其中之一就好
```
{
	"query":{
		"bool":{
			"should":[
					{"match":{"address":"880Holmes"}},
					{"match":{"city":"Brogan1"}}	
			]
		}
	}
}
```
* 按一种条件进行全文搜索，同样是在body中进行传参搜索的
```
{
	"query":{
		"match":{
			"address":"12"
		}
	}
}
```

**ES在项目中实际运用**
* 数据在存入数据库的时候写入ES中
* ES6.0之后一个索引中只能又一个Type
* ES中存放的数据只需要存放id和关键词相关的字段，例如bookId、bookName、bookContent
* ES完整工作流程：  
存放数据到数据库-->存放到ES  
搜索-->ES，返回包含id的slice，然后去数据库中查询出完整的数据返回给前端
* ES是一个http服务，ES的操作可以通过github.com/astaxie/beego/httplib这个库，也可以通过ES
自带的API。