package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/enq", enqhandler)
	http.HandleFunc("/fdump", fdump)
	http.HandleFunc("/cal00", cal00handler)
	http.HandleFunc("/cal01", calpmhandler)
	http.HandleFunc("/sum", sumhandler)
	http.HandleFunc("/bmi", bmihandler)

	http.HandleFunc("/kaiseki", kaisekihandler)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Codespace !")
}

func fdump(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	for k, v := range r.Form {
		fmt.Printf("%v : %v\n", k, v)
	}
}

func enqhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	fmt.Fprintln(w, r.FormValue("name")+"さん，ご協力ありがとうございます.\n年齢は"+r.FormValue("age")+"で，性別は"+r.FormValue("gend")+"で，出身地は"+r.FormValue("birthplace")+"ですね")
}

func cal00handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	price, _ := strconv.Atoi(r.FormValue("price"))
	num, _ := strconv.Atoi(r.FormValue("num"))
	fmt.Fprint(w, "合計金額は ")
	fmt.Fprintln(w, price*num)
}

func calpmhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	x, _ := strconv.Atoi(r.FormValue("x"))
	y, _ := strconv.Atoi(r.FormValue("y"))
	switch r.FormValue("cal0") {
	case "+":
		fmt.Fprintln(w, x+y)
	case "-":
		fmt.Fprintln(w, x-y)
	case "*":
		fmt.Fprintln(w, x*y)
	case "/":
		if y != 0 {
			fmt.Fprintln(w, x/y)
		} else {
			fmt.Fprintln(w, "エラー: 0で割ることはできません")
		}
	}
}

func sumhandler(w http.ResponseWriter, r *http.Request) {
	var sum, tt int
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	tokuten := strings.Split(r.FormValue("dd"), ",")
	fmt.Println(tokuten)
	for i := range tokuten {
		val := strings.TrimSpace(tokuten[i])
		tt, _ = strconv.Atoi(val)
		sum += tt
	}
	fmt.Fprintln(w, sum)
	fmt.Println(sum)
}

func bmihandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	weight, _ := strconv.ParseFloat(r.FormValue("w"), 64)
	height, _ := strconv.ParseFloat(r.FormValue("h"), 64)

	heightMeter := height / 100.0
	var bmi float64
	if heightMeter > 0 {
		bmi = weight / (heightMeter * heightMeter)
	}
	fmt.Fprintf(w, "身長: %.1fcm, 体重: %.1fkg のBMIは...\n", height, weight)
	fmt.Fprintf(w, "%.2f です。", bmi)
}

func kaisekihandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}

	input := r.FormValue("dd")
	tokutenList := strings.Split(input, ",")

	var sum int
	var count int
	dist := make([]int, 11)

	for _, s := range tokutenList {
		valStr := strings.TrimSpace(s)
		if valStr == "" { continue }

		score, err := strconv.Atoi(valStr)
		if err != nil {
			continue
		}

		sum += score
		count++

		idx := score / 10
		if idx >= 0 && idx <= 10 {
			dist[idx]++
		}
	}

	fmt.Fprintln(w, "【解析結果】")
	if count > 0 {
		avg := float64(sum) / float64(count)
		fmt.Fprintf(w, "平均点: %.2f 点\n", avg)
	} else {
		fmt.Fprintln(w, "有効なデータがありません")
	}

	fmt.Fprintln(w, "\n--- 得点分布 ---")
	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "%d - %d 点 : %d 人\n", i*10, i*10+9, dist[i])
	}
	fmt.Fprintf(w, "100 点 : %d 人\n", dist[10])
}
