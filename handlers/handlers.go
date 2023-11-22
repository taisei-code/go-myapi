package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// helloのハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
		// ハンドラの処理内容（何がきても、Hello, Worldを返す
		io.WriteString(w, "Hello, world!\n")
}

// articleのハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Article...\n")
}

// /article/listのハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article List\n")
}

// /article/1のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
		articleID := 1
		resString := fmt.Sprintf("Article No.%d\n", articleID)
		io.WriteString(w, resString)
}

// /article/niceのハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Nice...\n")
}

// /commentのハンドラ
func  PostCommentHandler(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Comment...\n")
}



	

	

	

