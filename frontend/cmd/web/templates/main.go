package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})
	fmt.Println("The Server is runnig on the PORT 8081")
	err := http.ListenAndServe(":8081", nil)

}
