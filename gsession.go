package gosession

import (
	"time"
)

var (
	Domain      string = "localhost"
	Path        string = "/"
	HttpOnly    bool   = false
	MaxAge      int    = 3600
	GoSessionId string = "gosession-id"
)

type GoSession struct {
	SId  string
	data map[string]interface{}
}

type GoSessionManager struct {
	gosessions  map[string]*GoSession
	CreateAts   map[string]time.Time
	GCInterval  time.Duration
	GoSessionId string
	MaxAge      int
	Domain      string
	HttpOnly    bool
	Path        string
}

func NewGoSessionManager(gosessionsId) *GoSessionManager {
	return &GoSessionManager{
		gosessions:  make(map[string]*GoSession),
		CreateAts:   make(map[string]time.Time),
		MaxAge:      MaxAge,
		Domain:      Domain,
		HttpOnly:    HttpOnly,
		Path:        Path,
		GoSessionId: GoSessionId,
		GCInterval:  time.Millisecond * 10,
	}
}

func StartGC(gms *GoSessionManager) {

}
