package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"os"
	"runtime"
	"time"
)

const logFile = "public/logs.json"

type Log struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Body  string `json:"body"`
	CTime int64  `json:"ctime"`
}

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/bbs", showHandler)
	http.HandleFunc("/write", writeHandler)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Codespace !")
}

func showHandler(w http.ResponseWriter, r *http.Request) {
	htmlLog := ""
	logs := loadLogs()

	for _, i := range logs {
		htmlLog += fmt.Sprintf(
			"<p>(%d) <span>%s</span>: %s <button onclick='addReply(%d)'>返信</button> --- %s</p>",
			i.ID,
			html.EscapeString(i.Name),
			html.EscapeString(i.Body),
			i.ID,
			time.Unix(i.CTime, 0).Format("2006/1/2 15:04"))
	}
	htmlBody := "<html><head><style>" +
		"p { border: 1px solid silver; padding: 1em;} " +
		"span { background-color: #eef; } " +
		"</style>" +
		"<script>" +
		"function addReply(id) {" +
		"  var body = document.getElementById('msgBody');" +
		"  body.value = '>>' + id + ' ' + body.value;" +
		"  body.focus();" +
		"}" +
		"</script>" +
		"</head><body><h1>BBS</h1>" +
		getForm() + htmlLog + "</body></html>"
	w.Write([]byte(htmlBody))
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var log Log
	log.Name = r.Form["name"][0]
	log.Body = r.Form["body"][0]
	if log.Name == "" {
		log.Name = "名無し"
	}
	logs := loadLogs()
	log.ID = len(logs) + 1
	log.CTime = time.Now().Unix()
	logs = append(logs, log)
	saveLogs(logs)
	http.Redirect(w, r, "/bbs", 302)
}


func getForm() string {
	return "<div><form action='/write' method='get'>" +
		"名前: <input type='text' name='name'><br>" +
		"本文: <input type='text' name='body' style='width:30em;' id='msgBody'><br>" +
		"<input type='submit' value='書込'>" +
		"</form></div><hr>"
}

func loadLogs() []Log {

	text, err := os.ReadFile(logFile)
	if err != nil {
		return make([]Log, 0)
	}

	var logs []Log
	json.Unmarshal([]byte(text), &logs)
	return logs
}


func saveLogs(logs []Log) {

	bytes, _ := json.Marshal(logs)

	os.WriteFile(logFile, bytes, 0644)
}
