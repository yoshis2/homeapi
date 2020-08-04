package ex_api

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Twitter struct {
}

func NewTwitter() *Twitter {
	return &Twitter{}
}

func (twitterOpen *Twitter) Open() *twitter.Client {
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECREET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	return twitter.NewClient(httpClient)
}
