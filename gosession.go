package gosession

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"sync"
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
	sync.Mutex
}

func NewGoSessionManager(gosessionsId string) *GoSessionManager {
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
func (gsm *GoSessionManager) GetGoSession(response http.ResponseWriter, request *http.Request) *GoSession {

	currentCookie, err := request.Cookie(gsm.GoSessionId)

	//创建一个新的gosession
	if err != nil || currentCookie == nil {
		return gsm.createGoSession(response, request)
	} else {
		//当前gosession是否存在
		goSession, has := gsm.hasGoSession(currentCookie.Value)
		if has { // 如果存在将直接返回
			return goSession
		} else {

			return gsm.createGoSession(response, request)
		}

	}

}

// 查询是否已经存在
func (gsm *GoSessionManager) hasGoSession(key string) (*GoSession, bool) {

	goSesion, ok := gsm.gosessions[key]

	return goSesion, ok
}

//创建一个新的gosession
func (gsm *GoSessionManager) createGoSession(response http.ResponseWriter, request *http.Request) *GoSession {

	var bys = make([]byte, 512)
	rand.Read(bys)
	cValue := base64.StdEncoding.EncodeToString(bys)

	c := &http.Cookie{
		Name:     gsm.GoSessionId,
		Value:    cValue,
		Path:     gsm.Path,
		Domain:   gsm.Domain,
		MaxAge:   gsm.MaxAge,
		HttpOnly: gsm.HttpOnly,
	}

	http.SetCookie(response, c)
	goSession := &GoSession{
		SId:  cValue,
		data: make(map[string]interface{}),
	}
	gsm.gosessions[cValue] = goSession
	gsm.CreateAts[cValue] = time.Now().Add(time.Duration(gsm.MaxAge))

	return goSession
}

func (gs *GoSession) Get(key string) interface{} {

	va, ok := gs.data[key]
	if ok {
		return va
	} else {
		return nil
	}
}

func (gs *GoSession) Delete(key string) {

	delete(gs.data, key)
}

//垃圾收集处理过期的 GoSession
func StartGC(gsm *GoSessionManager) {

	//
	//

	go func() {

	}()

}
