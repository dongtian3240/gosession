# gosession

provide  a session support for go web  in memory !


example :



package main

import (
	"fmt"
	"net/http"
	"time"
)

import "gosession"

var gsm *gosession.GoSessionManager

func main() {

	gsm = gosession.NewGoSessionManager("go--session--id", time.Millisecond*100)
	gosession.StartGC(gsm)
	http.HandleFunc("/session", gsessionHanlder)
	http.ListenAndServe(":8087", nil)
}

func gsessionHanlder(response http.ResponseWriter, request *http.Request) {
	gsm.UpdateGoSession(response, request)
	gs := gsm.GetGoSession(response, request)
	response.Write([]byte(gs.SId))
	gs.Set("name", "半夏")
	va := gs.Get("name")
	fmt.Println(va)
	gs.Delete("name")
	fmt.Println("va= ", gs.Get("name"))
}


look the console !


