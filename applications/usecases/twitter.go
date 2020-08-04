package usecases

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"homeapi/applications"
	"homeapi/applications/logging"
	"homeapi/applications/ports"
	"homeapi/applications/repository"
	"homeapi/applications/util"
	"homeapi/domain"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type TwitterUsecase struct {
	TwitterRepository repository.TwitterRepository
	DB                *gorm.DB
	RedisClient       *redis.Client
	TwitterClient     *twitter.Client
	Logging           logging.Logging
}

const FIRST_TURN = 0

func (usecase *TwitterUsecase) Get() *applications.UsecaseError {
	time.Sleep(time.Second * time.Duration(rand.Intn(1200))) // redisからターンを取得
	key := "tweetTurn"
	tweetTurn, err := usecase.RedisClient.Get(key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			err := usecase.RedisClient.Set(key, FIRST_TURN, 0).Err()
			if uerr := applications.GetUErrorByError(err); uerr != nil {
				usecase.Logging.Error(uerr)
				return uerr
			}
		}
	}

	if uerr := applications.GetUErrorByError(err); uerr != nil {
		usecase.Logging.Error(uerr)
		return uerr
	}

	// ターンから１プラス
	log.Printf("キーの内容  : %v -- %v", key, tweetTurn)
	tweetTurnInt, _ := strconv.Atoi(tweetTurn)
	tweetTurnInt++

	// ターンと同じIDのツイートメッセージ取得
	tweetContents, err := usecase.TwitterRepository.Get(usecase.DB, tweetTurnInt)
	if uerr := applications.GetUErrorByError(err); uerr != nil {
		usecase.Logging.Error(uerr)
		return uerr
	}

	//tweet, res, err := client.Statuses.Update("ツイートする本文", nil)
	_, _, err = usecase.TwitterClient.Statuses.Update(tweetContents.Message, nil)
	if uerr := applications.GetUErrorByError(err); uerr != nil {
		usecase.Logging.Error(uerr)
		return uerr
	}

	//ツイートテーブルのmaxID取得
	maxID, err := usecase.TwitterRepository.Last(usecase.DB)
	if uerr := applications.GetUErrorByError(err); uerr != nil {
		usecase.Logging.Error(uerr)
		return uerr
	}

	log.Printf("maxID : %v", maxID)
	if maxID == tweetTurnInt {
		tweetTurnInt = FIRST_TURN
	}

	tweetTurn = strconv.Itoa(tweetTurnInt)
	// 今のターンをredisにセットする
	if err := usecase.RedisClient.Set(key, tweetTurn, 0).Err(); err != nil {
		fmt.Println("redis.Client.Set Error:", err)
	}
	return nil
}

func (usecase *TwitterUsecase) Create(input *ports.TwitterInputPort) (*ports.TwitterInputPort, *applications.UsecaseError) {
	twitter := &domain.Twitter{
		Message: input.Message,
	}

	var err error
	twitter.CreatedAt, err = util.JapaneseNowTime()
	if uerr := applications.GetUErrorByError(err); uerr != nil {
		usecase.Logging.Error(uerr)
		return nil, uerr
	}

	twitter.UpdatedAt, err = util.JapaneseNowTime()
	if uerr := applications.GetUErrorByError(err); uerr != nil {
		usecase.Logging.Error(uerr)
		return nil, uerr
	}

	err = usecase.TwitterRepository.Insert(usecase.DB, twitter)
	log.Printf("エラーの内容 : %v", err)
	if uerr := applications.GetUErrorByError(err); uerr != nil {
		usecase.Logging.Error(uerr)
		return nil, uerr
	}
	return nil, nil
}
