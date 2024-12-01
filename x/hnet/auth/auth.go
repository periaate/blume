package auth

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/periaate/blume/maps"
)

func NewManager() *Manager {
	return &Manager{
		Links:    maps.NewExpiring[string, Link](),
		Sessions: maps.NewExpiring[string, Session](),
	}
}

type Session struct {
	Cookie string    `json:"cookie"`
	Label  string    `json:"label"`
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
	return m.Sessions.Set(s.Cookie, s, s.T)
}

func (m *Manager) NewLink(uses int, label string, duration time.Duration) (key string, ok bool) {
	if uses <= 0 {
		return
	}
	key = RandKey(32)
	ok = m.Links.Set(key, Link{
		Label:    label,
		Key:      key,
		T:        time.Now().Add(duration),
		Uses:     uses,
		Duration: duration,
	}, time.Now().Add(duration))

	return
}

func (m *Manager) UseLink(key string, w http.ResponseWriter) (sess Session, ok bool) {
	link, ok := m.Links.Get(key)
	if !ok {
		return
	}

	cookie := RandKey(32)
	sess = Session{
		Cookie: cookie,
		Label:  link.Label,
		T:      time.Now().Add(link.Duration),
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "X-Session",
		Value:   cookie,
		Expires: time.Now().Add(link.Duration),
	})

	m.Sessions.Set(sess.Cookie, sess, sess.T)

	link.Uses--
	if link.Uses <= 0 {
		m.Links.Del(key)
		return
	}

	m.Links.Set(key, link, link.T)
	return
}
