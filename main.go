package main

import (
	"flag"
	"fmt"
	"github.com/nats-io/nats"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"sync"
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
	var mutex = &sync.Mutex{}
	//	nc.Publish("nqjobs", []byte("http://www.netroby.com")) //To create a nqjobs
	nc.Subscribe("nqjobs", func(m *nats.Msg) {
		mutex.Lock()
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
		defer mutex.Unlock()
		runtime.Gosched()
	})
	runtime.Goexit()
}
