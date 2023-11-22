package main

import (
	"io"
	"log"
	"net/http"
)

func main () {
	// ハンドラの宣言
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		// ハンドラの処理内容（何がきても、Hello, Worldを返す
		io.WriteString(w, "Hello, world!\n")
	}

	// 定義したハンドラをサーバーで使う登録
	// http.HandleFunc("/", helloHandler)
	http.HandleFunc("/hello", helloHandler)

	// サーバー起動時のログ出力
	log.Println("server start at port 8080")

	// ListenAndServe関数でサーバー起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}