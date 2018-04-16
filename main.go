package main

import (
	"log"
	"net/http"
	"odin_img_upload/config"
	"time"
)

func main() {
	config.Init()
	mux := NewRouter()
	mux.PathPrefix("/imgs/").Handler(http.StripPrefix("/imgs/", http.FileServer(http.Dir(config.GlobalConf.NewsPath))))
	svr := http.Server{
		Addr:         config.GlobalConf.Addr,
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(svr.ListenAndServe())
}
