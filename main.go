package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"golang.org/x/net/http2"
)

var addr = flag.String("addr", "localhost:8004", "http service address")

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Nice nice nice ")

	//w.Write([]byte("test"))

	//homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func index_main(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1><center> Hello from Go! </h1></center>"))

}

func main() {

	var srv http.Server
	srv.Addr = ":8002"
	//Enable http2
	http2.ConfigureServer(&srv, nil)

	http.HandleFunc("/", index_main)

	err := srv.ListenAndServeTLS("certs/localhost.cert", "certs/localhost.key")

	if err != nil {
		fmt.Println(err)
	}

	//srv.ListenAndServe()
	// flag.Parse()
	// log.SetFlags(0)
	// http.HandleFunc("/echo", echo)
	// http.HandleFunc("/", home)
	// log.Fatal(http.ListenAndServe(*addr, nil))
}

// openssl genrsa -out localhost.key 2048
// openssl req -new -x509 -key localhost.key -out localhost.cert -days 3650 -subj /CN=localhost
