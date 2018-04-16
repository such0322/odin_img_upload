package handler

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"odin_img_upload/config"
)

func Index(w http.ResponseWriter, r *http.Request) {
	//获取news目录
	newsPath := config.GlobalConf.NewsPath
	dir, err := ioutil.ReadDir(newsPath)
	if err != nil {
		fmt.Println(err)
	}
	var dirList []string
	for _, vo := range dir {
		if vo.IsDir() {
			dirList = append(dirList, vo.Name())
		}
	}
	tpl := template.Must(template.New("index.html").ParseFiles("templates/index.html"))
	tpl.Execute(w, dirList)
}
