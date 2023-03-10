package main

import (
	"fmt"
	"net/http"

	"github.com/atsutama2/go_github_actions_test/service"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// HelloServiceをインスタンス化し、ハンドラーに渡します。
	helloService := &service.HelloService{}
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, helloService.GetHello())
	})

	http.ListenAndServe(":8080", r)
}
