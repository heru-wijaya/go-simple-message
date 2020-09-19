package main

import (
	"flag"
	"fmt"
	"go-simple-message/controller"
	"net/http"

	"golang.org/x/net/websocket"
)

type channel struct {
	clients          map[string]*websocket.Conn
	addClientChan    chan *websocket.Conn
	removeClientChan chan *websocket.Conn
	broadcastChan    chan controller.CreateMessageInput
}

var (
	port = flag.String("port", "9000", "port used for ws connection")
)

func main() {
	flag.Parse()
	fmt.Println(server(*port))
}

// server creates a websocket server at port <port> and registers the sole handler
func server(port string) error {
	h := newChannel()
	mux := http.NewServeMux()
	mux.Handle("/", websocket.Handler(func(ws *websocket.Conn) {
		handler(ws, h)
	}))

	s := http.Server{Addr: ":" + port, Handler: mux}
	return s.ListenAndServe()
}

// handler registers a new chat client conn;
// It runs the channel, adds the client to the connection pool
// and broadcasts received message
func handler(ws *websocket.Conn, c *channel) {
	go c.run()

	c.addClientChan <- ws

	for {
		var m controller.CreateMessageInput
		err := websocket.JSON.Receive(ws, &m)
		if err != nil {
			c.broadcastChan <- controller.CreateMessageInput{err.Error()}
			c.removeClient(ws)
			return
		}
		c.broadcastChan <- m
	}
}

// newChannel returns a new channel
func newChannel() *channel {
	return &channel{
		clients:          make(map[string]*websocket.Conn),
		addClientChan:    make(chan *websocket.Conn),
		removeClientChan: make(chan *websocket.Conn),
		broadcastChan:    make(chan controller.CreateMessageInput),
	}
}

// run receives from the channels and calls the appropriate method
func (c *channel) run() {
	for {
		select {
		case conn := <-c.addClientChan:
			c.addClient(conn)
		case conn := <-c.removeClientChan:
			c.removeClient(conn)
		case m := <-c.broadcastChan:
			c.broadcastMessage(m)
		}
	}
}

// removeClient removes a conn from the pool
func (c *channel) removeClient(conn *websocket.Conn) {
	delete(c.clients, conn.LocalAddr().String())
}

// addClient adds a conn to the pool
func (c *channel) addClient(conn *websocket.Conn) {
	c.clients[conn.RemoteAddr().String()] = conn
}

// broadcastMessage sends a message to all client conns in the pool
func (c *channel) broadcastMessage(m controller.CreateMessageInput) {
	for _, conn := range c.clients {
		err := websocket.JSON.Send(conn, m)
		if err != nil {
			fmt.Println("Error broadcastMessage: ", err)
			return
		}
	}
}
