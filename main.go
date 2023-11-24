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
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)


	// サーバー起動時のログ出力
	log.Println("server start at port 8080")

	// ListenAndServe関数でサーバー起動
	log.Fatal(http.ListenAndServe(":8080", r))
}