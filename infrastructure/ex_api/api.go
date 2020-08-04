package ex_api

import "github.com/dghubble/go-twitter/twitter"

type TwitterInterface interface {
	Open() *twitter.Client
}
