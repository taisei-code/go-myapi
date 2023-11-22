package main

import (
	"log"
	"net/http"

	"github.com/yourname/reponame/handlers"
)

func main () {
	
	// 定義したハンドラをサーバーで使う登録
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)


	// サーバー起動時のログ出力
	log.Println("server start at port 8080")

	// ListenAndServe関数でサーバー起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}