package gosession

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
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

type SyncGoSession interface {
	UpdateSession(response http.ResponseWriter, request *http.Request)
}
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

func NewGoSessionManager(gosessionsId string, gcInterval time.Duration) *GoSessionManager {
	if gosessionsId == "" {
		gosessionsId = GoSessionId
	}

	if gcInterval <= 0 {
		gcInterval = time.Millisecond * 100
	}
	return &GoSessionManager{
		gosessions:  make(map[string]*GoSession),
		CreateAts:   make(map[string]time.Time),
		MaxAge:      MaxAge,
		Domain:      Domain,
		HttpOnly:    HttpOnly,
		Path:        Path,
		GoSessionId: gosessionsId,
		GCInterval:  gcInterval,
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

	var bys = make([]byte, 32)
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
	gsm.CreateAts[cValue] = time.Now().Add(time.Second * time.Duration(gsm.MaxAge))

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

func (gs *GoSession) Set(key string, va interface{}) {

	gs.data[key] = va

}

func (gs *GoSession) Delete(key string) {

	delete(gs.data, key)
}

//垃圾收集处理过期的 GoSession
func StartGC(gsm *GoSessionManager) {

	go func() {

		for _ = range time.Tick(gsm.GCInterval) {
			fmt.Println("************ GC **************")
			for k, v := range gsm.CreateAts {

				if gsm.expiredSession(v) == false {

					gsm.Delete(k)
				}
			}
		}
	}()

}

//判断是否过期
func (gsm *GoSessionManager) expiredSession(at time.Time) bool {

	return time.Now().Before(at)
}

//根据key删除gosession
func (gsm *GoSessionManager) Delete(key string) {
	gsm.Lock()
	defer gsm.Unlock()

	delete(gsm.CreateAts, key)
	delete(gsm.gosessions, key)

}

// 更新gosession
func (gsm *GoSessionManager) UpdateGoSession(response http.ResponseWriter, request *http.Request) {
	currentCookie, err := request.Cookie(gsm.GoSessionId)

	var key string
	if err != nil || currentCookie == nil {
		return
	}

	key = currentCookie.Value
	_, ok := gsm.CreateAts[key]
	if ok {
		fmt.Println("=======update goSession ===========")
		gsm.CreateAts[key] = time.Now().Add(time.Second * time.Duration(gsm.MaxAge))
		fmt.Println("at ", gsm.CreateAts[key], "key=", key)
		currentCookie.MaxAge = gsm.MaxAge

		http.SetCookie(response, currentCookie)
	}

}
