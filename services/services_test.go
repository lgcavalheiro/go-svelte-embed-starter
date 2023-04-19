package services

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func HandleErrorMsg(expected, received interface{}) string {
	return fmt.Sprintf("[ERROR] Expected %v but got %v\n", expected, received)
}

func TestHealthcheck(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	Healthcheck(rec, req)

	if http.StatusOK != rec.Result().StatusCode {
		t.Error(HandleErrorMsg(http.StatusOK, rec.Result().StatusCode))
	}
	if rec.Body.String() != "OK" {
		t.Error(HandleErrorMsg("OK", rec.Body.String()))
	}
}

func TestDouble(t *testing.T) {
	q := make(url.Values)
	q.Set("number", "2")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
	rec := httptest.NewRecorder()

	Double(rec, req)

	if http.StatusOK != rec.Code {
		t.Error(http.StatusOK, rec.Code)
	}
	if rec.Body.String() != "{\"result\":4}\n" {
		t.Error("{\"result\":4}\n", rec.Body.String())
	}
}

func TestDoubleFailure(t *testing.T) {
	q := make(url.Values)
	q.Set("number", "NAN")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
	rec := httptest.NewRecorder()

	Double(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Error(http.StatusBadRequest, rec.Code)
	}
	if rec.Body.String() != "\"param number was invalid\"\n" {
		t.Error("\"param number was invalid\"\n", rec.Body.String())
	}
}
