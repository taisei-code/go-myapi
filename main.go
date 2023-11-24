package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourname/reponame/handlers"
)

func main () {
	// ルータrを明示的に宣言
	r := mux.NewRouter()

	// 定義したハンドラをサーバーで使う登録
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler)
	r.HandleFunc("/article/list", handlers.ArticleListHandler)
	r.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler)
	r.HandleFunc("/comment", handlers.PostCommentHandler)


	// サーバー起動時のログ出力
	log.Println("server start at port 8080")

	// ListenAndServe関数でサーバー起動
	log.Fatal(http.ListenAndServe(":8080", r))
}