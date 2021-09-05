package handler

import (
	"encoding/json"
	"opa-backend/configs"
	"opa-backend/utils"

	"github.com/gin-gonic/gin"
)

func CreateCode(config configs.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := newCreateCodeResest()

		code, err := createCode(&req, config)
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

func createCode(orderCode *createCodeResest, config configs.Config) (*createCodeResponse, error) {
	method := "POST"
	path := "/v2/codes"
	url := config.BASEURL + path
	data, err := json.Marshal(orderCode)
	if err != nil {
		return nil, err
	}

	header, err := utils.GetHeader(method, path, data, config)
	if err != nil {
		return nil, err
	}
	query := map[string]string{
		"assumeMerchant": config.ASSUMEMERCHANT,
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

type createCodeResest struct {
	MerchantPaymentID string `json:"merchantPaymentId"`
	Amount            amount `json:"amount"`
	CodeType          string `json:"codeType"`
}

type createCodeResponse struct {
	ResultInfo resultInfo `json:"resultInfo"`
	Data       data       `json:"data"`
}
