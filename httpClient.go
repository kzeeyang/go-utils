package utils

import (
	"io/ioutil"
	"net/http"
)

// ClientRequest - http client
func ClientRequest(reqOption *http.Request) (body []byte, err error) {
	client := &http.Client{}

	resp, err := client.Do(reqOption)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ = ioutil.ReadAll(resp.Body)

	return
}
