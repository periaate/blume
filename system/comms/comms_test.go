package comms

import (
	"log/slog"
	"testing"
	"time"

	"github.com/periaate/blume/clog"
)

func registerSpace(space string) {
	c, err := Dial("register", space)
	if err != nil {
		return
	}
	conn := c.Conn

	for {
		if _, p, err := conn.ReadMessage(); err == nil {
			clog.Info("msg received", "msg", string(p))
		}
	}
}

func TestCall(t *testing.T) {
	clog.SetLogLoggerLevel(slog.LevelDebug)
	go Host()
	go registerSpace("space")
	time.Sleep(time.Millisecond * 100)

	c, err := Dial("connect", "space")
	if err != nil {
		t.Fatal(err)
	}

	c.Call([]byte("hello world"))
	time.Sleep(time.Millisecond * 10)
}
