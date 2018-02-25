package server

import (
	"testing"
	"net/http"
	"fmt"
	"time"
	"runtime"
)

func TestProxy_Run(t *testing.T) {
	testServer()
	ch := make(chan struct{})
	p := Proxy{

	}
	go p.Run(ch)

	httpCli := http.Client{
		Timeout: time.Second,
	}
	a, err := httpCli.Get("http://127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(runtime.NumGoroutine())
	var buf = make([]byte, 10)
	a.Body.Read(buf)

	fmt.Printf("get %s\n", buf)

	a.Body.Close()

	<-ch
}

func testServer() {
	var handle http.HandlerFunc = func(rw http.ResponseWriter, request *http.Request) {
		fmt.Fprint(rw, "hello")
	}

	go func() {
		http.ListenAndServe("127.0.0.1:6464", handle)
	}()
}
