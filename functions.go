package functions

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"context"
)

func GetChinchin(ctx context.Context, m PubSubMessage) {

	api := connectTwitterAPI()
	serachRes, _ := api.GetSearch(`ちんちん`, nil)
	for _, tweet := range serachRes.Statuses {
		fmt.Println(tweet.Text)
	}

}

func connectTwitterAPI() *anaconda.TwitterApi {

	twitterAccount := TwitterAccount{
		AccessToken:       "1204332742592233473-VLLAJpwBwYFRnUxNyqKruC9hIhUx6C",
		AccessTokenSecret: "NYug91BBieMGvw3e1kKTlncroQ4HieR6rlWUH8NJuXSWs",
		ConsumerKey:       "vbtvrinHC9TJQkJWLc9zTU2xC",
		ConsumerSecret:    "oRYaeKci98jP886HMQ1da6uFvCZTvpo2YbANRCBMbopSJXE4kR",
	}

	return anaconda.NewTwitterApiWithCredentials(twitterAccount.AccessToken, twitterAccount.AccessTokenSecret, twitterAccount.ConsumerKey, twitterAccount.ConsumerSecret)

}

type TwitterAccount struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}

type PubSubMessage struct {
	Data []byte `json:"data"`
}