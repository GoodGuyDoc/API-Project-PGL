package session

import (
	"github.com/gorilla/sessions"
)

var (
	Key   = []byte("super-secret-key") // unguessable. very secure.
	Store = sessions.NewCookieStore(Key)
)
