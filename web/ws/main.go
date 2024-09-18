package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func wsbase(w http.ResponseWriter, r *http.Request, h func(c *websocket.Conn)) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		h(c)
	}
}

func handleTest(c *websocket.Conn) {
	mt, message, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
	}
	log.Printf("recv: %s", message)
	err = c.WriteMessage(mt, message)
	if err != nil {
		log.Println("write:", err)
	}
}

func Test(w http.ResponseWriter, r *http.Request) {
	wsbase(w, r, handleTest)
}
