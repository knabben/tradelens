package request

import (
	"fmt"
	"net/http"
)

var (
	baseUrl string = "https://platform-sandbox.tradelens.com"
)
// Get makes a request in the platform
func Get(url string, token string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl + url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Bearer " + token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	return resp
}