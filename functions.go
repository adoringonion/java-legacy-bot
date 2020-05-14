package functions

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/dghubble/go-twitter/twitter"
)

func GetTweets(ctx context.Context, m PubSubMessage) error {

	api := ConnectTwitterAPI()
	serachRes, _ := api.GetSearch(`テスト`, nil)
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

	if f, err := os.Stat("./env.json"); os.IsNotExist(err) || f.IsDir() {
		row, err := ioutil.ReadFile("./env.json")
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		
	}
	AccessToken := os.Getenv("ACCESS_TOKEN")
	AccessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	ConsumerKey := os.Getenv("CONSUMER_KEY")
	ConsumerSecret := os.Getenv("CONSUMER_SECRET")

	return anaconda.NewTwitterApiWithCredentials(AccessToken, AccessTokenSecret, ConsumerKey, ConsumerSecret)

}

type PubSubMessage struct {
	Data []byte `json:"data"`
}
