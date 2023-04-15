package routes

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RoutesSuite struct {
	suite.Suite
}

func (s *RoutesSuite) SetupTest() {
	RegisterApiRoutes()
	RegisterWebRoutes()
}

func (s *RoutesSuite) TestRegisterRoutes() {
	routes := reflect.ValueOf(http.DefaultServeMux).Elem().FieldByIndex([]int{1}).MapRange()

	expected := []string{"/api/healthcheck", "/api/double", "/"}
	for routes.Next() {
		assert.Contains(s.T(), expected, routes.Key().String())
	}
}

func TestRunRoutesSuite(t *testing.T) {
	suite.Run(t, new(RoutesSuite))
}
