package registry

import (
	"net/http"
)

var Center = map[string]http.Handler{}

func Register(serviceName string, h http.Handler) {
	Center[serviceName] = h
}
