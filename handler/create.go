package handler

import (
	"io"
	"log"
	"net/http"
	"odin_img_upload/config"
	"os"
)

func Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newsName := r.FormValue("new_news")
	err := os.Mkdir(config.GlobalConf.NewsPath+newsName, os.ModePerm)
	if err != nil {
		log.Println(err)
		w.Write([]byte("news已存在"))
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)

}

func CreateImg(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(16 << 20)

	img, imgInfo, err := r.FormFile("img")
	if err != nil {
		log.Println(err)
		w.Write([]byte("上传失败"))
		return
	}
	defer img.Close()
	contentType := imgInfo.Header["Content-Type"][0]
	if contentType != "image/jpeg" && contentType != "image/png" {
		w.Write([]byte("文件格式错误，只允许上传jpg和png格式图片"))
		return
	}
	path := r.FormValue("p")
	imgPath := config.GlobalConf.NewsPath + path + "/" + imgInfo.Filename
	_, err = os.Stat(imgPath)
	if !os.IsNotExist(err) {
		w.Write([]byte("文件已存在"))
		return
	}

	file, err := os.Create(imgPath)
	if err != nil {
		log.Println(err)
		w.Write([]byte("创建文件失败"))
		return
	}
	defer file.Close()

	io.Copy(file, img)
	http.Redirect(w, r, "/newsImgs?news="+path, http.StatusFound)

}
