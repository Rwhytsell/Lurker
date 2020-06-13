package main

import (
	"flag"
	"log"
	"math/rand"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "irc-ws.chat.twitch.tv", "http service address")
var channel = flag.String("chan", "summit1g", "Target IRC Channel")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	username := "justinfan" + strconv.Itoa(rand.Intn(49999 - 10000) + 10000)
	log.Printf("Connecting as %s", username)
	c.WriteMessage(websocket.TextMessage, []byte("CAP REQ :twitch.tv/tags twitch.tv/commands"))
	c.WriteMessage(websocket.TextMessage, []byte("PASS SCHMOOPIIE"))
	c.WriteMessage(websocket.TextMessage, []byte("NICK " + username))
	c.WriteMessage(websocket.TextMessage, []byte("USER " + username + " 8 * :" + username))

	log.Printf("Joining channel #%s", *channel)
	c.WriteMessage(websocket.TextMessage, []byte("JOIN #" + *channel))

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
