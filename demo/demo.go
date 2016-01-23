package main

import (
	"flag"
	"fmt"
	"github.com/nats-io/nats"
)

func main() {
	host := flag.String("host", "127.0.0.1", "a string")
	port := flag.Int("port", 4222, "a int")
	flag.Parse()
	fmt.Println(*host)
	fmt.Println(*port)
	fmt.Println(fmt.Sprintf("nats://%s:%d", *host, *port))
	nc, err := nats.Connect(fmt.Sprintf("nats://%s:%d", *host, *port))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success connected to nats server")
	}
	nc.Publish("nqjobs", []byte("http://www.netroby.com"))
	nc.Close()
}
