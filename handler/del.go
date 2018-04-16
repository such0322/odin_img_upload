package handler

import (
	"log"
	"net/http"
	"odin_img_upload/config"
	"os"
)

func DelNews(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	news := r.FormValue("news")
	stat, err := os.Stat(config.GlobalConf.NewsPath + news)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	if stat.IsDir() {
		os.RemoveAll(config.GlobalConf.NewsPath + news)
	}
	http.Redirect(w, r, "/", http.StatusFound)

}

func DelImg(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("p")
	file := r.FormValue("f")
	imgPath := config.GlobalConf.NewsPath + path + "/" + file
	err := os.Remove(imgPath)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/newsImgs?news="+path, http.StatusFound)
}
