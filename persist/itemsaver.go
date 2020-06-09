package persist

import (
	"blogserver/common"
	"context"
	"github.com/olivere/elastic"
)

func Save(item interface{}) error {
	client, err := elastic.NewClient(
		//must set false on docker
		elastic.SetSniff(false),
	)

	if err != nil {
		return err
	}

	_, err = client.Index().
		Index(common.ESIndex).
		Type(common.ESType).
		BodyJson(item).
		Do(context.Background())

	if err != nil {
		return err
	}

	//fmt.Printf("%v",resp)
	return nil
}
