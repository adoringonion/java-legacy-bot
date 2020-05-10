package functions

import (
	"context"
	"github.com/ChimeraCoder/anaconda"
	"log"
	"os"
)

func GetChinchin(ctx context.Context, m PubSubMessage) error {

	api := ConnectTwitterAPI()
	serachRes, _ := api.GetSearch(`ちんちん`, nil)
	for _, tweet := range serachRes.Statuses {
		log.Println(tweet.Text)
	}

	return nil
}

func GetTweetsCount(searchWord string) (int, error) {
	api := ConnectTwitterAPI()
	searchRes, err := api.GetSearch(searchWord, nil)
	return len(searchRes.Statuses), err
}

func ConnectTwitterAPI() *anaconda.TwitterApi {

	AccessToken := os.Getenv("ACCESS_TOKEN")
	AccessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	ConsumerKey := os.Getenv("CONSUMER_KEY")
	ConsumerSecret := os.Getenv("CONSUMER_SECRET")

	return anaconda.NewTwitterApiWithCredentials(AccessToken, AccessTokenSecret, ConsumerKey, ConsumerSecret)

}

type PubSubMessage struct {
	Data []byte `json:"data"`
}
