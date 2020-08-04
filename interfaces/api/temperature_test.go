package api

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

var (
	testTemp = `{"temp":"22", "humi":"60", "created_at":2018-03-10 15:00:00:00}`
)

func TestSetTemp(t *testing.T) {
	api := TemperatureJSON{}

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/api/v1/temperatures", strings.NewReader(testTemp))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	tempReal := api.Insert(c)
	log.Printf("tempReal : %v", tempReal)
	log.Printf("rec : %v", rec.Code)

	assert.Equal(t, 200, rec.Code)
}
