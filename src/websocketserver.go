package main

import (
	// 	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// func main() {
// 	websocketserver()
// }

func websocketserver() {
	upgrader := &websocket.Upgrader{
		//如果有 cross domain 的需求，可加入這個，不檢查 cross domain
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade:", err)
			return
		}
		defer func() {
			log.Println("disconnect !!")
			c.Close()
		}()
		for {
			mtype, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
				// return
			}
			log.Printf("## server receive: %s\n", string(msg))
			err = c.WriteMessage(mtype, msg)
			log.Println("## server write:", string(msg))
			if err != nil {
				log.Println("server write:", err)
				break
				// return
			}
		}
	})
	log.Println("server start at :8899")
	log.Fatal(http.ListenAndServe(":8899", nil))

}
