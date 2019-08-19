package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
//App encapsulates Env,Router and middleware
type App struct {
	Router *mux.Router
}
type shortenReq struct {
	Url string `json:"url" validate:"nonzero"`
	ExpirationInMinutes int64 `json:"expiration_in_minutes" validate:"min=0"`
}
type  shortlinkResp struct {
	ShortLink string `json:"short_link"`
}
//Initialize is initialization of app
func (a *App) Initialize()  {
	//set log formatter
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	a.Router = mux.NewRouter()
	a.InitializeRouters()
}
func (a *App) InitializeRouters() {
	a.Router.HandleFunc("/api/shorten",a.createShortLink,).Methods("POST")
	a.Router.HandleFunc("/api/info",a.getShortLinkInfo,).Methods("GEt")
	a.Router.HandleFunc("/{shortLink:[a-zA-Z0-9]{1,11}",a.Redirect,).Methods("GET")
}
func (a *App) createShortLink(w http.ResponseWriter,r *http.Request) {
	var req shortenReq
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&req)
	fmt.Println(req)
	if  err != nil{
		return
	}
	//arr := validator.Validate(req)
	//if arr != nil {
	//	return
	//}
	defer r.Body.Close()
	fmt.Println(11)
	fmt.Println(req)
}
func (a *App) getShortLinkInfo(w http.ResponseWriter,r *http.Request) {
	vale :=r.URL.Query()
	s := vale.Get("ShortLink")
	fmt.Println(s)
}
func (a *App) Redirect(w http.ResponseWriter,r *http.Request) {
	vale :=mux.Vars(r)
	fmt.Println(vale)
}
// run  starts listen and server
func (a *App) Run(addr string)  {
	_ = http.ListenAndServe(addr, a.Router)
}
func main() {
	a := App{}
	a.Initialize()
	a.Run("127.0.0.1:8000")
}