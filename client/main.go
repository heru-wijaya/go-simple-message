package main

import (
	"bufio"
	"flag"
	"fmt"
	"go-simple-message/controller"
	"math/rand"
	"os"
	"time"

	"golang.org/x/net/websocket"
)

var (
	port = flag.String("port", "9000", "port used for ws connection")
)

func main() {
	flag.Parse()

	// connect
	ws, err := connect()
	if err != nil {
		fmt.Println(err)
	}
	defer ws.Close()

	// receive
	var m controller.CreateMessageInput
	go func() {
		for {
			err := websocket.JSON.Receive(ws, &m)
			if err != nil {
				fmt.Println("Error receiving message: ", err.Error())
				break
			}
			fmt.Println("Message: ", m)
		}
	}()

	// send
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		m := controller.CreateMessageInput{
			Body: text,
		}
		err = websocket.JSON.Send(ws, m)
		if err != nil {
			fmt.Println("Error sending message: ", err.Error())
			break
		}
	}
}

// connect connects to the local chat server at port <port>
func connect() (*websocket.Conn, error) {
	return websocket.Dial(fmt.Sprintf("ws://localhost:%s", *port), "", mockIP())
}

// mockedIP is a demo-only utility that generates a random IP address for this client
func mockIP() string {
	var arr [4]int
	for i := 0; i < 4; i++ {
		rand.Seed(time.Now().UnixNano())
		arr[i] = rand.Intn(256)
	}
	return fmt.Sprintf("http://%d.%d.%d.%d", arr[0], arr[1], arr[2], arr[3])
}
