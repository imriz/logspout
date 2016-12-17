package tcp

import (
	"net"
	"log"
	"time"
	"github.com/gliderlabs/logspout/adapters/raw"
	"github.com/gliderlabs/logspout/router"
)

func init() {
	router.AdapterTransports.Register(new(tcpTransport), "tcp")
	// convenience adapters around raw adapter
	router.AdapterFactories.Register(rawTCPAdapter, "tcp")
}

func rawTCPAdapter(route *router.Route) (router.LogAdapter, error) {
	route.Adapter = "raw+tcp"
	return raw.NewRawAdapter(route)
}

type tcpTransport int

func (_ *tcpTransport) Dial(addr string, options map[string]string) (net.Conn, error) {
	log.Printf("tcp: Resolving %v\n", addr)
	raddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	log.Printf("tcp: Connecting to %v\n", raddr)
	conn, err := net.DialTimeout("tcp", addr,time.Duration(10)*time.Second)
	if err != nil {
	        log.Printf("tcp: Error connecting to %v\n", raddr)
		return nil, err
	}
	log.Printf("tcp: Connected to %v\n", raddr)
	return conn, nil
}
