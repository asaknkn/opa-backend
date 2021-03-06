package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"opa-backend/configs"
	"strconv"
	"strings"
	"time"
)

func DoHttpRequest(method, url string, header, query map[string]string, data []byte) (int, []byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return 0, nil, err
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
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}
	fmt.Println("#### responsBody ####")
	fmt.Println(string(body))
	return res.StatusCode, body, nil
}

func GetHeader(method, path string, body []byte, config configs.ApiConfig) (map[string]string, error) {
	// md5
	contentType := "empty"
	hash := "empty"
	if body != nil {
		fmt.Println(" ####reqestBody ####")
		fmt.Println(string(body))

		contentType = "application/json;charset=UTF-8;"

		md := md5.New()
		defer md.Reset()
		md.Write([]byte(contentType))
		md.Write(body)

		hash = base64.StdEncoding.EncodeToString(md.Sum(nil))
	}
	delimiter := "\n"
	nonce, err := newNonce(8)
	if err != nil {
		return nil, err
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	s := []string{path, method, nonce, timestamp, contentType, hash}
	hmacData := []byte(strings.Join(s, delimiter))
	fmt.Println("#### hmacData ####")
	fmt.Println(string(hmacData))

	apiKey := config.ApiKey
	apiSecret := config.ApiSecret
	mac := hmac.New(sha256.New, []byte(apiSecret))
	mac.Write(hmacData)
	macData := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	authHeader := "hmac OPA-Auth:" + apiKey + ":" + macData + ":" + nonce + ":" + timestamp + ":" + hash
	fmt.Println("#### authHeader ####")
	fmt.Println("hmac OPA-Auth: + apiKey + : + macData + : + nonce + : + timestamp + : + hash")
	fmt.Println(authHeader)

	return map[string]string{
		"Authorization": authHeader,
		"Content-Type":  contentType,
	}, nil
}

func newNonce(d int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// ???????????????
	b := make([]byte, d)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	// letters ??????????????????????????????????????????????????????
	var result string
	for _, v := range b {
		// index ??? letters ????????????????????????????????????
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}

func GetQuery(assumeMerchant string) map[string]string {
	var query map[string]string

	if assumeMerchant != "" {
		query = map[string]string{
			"assumeMerchant": assumeMerchant,
		}
	} else {
		query = nil
	}

	return query
}
