package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	f, e := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" +
		strings.Split(r.URL.EscapedPath(), "/")[2]) // strings.Split会将URL拆分成三个字符串（" ", "objects", "object_name"），获得到文件名object_name
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	// 如果成功则将r.Body写入到文件f当中
	io.Copy(f, r.Body)
}
