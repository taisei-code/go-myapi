package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// helloのハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {

	//reqのMethodフィールドがGETなら
	if req.Method == http.MethodGet {
		// 通常通りにレスポンスを返す
		io.WriteString(w, "Hello, world!\n")

	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}

}

// articleのハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Article...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// /article/listのハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodGet {
		io.WriteString(w, "Article List\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// /article/1のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodGet {
		articleID := 1
		resString := fmt.Sprintf("Article No.%d\n", articleID)
		io.WriteString(w, resString)
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// /article/niceのハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {

		if req.Method == http.MethodPost {
			io.WriteString(w, "Posting Nice...\n")
		} else {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
}

// /commentのハンドラ
func  PostCommentHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Comment...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}



	

	

	

