package api

import (
	"fmt"
	"net/http"

	"github.com/tomjamescn/siri/registry"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var (
		service = r.URL.Query().Get("service")
	)

	if h, ok := registry.Center[service]; ok {
		h.ServeHTTP(w, r)
	} else {
		fmt.Fprintf(w, "Not Support")
	}
}
