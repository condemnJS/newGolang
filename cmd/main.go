package main

import (
	"fmt"
	"net/http"
	"example-api-project/pkg/transport"
	"example-api-project/pkg/types"
)

func main() {
	r := transport.Router()
	fmt.Println(http.ListenAndServe(types.PORT, r))
}
