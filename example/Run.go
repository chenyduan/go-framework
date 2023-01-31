package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("hello go!")
	http.HandleFunc("/", root)
}
func root(w http.ResponseWriter, r *http.Request) {

}
