package persist

import (
	"../engine"
	"context"
	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index string) (chan engine.Item,error) {
	// must turn off sniff for docker
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil,err
	}

	out := make(chan engine.Item)

	go func() {

		itemCount := 0
		for {
			item := <-out
			log.Printf("item saver, got item #%d: %v", itemCount, item)
			itemCount++

			err := save(client, item,index)
			if err != nil {
				log.Printf("item saver err: saving %v: %v", item, err)
			}
		}
	}()

	return out,nil
}

func save(client *elastic.Client, item engine.Item,index string) error {

	if item.Type == "" {
		return errors.New("must supply type")
	}

	indexService := client.Index().Index(index).Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())
	if err != nil {
		return nil
	}

	return nil
}
