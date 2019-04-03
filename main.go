package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from me :)")
	})
	log.Fatal(http.ListenAndServeTLS(":8082", "temporary.atatiki.com.crt", "temporary.atatiki.com.key", nil))
}
