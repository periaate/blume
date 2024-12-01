package auth

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/periaate/blume/clog"
	"github.com/periaate/blume/gen"
	"github.com/periaate/blume/maps"
)

func NewManager(opts ...gen.Option[Manager]) *Manager {
	m := &Manager{
		Links:    maps.NewExpiring[string, Link](),
		Sessions: maps.NewExpiring[string, Session](),
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

type Session struct {
	Cookie string    `json:"cookie"`
	Label  string    `json:"label"`
	Host   string    `json:"host"`
	T      time.Time `json:"expires"`
}

func (s *Session) Encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(&s)
}

func (s *Session) Decode(r io.Reader) error {
	return json.NewDecoder(r).Decode(&s)
}

func (s *Session) Reader() io.Reader {
	rw := bytes.NewBuffer([]byte{})
	s.Encode(rw)
	return rw
}

type Link struct {
	Key   string
	Host  string
	Label string
	// T is the time when the link will expire
	T time.Time
	// Uses is the number of times the link can be used
	Uses int

	// Duration is the duration of generated sessions
	Duration time.Duration
}

func RandKey(length int) string {
	val := make([]byte, length)
	rand.Read(val)
	return base64.URLEncoding.EncodeToString(val)
}

type Manager struct {
	Links    *maps.Expiring[string, Link]
	Sessions *maps.Expiring[string, Session]
	mut      sync.Mutex
}

func (m *Manager) Register(s Session) (ok bool) {
	m.mut.Lock()
	defer m.mut.Unlock()
	m.Sessions.Sync.Values[s.Cookie] = maps.ExpItem[Session]{
		Value:   s,
		Expires: s.T,
	}
	return true
}

func (m *Manager) NewLink(uses int, label string, host string, duration time.Duration) (key string, ok bool) {
	if uses <= 0 {
		return
	}
	clog.Info("creating link", "label", label, "host", host, "duration", duration, "uses", uses)
	key = RandKey(32)
	ok = m.Links.Set(key, Link{
		Label:    label,
		Host:     host,
		Key:      key,
		T:        time.Now().Add(duration),
		Uses:     uses,
		Duration: duration,
	}, time.Now().Add(duration))

	return
}

// UseLink uses a link to generate a session.
func (m *Manager) UseLink(key string, w http.ResponseWriter) (sess Session, ok bool) {
	m.mut.Lock()
	defer m.mut.Unlock()
	link, ok := m.Links.Get(key)
	if !ok {
		clog.Error("link not found", "key", key)
		return
	}

	link.Uses--
	if link.Uses <= 0 {
		clog.Info("link expired", "key", key)
		m.Links.Del(key)
	} else {
		if !m.Links.Set(key, link, link.T) {
			clog.Error("error updating link", "key", key)
			m.Links.Del(key)
		}
	}

	cookie := RandKey(32)
	sess = Session{
		Cookie: cookie,
		Label:  link.Label,
		Host:   link.Host,
		T:      time.Now().Add(link.Duration),
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "X-Session",
		Value:   cookie,
		Expires: time.Now().Add(link.Duration),
	})

	fmt.Println("ABC")
	m.Sessions.Set(sess.Cookie, sess, sess.T)
	fmt.Println("DEF")
	return
}
