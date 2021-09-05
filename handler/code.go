package handler

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"opa-backend/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

func CreateCode() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := newCreateCodeResest()

		var goenv Env
		envconfig.Process("OPA", &goenv)
		apiKey := goenv.ApiKey
		apiSecret := goenv.ApiSecret

		code, err := createCode(&req, apiKey, apiSecret)
		if err != nil {
			c.JSON(400, code)
		}

		c.JSON(200, code)
	}
}

func newCreateCodeResest() createCodeResest {
	return createCodeResest{
		MerchantPaymentID: "1234678",
		Amount: Amount{
			Amount:   100,
			Currency: "JPY",
		},
		CodeType: "ORDER_QR",
	}
}

func createCode(orderCode *createCodeResest, apiKey, apiSecret string) (*CreateCodeResponse, error) {
	method := "POST"
	path := "/v2/codes"
	url := "https://stg-api.sandbox.paypay.ne.jp" + path
	data, err := json.Marshal(orderCode)
	if err != nil {
		return nil, err
	}

	header := getHeader(method, path, apiKey, apiSecret, data)
	query := map[string]string{
		"assumeMerchant": "407910645125136384",
	}
	res, err := utils.DoHttpRequest(method, url, header, query, data)
	if err != nil {
		return nil, err
	}

	var code CreateCodeResponse
	err = json.Unmarshal(res, &code)
	if err != nil {
		return nil, err
	}

	return &code, nil
}

func getHeader(method, path, apiKey, apiSecret string, body []byte) map[string]string {

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

// type createCodeResest struct {
// 	MerchantPaymentID   string      `json:"merchantPaymentId" validate:"required"`
// 	Amount              Amount      `json:"amount" validate:"required"`
// 	OrderDescription    string      `json:"orderDescription"`
// 	OrderItems          []OrderItem `json:"orderItems"`
// 	CodeType            string      `json:"codeType" validate:"required"`
// 	StoreInfo           string      `json:"storeInfo"`
// 	StoreId             string      `json:"storeId"`
// 	TerminalId          string      `json:"terminalId"`
// 	RequestedAt         int         `json:"requestedAt"`
// 	RedirectUrl         string      `json:"redirectUrl"`
// 	RedirectType        string      `json:"redirectType"`
// 	UserAgent           string      `json:"userAgent"`
// 	IsAuthorization     bool        `json:"isAuthorization"`
// 	AuthorizationExpiry int         `json:"authorizationExpiry"`
// }

type Env struct {
	ApiKey    string `envconfig:"APIKEY" split_words:"true"`
	ApiSecret string `envconfig:"APISECRET" split_words:"true"`
}

type createCodeResest struct {
	MerchantPaymentID string `json:"merchantPaymentId"`
	Amount            Amount `json:"amount"`
	CodeType          string `json:"codeType"`
}

type CreateCodeResponse struct {
	ResultInfo ResultInfo `json:"resultInfo"`
	Data       Data       `json:"data"`
}

type ResultInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	CodeID  string `json:"codeId"`
}

type Data struct {
	CodeID              string      `json:"codeId"`
	Url                 string      `json:"url"`
	Deeplink            string      `json:"deeplink"`
	ExpiryDate          int         `json:"expiryDate"`
	MerchantPaymentId   string      `json:"merchantPaymentId"`
	Amount              Amount      `json:"amount"`
	OrderDescription    string      `json:"orderDescription"`
	OrderItems          []OrderItem `json:"orderItems"`
	CodeType            string      `json:"codeType"`
	StoreInfo           string      `json:"storeInfo"`
	StoreID             string      `json:"storeId"`
	TerminalID          string      `json:"terminalId"`
	RequestedAt         int         `json:"requestedAt"`
	RedirectUrl         string      `json:"redirectUrl"`
	RedirectType        string      `json:"redirectType"`
	IsAuthorization     bool        `json:"isAuthorization"`
	AuthorizationExpiry int         `json:"authorizationExpiry"`
}

type Amount struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

type OrderItem struct {
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Quantity  string    `json:"quantity"`
	ProductID string    `json:"productId"`
	UnitPrice UnitPrice `json:"unit_price"`
}

type UnitPrice struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}
