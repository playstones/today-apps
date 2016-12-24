package main

import (
	"net/http"

	"today-apps/model"

	"encoding/json"

	"github.com/gorilla/mux"
)

func main() {
	listenServer()
}

func listenServer() {
	r := mux.NewRouter()
	r.HandleFunc("/api/apps", apps).Methods("GET")
	http.Handle("/api/apps", r)
	//http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	http.ListenAndServe(":8080", nil)
}

func apps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	apps := allApps()

	res, err := json.Marshal(apps)

	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	w.Write(res)
}

func allApps() []model.App {
	var a model.App
	a.Identifier = "today.app.wechat.scanqrcode"
	a.Title = "微信扫一扫"
	a.URL = "weixin://scanqrcode"
	a.DisplayName = "微信扫一扫"
	a.IconURL = ""

	apps := []model.App{a}
	return apps
}
