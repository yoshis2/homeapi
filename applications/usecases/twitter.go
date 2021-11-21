package usecases

import (
	"context"
	"strconv"

	"homeapi/applications/logging"
	"homeapi/applications/ports"
	"homeapi/applications/repository"
	"homeapi/applications/util"
	"homeapi/domain"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type TwitterUsecase struct {
	TwitterRepository repository.TwitterRepository
	DB                *gorm.DB
	RedisClient       *redis.Client
	TwitterClient     *twitter.Client
	Logging           logging.Logging
	Validator         *validator.Validate
}

const FIRST_TURN = 0

func (usecase *TwitterUsecase) Get(ctx context.Context) error {
	key := "tweetTurn"
	tweetTurn, err := usecase.RedisClient.Get(ctx, key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			if err := usecase.RedisClient.Set(ctx, key, FIRST_TURN, 0).Err(); err != nil {
				usecase.Logging.Error(err)
				return err
			}
		}
	}

	// ターンから１プラス
	tweetTurnInt, _ := strconv.Atoi(tweetTurn)
	tweetTurnInt++

	// ターンと同じIDのツイートメッセージ取得
	tweetContents, err := usecase.TwitterRepository.Get(ctx, usecase.DB, tweetTurnInt)
	if err != nil {
		usecase.Logging.Error(err)
		return err
	}

	//tweet, res, err := client.Statuses.Update("ツイートする本文", nil)
	_, _, err = usecase.TwitterClient.Statuses.Update(tweetContents.Message, nil)
	if err != nil {
		usecase.Logging.Error(err)
		return err
	}

	//ツイートテーブルのmaxID取得
	maxID, err := usecase.TwitterRepository.Last(ctx, usecase.DB)
	if err != nil {
		usecase.Logging.Error(err)
		return err
	}

	if maxID == tweetTurnInt {
		tweetTurnInt = FIRST_TURN
	}

	tweetTurn = strconv.Itoa(tweetTurnInt)
	// 今のターンをredisにセットする
	if err := usecase.RedisClient.Set(ctx, key, tweetTurn, 0).Err(); err != nil {
		usecase.Logging.Error(err)
		return err
	}

	return nil
}

func (usecase *TwitterUsecase) Create(ctx context.Context, input *ports.TwitterInputPort) (*ports.TwitterInputPort, error) {
	now, err := util.JapaneseNowTime()
	if err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	twitter := &domain.Twitter{
		Message:   input.Message,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := usecase.TwitterRepository.Insert(ctx, usecase.DB, twitter); err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}
	return &ports.TwitterInputPort{Message: input.Message}, nil
}
