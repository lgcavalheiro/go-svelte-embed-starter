package services

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServicesSuite struct {
	suite.Suite
	Echo *echo.Echo
}

func (s *ServicesSuite) SetupTest() {
	s.Echo = echo.New()
}

func (s *ServicesSuite) TestHealthcheck() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := s.Echo.NewContext(req, rec)
	c.SetPath("/")

	if assert.NoError(s.T(), Healthcheck(c)) {
		assert.Equal(s.T(), http.StatusOK, rec.Code)
		assert.Equal(s.T(), "OK", rec.Body.String())
	}
}

func (s *ServicesSuite) TestDouble() {
	q := make(url.Values)
	q.Set("number", "2")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := s.Echo.NewContext(req, rec)
	c.SetPath("/")

	if assert.NoError(s.T(), Double(c)) {
		assert.Equal(s.T(), http.StatusOK, rec.Code)
		assert.Equal(s.T(), "{\"result\":4}\n", rec.Body.String())
	}
}

func (s *ServicesSuite) TestDoubleFailure() {
	q := make(url.Values)
	q.Set("number", "NAN")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := s.Echo.NewContext(req, rec)
	c.SetPath("/")

	r := Double(c).(*echo.HTTPError)
	if r.Error() != "" {
		assert.Equal(s.T(), http.StatusBadRequest, r.Code)
		assert.Equal(s.T(), "code=400, message=param number was invalid", r.Error())
	}
}

func TestRunServicesSuite(t *testing.T) {
	suite.Run(t, new(ServicesSuite))
}
