package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func DoHttpRequest(method, url string, header, query map[string]string, data []byte) ([]byte, error) {
	if method != "GET" && method != "POST" {
		return nil, errors.New("method's neither GET or POST")
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	for key, value := range header {
		req.Header.Add(key, value)
	}

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	fmt.Println("req: ", req)
	fmt.Println("res: ", res)
	fmt.Println("err: ", err)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
