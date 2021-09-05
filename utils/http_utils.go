package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"opa-backend/configs"
	"strconv"
	"strings"
	"time"
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

func GetHeader(method, path string, body []byte, config configs.Config) map[string]string {

	// md5
	contentType := "empty"
	hash := "empty"
	if method != "GET" {
		fmt.Println(" ####reqestBody ####")
		fmt.Println(string(body))
		//contentType = "application/json"
		contentType = "application/json;charset=UTF-8;"

		md := md5.New()
		defer md.Reset()
		md.Write([]byte(contentType))

		//fmt.Println(" #### data ####")
		//fmt.Println(`{"sampleRequestBodyKey1":"sampleRequestBodyValue1","sampleRequestBodyKey2":"sampleRequestBodyValue2"}`)
		//md.Write([]byte(`{"sampleRequestBodyKey1":"sampleRequestBodyValue1","sampleRequestBodyKey2":"sampleRequestBodyValue2"}`))

		// type test struct {
		// 	SampleRequestBodyKey1 string
		// 	SampleRequestBodyKey2 string
		// }

		// testData := test{
		// 	SampleRequestBodyKey1: "sampleRequestBodyValue11",
		// 	SampleRequestBodyKey2: "sampleRequestBodyValue22",
		// }

		// testBody, _ := json.Marshal(testData)

		// fmt.Println(" #### data ####")
		// fmt.Println(body)
		// fmt.Println(string(testBody))
		// fmt.Println(`{"SampleRequestBodyKey1":"sampleRequestBodyValue11","SampleRequestBodyKey2":"sampleRequestBodyValue22"}`)
		// md.Write([]byte(`{"SampleRequestBodyKey1":"sampleRequestBodyValue11","SampleRequestBodyKey2":"sampleRequestBodyValue22"}`))
		// md.Write(testBody)
		md.Write(body)

		hash = base64.StdEncoding.EncodeToString(md.Sum(nil))
		fmt.Println(" #### hash ####")
		fmt.Println(hash)
	}
	delimiter := "\n"
	nonce := "181dfc"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	//timestamp := "1629853169"
	fmt.Println(" #### timestamp ####")
	fmt.Println(timestamp)

	/*
		hmacDataa := []byte(path + delimiter +
			method + delimiter +
			nonce + delimiter +
			timestamp + delimiter +
			contentType + delimiter +
			hash + delimiter)
	*/
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
	fmt.Println(authHeader)
	fmt.Println()

	return map[string]string{
		"Authorization": authHeader,
		"Content-Type":  contentType,
		//"X-ASSUME-MERCHANT": "407910645125136384",
	}
}
