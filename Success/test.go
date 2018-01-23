package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/bouk/monkey"
)

func main() {
	var d *http.Client
	monkey.PatchInstanceMethod(reflect.TypeOf(d), "Get", func(_ *http.Client, _ string) (resp *http.Response, err error) {
		return nil, fmt.Errorf("no dialing allowed")
	})
	_, err := http.Get("http://google.com")
	fmt.Println(err) // Get http://google.com: no dialing allowed
}
