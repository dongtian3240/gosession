package gosession

import (
	"net/http"
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

//获取gosession
func (gsm *GoSessionManager) getGoSession(response http.ResponseWriter, request *http.Request) *GoSession {

	currentCookie, err := request.Cookie(gsm.GoSessionId)
	if err != nil || currentCookie == nil { //创建一个新的gosession

	} else {

	}

}

func StartGC(gms *GoSessionManager) {

}
