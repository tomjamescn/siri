package api

import (
	"fmt"
	"net/http"

	_ "github.com/tomjamescn/siri-whatismyip"
	"github.com/tomjamescn/siri/registry"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var (
		service = r.URL.Query().Get("service")
	)

	if h, ok := registry.Center[service]; ok {
		h.ServeHTTP(w, r)
	} else {
		fmt.Fprintf(w, "Not Support. service: %s", service)
	}
}
