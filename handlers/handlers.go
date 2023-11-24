package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// helloのハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
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

// /article/{id}のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
		articleID, err := strconv.Atoi(mux.Vars(req)["id"])
		if err != nil {
			// 404エラー（BadRequest)を返す
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
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



	

	

	

