package services

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServicesSuite struct {
	suite.Suite
}

func (s *ServicesSuite) TestHealthcheck() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	Healthcheck(rec, req)

	assert.Equal(s.T(), http.StatusOK, rec.Result().StatusCode)
	assert.Equal(s.T(), "OK", rec.Body.String())
}

func (s *ServicesSuite) TestDouble() {
	q := make(url.Values)
	q.Set("number", "2")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
	rec := httptest.NewRecorder()

	Double(rec, req)

	assert.Equal(s.T(), http.StatusOK, rec.Code)
	assert.Equal(s.T(), "{\"result\":4}\n", rec.Body.String())
}

func (s *ServicesSuite) TestDoubleFailure() {
	q := make(url.Values)
	q.Set("number", "NAN")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
	rec := httptest.NewRecorder()

	Double(rec, req)

	assert.Equal(s.T(), http.StatusBadRequest, rec.Code)
	assert.Equal(s.T(), "\"param number was invalid\"\n", rec.Body.String())
}

func TestRunServicesSuite(t *testing.T) {
	suite.Run(t, new(ServicesSuite))
}
