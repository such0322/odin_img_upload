package main

import (
	"log"
	"net/http"
	"odin_img_upload/config"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	if !config.GlobalConf.Logger {
		return inner
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf("%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start))
	})
}
