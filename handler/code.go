package handler

import (
	"encoding/json"
	"opa-backend/utils"

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
		Amount: amount{
			Amount:   100,
			Currency: "JPY",
		},
		CodeType: "ORDER_QR",
	}
}

func createCode(orderCode *createCodeResest, apiKey, apiSecret string) (*createCodeResponse, error) {
	method := "POST"
	path := "/v2/codes"
	url := "https://stg-api.sandbox.paypay.ne.jp" + path
	data, err := json.Marshal(orderCode)
	if err != nil {
		return nil, err
	}

	header := utils.GetHeader(method, path, apiKey, apiSecret, data)
	query := map[string]string{
		"assumeMerchant": "407910645125136384",
	}
	res, err := utils.DoHttpRequest(method, url, header, query, data)
	if err != nil {
		return nil, err
	}

	var code createCodeResponse
	err = json.Unmarshal(res, &code)
	if err != nil {
		return nil, err
	}

	return &code, nil
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
	Amount            amount `json:"amount"`
	CodeType          string `json:"codeType"`
}

type createCodeResponse struct {
	ResultInfo resultInfo `json:"resultInfo"`
	Data       data       `json:"data"`
}
