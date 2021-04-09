package main

import (
	"fmt"
	"net/http"
	"example-api-project/pkg/transport"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	r := transport.Router()
	fmt.Println(http.ListenAndServe(":"+port, r))
}
