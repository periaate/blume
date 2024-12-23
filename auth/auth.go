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
	"github.com/periaate/blume/yap"
)

func NewManager(opts ...func(*Manager)) *Manager {
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

func (s *Session) Encode(w io.Writer) error { return json.NewEncoder(w).Encode(&s) }

func (s *Session) Decode(r io.Reader) error { return json.NewDecoder(r).Decode(&s) }

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
	return m.Sessions.Set(s.Cookie, s, time.Until(s.T)).Ok
}

func (m *Manager) NewLink(uses int, label string, host string, duration time.Duration) (key string, ok bool) {
	if uses <= 0 {
		return
	}
	yap.Info("creating link", "label", label, "host", host, "duration", duration, "uses", uses)
	key = RandKey(32)
	ok = m.Links.Set(key, Link{
		Label:    label,
		Host:     host,
		Key:      key,
		T:        time.Now().Add(duration),
		Uses:     uses,
		Duration: duration,
	}, duration).Ok

	return
}

// UseLink uses a link to generate a session.
func (m *Manager) UseLink(key string, w http.ResponseWriter) (sess Session, ok bool) {
	m.mut.Lock()
	defer m.mut.Unlock()
	opt := m.Links.Get(key)
	if !opt.Ok {
		yap.Error("link not found", "key", key)
		return
	}

	link := opt.Value

	link.Uses--
	if link.Uses <= 0 {
		yap.Info("link expired", "key", key)
		m.Links.Del(key)
	} else {
		if !m.Links.Set(key, link, time.Until(link.T)).Ok {
			yap.Error("error updating link", "key", key)
			m.Links.Del(key)
			return // this should return ? Write tests dummy!!
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
		Domain:  link.Host,
		Expires: time.Now().Add(link.Duration),
	})

	m.Sessions.Set(sess.Cookie, sess, time.Until(sess.T))
	return
}
