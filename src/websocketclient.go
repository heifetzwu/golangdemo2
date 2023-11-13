package main

import (
	"log"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

// func main() {
// 	websocketclient()
// }

func websocketclient() {
	c, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8899/echo", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	for i := 0; i < 2; i++ {
		msgstr := "ithome30day#" + strconv.Itoa(i)

		err = c.WriteMessage(websocket.TextMessage, []byte(msgstr))
		log.Printf("clinet write: %s\n", msgstr)
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(10 * time.Second)
	}

	time.Sleep(5 * time.Second)
	_, msg, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	log.Printf("clinet receive: %s\n", msg)
}
