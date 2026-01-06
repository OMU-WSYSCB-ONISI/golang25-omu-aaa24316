package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    // "./public" ディレクトリ（index.htmlが入っている場所）を
    // ファイルサーバーとして配信します。
    fs := http.FileServer(http.Dir("./public"))
    http.Handle("/", fs)

    fmt.Println("サーバーを http://localhost:8080 で起動します")

    // サーバーを8080番ポートで起動
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
