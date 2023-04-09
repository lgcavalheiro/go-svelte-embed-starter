package routes

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RoutesSuite struct {
	suite.Suite
	Echo *echo.Echo
}

func (s *RoutesSuite) SetupTest() {
	s.Echo = echo.New()
	RegisterApiRoutes(s.Echo)
	RegisterWebRoutes(s.Echo)
}

func (s *RoutesSuite) TestRegisterRoutes() {
	routes := s.Echo.Routes()

	if assert.Equal(s.T(), len(routes), 3) {
		expected := []string{"/api/healthcheck", "/api/double", "/*"}
		for _, route := range routes {
			assert.Contains(s.T(), expected, route.Path)
		}
	}
}

func TestRunRoutesSuite(t *testing.T) {
	suite.Run(t, new(RoutesSuite))
}
