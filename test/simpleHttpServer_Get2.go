package main

//更简单的http-get服务器，本质上都是调用了http.Handle.
import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main () {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8000", nil)
}