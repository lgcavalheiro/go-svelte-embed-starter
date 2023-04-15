package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type DoubleResult struct {
	Result int `json:"result"`
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func Double(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	number, err := strconv.Atoi(r.URL.Query().Get("number"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("param number was invalid")
		return
	}

	json.NewEncoder(w).Encode(&DoubleResult{Result: number * 2})
}
