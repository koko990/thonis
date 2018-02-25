package server

import (
	"net"
	"github.com/thonis/pkg/log"
	"net/http"
	"net/url"
	"net/http/httputil"
	"fmt"
)

type Proxy struct {
	SourceHost      net.IP
	SourcePort      int
	DestinationHost net.IP
	DestinationPort int
}

func (p *Proxy) Run(ch chan<- struct{}) {
	log.Infoln(*p)
	startServer()
	ch <- struct{}{}
}

type handle struct {
	host string
	port string
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse("http://" + this.host + ":" + this.port)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

type Ruler interface {
}

func startServer() {

	//被代理的服务器host和port
	h := &handle{host: "127.0.0.1", port: "6464"}
	s := http.Server{
		Addr:    ":8888",
		Handler: h,
		ConnState: func(conn net.Conn, state http.ConnState) {
			fmt.Println(conn.RemoteAddr())
			fmt.Println(state.String())
		},
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("dddddd")
}
