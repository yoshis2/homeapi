package server

import (
	"homeapi/applications/logging"
	"net/http/httptest"
	"testing"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/gavv/httpexpect"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

type RoutingModuleTestSuite struct {
	suite.Suite

	server        *httptest.Server
	db            *gorm.DB
	redisClient   *redis.Client
	twitterClient *twitter.Client
	logging       logging.Logging
	validate      *validator.Validate
	mockCtrl      *gomock.Controller
}

func (route *RoutingModuleTestSuite) SetupSuite(t *testing.T) {
	route.mockCtrl = gomock.NewController(route.T())

	route.redisClient, _ = redismock.NewClientMock()
	// route.db, _, _ = databases.GormMock()

	suite.Run(t, new(RoutingModuleTestSuite))
}

func (route *RoutingModuleTestSuite) TestRun() {
	e := httpexpect.New(route.T(), route.server.URL)

	e.GET("/api/v1/temperatures1").
		Expect().
		Status(401).
		JSON()
	route.Run("temperature checking", func() {
		e.GET("/api/v1/temperatures1").
			Expect().
			Status(401).
			JSON()
	})
}
