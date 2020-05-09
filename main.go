package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"github.com/ChimeraCoder/anaconda"
)

func main () {

	twitterAccount := TwitterAccount {
		AccessToken: "1204332742592233473-VLLAJpwBwYFRnUxNyqKruC9hIhUx6C",
		AccessTokenSecret: "NYug91BBieMGvw3e1kKTlncroQ4HieR6rlWUH8NJuXSWs",
		ConsumerKey: "vbtvrinHC9TJQkJWLc9zTU2xC",
		ConsumerSecret: "oRYaeKci98jP886HMQ1da6uFvCZTvpo2YbANRCBMbopSJXE4kR"
	}
}

type TwitterAccount struct {
    AccessToken       string `json:"accessToken"`
    AccessTokenSecret string `json:"accessTokenSecret"`
    ConsumerKey       string `json:"consumerKey"`
    ConsumerSecret    string `json:"consumerSecret"`
}