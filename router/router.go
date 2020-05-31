package router

import "net/http"

func Handle() *http.ServeMux {
	handler := http.ServeMux{}

	handler.HandleFunc("/", tweets)
	handler.HandleFunc("/tweet", tweet)
	handler.HandleFunc("/followed", followed)

	return &handler
}
