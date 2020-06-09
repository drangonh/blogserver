package persist

import (
	"blogserver/common"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"testing"
)

func TestSave(t *testing.T) {
	var obj map[string]string
	obj = make(map[string]string)
	obj["name"] = "hl"
	obj["des"] = "hua 花公子"

	err := Save(obj)

	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	q := elastic.NewMatchQuery("des", "hua")
	resp, err := client.Search().Index(common.ESIndex).Type(common.ESType).Query(q).Do(context.Background())
	if err != nil {
		panic(err)
	}

	printEmployee(resp, err)
}

//打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	for _, hit := range res.Hits.Hits {
		var item interface{}
		err := json.Unmarshal(hit.Source, &item) //另外一种取数据的方法
		if err != nil {
			fmt.Print(err.Error())
		}
		//items = append(items,item)

		fmt.Println(item)
	}
}
