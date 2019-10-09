package test

import (
	"../noodle"
	"fmt"
	"github.com/tungyao/cedar"
	"net/http"
	"testing"
)

func TestNoodle(t *testing.T) {

	r:=cedar.NewRouter()
	r.Get("/index", noodle.Hello(func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, "hello world")
	}))
	_ = http.ListenAndServe(":80", r)
}
