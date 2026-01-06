package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/info", infohandler)
	http.ListenAndServe(":8080", nil)
}

func infohandler(w http.ResponseWriter, r *http.Request) {
	jst, _ := time.LoadLocation("Asia/Tokyo")

	var h map[string][]string
	h = r.Header

	fmt.Fprintf(w, "今の時刻は%sで，利用しているブラウザは%s，ですね",
		time.Now().In(jst).Format("15:04"),
		h["User-Agent"][0])
}
