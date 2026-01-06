package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/webfortune", fortunehandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fortunehandler(w http.ResponseWriter, r *http.Request) {
	seed := time.Now().UnixNano()
	d := rand.New(rand.NewSource(seed))

	diceValue := d.Int31n(6) + 1

	var fortune string
	if diceValue == 1 {
		fortune = "凶"
	} else if diceValue == 2 || diceValue == 3 {
		fortune = "吉"
	} else if diceValue == 4 || diceValue == 5 {
		fortune = "中吉"
	} else {
		fortune = "大吉"
	}

	fmt.Fprintf(w, "今の運勢は%sです", fortune)
}
