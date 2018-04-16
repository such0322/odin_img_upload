package handler

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"odin_img_upload/config"
	"strings"
)

func NewsImgs(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	news := r.FormValue("news")
	path := config.GlobalConf.NewsPath + news
	list, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
		return
	}
	var imgs []string
	for _, vo := range list {
		if !strings.HasPrefix(vo.Name(), ".") {
			imgs = append(imgs, vo.Name())
		}
	}
	var data struct {
		Imgs []string
		News string
	}
	data.Imgs = imgs
	data.News = news
	tpl := template.Must(template.New("imgs.html").ParseFiles("templates/imgs.html"))
	tpl.Execute(w, data)
}
