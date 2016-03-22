# gosession

provide  a session  for go web  in memory !


example :

package main

import (
	"fmt"
	"net/http"
)

import "gosession"

var gsm *gosession.GoSessionManager

func main() {

	gsm = gosession.NewGoSessionManager("go--session--id")

	http.HandleFunc("/session", gsessionHanlder)
	http.ListenAndServe(":8087", nil)
}

func gsessionHanlder(response http.ResponseWriter, request *http.Request) {

	gs := gsm.GetGoSession(response, request)
	response.Write([]byte(gs.SId))
	gs.Set("name", "半夏")
	va := gs.Get("name")
	fmt.Println(va)
	gs.Delete("name")
	fmt.Println("va= ", gs.Get("name"))
}


the output :
半夏
va=  "<nil>"

the broswer cookie  key is go--session--id


