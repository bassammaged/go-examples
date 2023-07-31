package httpclient

import (
	"io"
	"net/http"
	"net/url"
)

const URL = "https://dummyjson.com/"

func sendGetReq(paramerizedAbosultePath string) ([]byte, error) {
	// Prepare the URL and validate it
	url := URL + paramerizedAbosultePath
	if err := parseURL(url); err != nil {
		return nil, err
	}

	// Send the request
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// Fetch the response
	resBodyInBytes, err := fetchResponseBody(res)
	if err != nil {
		return nil, err
	}

	return resBodyInBytes, nil
}

func parseURL(urLocator string) error {
	_, err := url.ParseRequestURI(urLocator)
	return err
}

func fetchResponseBody(response *http.Response) ([]byte, error) {
	resBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return resBytes, nil
}
