package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/mock"
)

func TestService(t *testing.T) {
	str := `<html><body><h1>Hello World</h1></body></html>`

	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, str)
	}

	// ugly hack
	registry.DefaultRegistry = registry.Registry(mock.NewRegistry())

	service := NewService(
		Name("go.micro.web.test"),
	)

	service.HandleFunc("/", fn)

	go func() {
		if err := service.Run(); err != nil {
			t.Fatal(err)
		}
	}()

	// another ugly hack
	time.Sleep(time.Millisecond * 100)

	s, err := registry.GetService("go.micro.web.test")
	if err != nil {
		t.Fatal(err)
	}

	rsp, err := http.Get(fmt.Sprintf("http://%s:%d", s[0].Nodes[0].Address, s[0].Nodes[0].Port))
	if err != nil {
		t.Fatal(err)
	}
	defer rsp.Body.Close()

	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(b) != str {
		t.Errorf("Expected %s got %s", str, string(b))
	}
}
