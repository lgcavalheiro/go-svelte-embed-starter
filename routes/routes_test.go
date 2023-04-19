package routes

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestRegisterRoutes(t *testing.T) {
	RegisterApiRoutes()
	RegisterWebRoutes()

	routes := reflect.ValueOf(http.DefaultServeMux).Elem().FieldByIndex([]int{1}).MapRange()

	expected := []string{"/api/healthcheck", "/api/double", "/"}
	for routes.Next() {
		strings.Contains(strings.Join(expected, " "), routes.Key().String())
	}
}
