package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath" // path/filepath をインポート
)

func main() {
	// 1. 実行中の場所（カレントディレクトリ）を取得
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("作業ディレクトリの取得に失敗: %v", err)
	}

    // 2. ★ターミナルに実行場所を表示
	log.Printf("現在の作業ディレクトリ: %s", wd)

    // 3. index.html への絶対パスを計算
	fullPath := filepath.Join(wd, "public", "index.html")

    // 4. ★ターミナルに、探しに行くファイルパスを表示
	log.Printf("探しに行くファイル: %s", fullPath)

    // 5. ★もしファイルが見つからなければ、サーバー起動前にエラーで停止
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		log.Fatalf("！！！ファイルが見つかりません！！！: %s", fullPath)
	}

	// 6. ファイルが見つかった場合のみ、サーバーを起動
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, fullPath)
	})

	fmt.Println("サーバーを http://localhost:8080 で起動します")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
