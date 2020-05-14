package functions

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
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

func ConnectTwitterAPI() *twitter.Client {

	var twitterAuth TwitterAuth

	if f, err := os.Stat("./env.json"); os.IsNotExist(err) || f.IsDir() {
		row, err := ioutil.ReadFile("./env.json")
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		json.Unmarshal(row, &twitterAuth)

	}else {
		twitterAuth = TwitterAuth{os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"), os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET")}
	}

	config := oauth1.NewConfig(twitterAuth.ConsumerKey, twitterAuth.ConsumerSecret)
	token := oauth1.NewToken(twitterAuth.AccessToken, twitterAuth.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	return client
}

type PubSubMessage struct {
	Data []byte `json:"data"`
}

// TwitterAuth twitter認証用の構造体
type TwitterAuth struct {
	AccessToken string
	AccessTokenSecret string
	ConsumerKey string
	ConsumerSecret string
}