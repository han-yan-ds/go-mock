package main

import (
	"fmt"
	"net/http"

	mocks "github.com/han-yan-ds/go-mocks/mocks"
)

func main() {
	mocks.Requester = &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
	res, _ := mocks.Requester.Do(req)
	fmt.Println(res.StatusCode)
}
