package main

import (
	"flag"
	"fmt"
	"github.com/nats-io/nats"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"time"
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
	nc.Subscribe("nqjobs", func(m *nats.Msg) {
		fmt.Println("Time now: ", time.Now().Format("15:04:05.000"))
		url := string(m.Data)
		fmt.Println(string(m.Data))
		if strings.HasPrefix(url, "http") {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(body))
		}
	})
	runtime.Goexit()
}
