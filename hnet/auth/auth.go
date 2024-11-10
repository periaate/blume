package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/periaate/blume/clog"
)

func NewManager() *Manager {
	return &Manager{
		Links:    make(map[string]*Link),
		Sessions: make(map[string]*Session),
	}
}

type Session struct {
	Key string
	T   time.Time
}

type Link struct {
	Key string
	// T is the time when the link will expire
	T time.Time
	// Uses is the number of times the link can be used
	Uses int

	// Duration is the duration of generated sessions
	Duration time.Duration
}

func (l *Link) Use(w http.ResponseWriter) (sess *Session, err error) {
	if l.Uses <= 0 {
		err = fmt.Errorf("link has no uses")
		return
	}
	if isExpired(l.T) {
		err = fmt.Errorf("link has expired")
		return
	}

	l.Uses--

	key := RandKey(32)
	sess = &Session{
		Key: key,
		T:   time.Now().Add(l.Duration),
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "X-Session",
		Value:   key,
		Expires: sess.T,
	})

	return
}

func isExpired(t time.Time) bool { return t.Before(time.Now()) }

func RandKey(length int) string {
	val := make([]byte, length)
	rand.Read(val)
	return base64.URLEncoding.EncodeToString(val)
}

type Manager struct {
	Links    map[string]*Link
	Sessions map[string]*Session
	mut      sync.Mutex
}

func (m *Manager) GetSession(key string) (sess *Session, ok bool) {
	m.mut.Lock()
	defer m.mut.Unlock()

	sess, ok = m.Sessions[key]
	if isExpired(sess.T) {
		delete(m.Sessions, key)
		ok = false
	}

	return
}

func (m *Manager) NewLink(uses int, duration time.Duration) (key string) {
	m.mut.Lock()
	defer m.mut.Unlock()

	key = RandKey(32)
	m.Links[key] = &Link{
		Key:      key,
		T:        time.Now().Add(duration),
		Uses:     uses,
		Duration: duration,
	}

	return
}

func (m *Manager) Stringify() {
	m.mut.Lock()
	defer m.mut.Unlock()
	clog.Info("found sessions", "len", len(m.Sessions))
	for _, v := range m.Sessions {
		clog.Info("session found", "key", v)
	}
}

func (m *Manager) UseLink(key string, w http.ResponseWriter) (sess *Session, err error) {
	m.mut.Lock()
	defer m.mut.Unlock()

	link, ok := m.Links[key]
	if !ok {
		err = fmt.Errorf("link not found")
		return
	}

	sess, err = link.Use(w)
	if err != nil {
		delete(m.Links, key)
	}

	m.Sessions[sess.Key] = sess

	return
}

func (m *Manager) IsValidSession(key string) bool {
	m.mut.Lock()
	defer m.mut.Unlock()

	sess, ok := m.Sessions[key]
	if !ok {
		return false
	}

	if isExpired(sess.T) {
		delete(m.Sessions, key)
		return false
	}

	return true
}
