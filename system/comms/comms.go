package comms

import (
	"fmt"
	"net/http"
	"os"

	"blume/clog"

	ws "github.com/gorilla/websocket"
)

var addr = "localhost:8888"

func init() {
	a := os.Getenv("BLUME_COMMS_ADDR")
	if len(a) != 0 {
		addr = a
	}
}

func Dial(mode, arg string) (c *Client, err error) {
	c = &Client{}
	d := ws.Dialer{}
	URL := fmt.Sprintf("ws://%s/%s/%s/", addr, mode, arg)
	conn, _, err := d.Dial(URL, nil)
	if err != nil {
		return
	}

	c.Conn = conn
	return
}

// Client
type Client struct{ Conn *ws.Conn }

func (c *Client) Call(data []byte) (err error) { return c.Conn.WriteMessage(ws.TextMessage, data) }

// Broker
type Broker struct{ Spaces map[string]*ws.Conn }

var upgrader = ws.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func (b *Broker) Connect(w http.ResponseWriter, r *http.Request) {
	space := r.PathValue("space")
	if len(space) == 0 {
		return
	}
	clog.Debug("connection request", "name", space)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		err = fmt.Errorf("could not open websocket connection")
		clog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bConn, ok := b.Spaces[space]
	if !ok {
		return
	}

	go func() {
		for {
			if messageType, p, err := conn.ReadMessage(); err == nil {
				bConn.WriteMessage(messageType, p)
			}
		}
	}()
}

func (b *Broker) Register(w http.ResponseWriter, r *http.Request) {
	reg := r.PathValue("name")
	if len(reg) == 0 {
		return
	}
	clog.Debug("register request", "name", reg)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		err = fmt.Errorf("could not open websocket connection")
		clog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, ok := b.Spaces[reg]
	if ok {
		return
	}

	b.Spaces[reg] = conn
}

func (b *Broker) Mux(mux *http.ServeMux) {
	mux.HandleFunc("/connect/{space}/", b.Connect)
	mux.HandleFunc("/register/{name}/", b.Register)
}

func Host() {
	mux := &http.ServeMux{}
	b := &Broker{make(map[string]*ws.Conn)}
	b.Mux(mux)
	http.ListenAndServe(addr, mux)
}
