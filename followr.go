package main

import (
	"fmt"
	"net/http"
)

func main() {

	init_db()
	handle_routes()

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))

}
