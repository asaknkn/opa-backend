package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"opa-backend/configs"
	"opa-backend/utils"

	"github.com/gin-gonic/gin"
)

func CreateCode(config configs.ApiConfig) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req createCodeRequest
		err := c.ShouldBindJSON(&req)
		if err != nil {
			fmt.Println("#### erro bind json ####")
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		statusCode, code, err := createCode(&req, config)
		if err != nil {
			c.JSON(http.StatusBadRequest, code)
		}
		c.JSON(statusCode, code)
	}
}

func createCode(orderCode *createCodeRequest, config configs.ApiConfig) (int, *createCodeResponse, error) {
	method := "POST"
	path := "/v2/codes"
	url := config.BASEURL + path
	data, err := json.Marshal(orderCode)
	if err != nil {
		return 0, nil, err
	}

	header, err := utils.GetHeader(method, path, data, config)
	if err != nil {
		return 0, nil, err
	}

	query := utils.GetQuery(config.ASSUMEMERCHANT)

	statusCode, res, err := utils.DoHttpRequest(method, url, header, query, data)
	if err != nil {
		return 0, nil, err
	}

	var code createCodeResponse
	err = json.Unmarshal(res, &code)
	if err != nil {
		return 0, nil, err
	}

	return statusCode, &code, nil
}

type createCodeRequest struct {
	MerchantPaymentID   string             `json:"merchantPaymentId" validate:"required"`
	Amount              amount             `json:"amount" validate:"required"`
	OrderDescription    string             `json:"orderDescription,omitempty"`
	OrderItems          []orderRequestItem `json:"orderItems,omitempty"`
	CodeType            string             `json:"codeType" validate:"required"`
	StoreInfo           string             `json:"storeInfo,omitempty"`
	StoreId             string             `json:"storeId,omitempty"`
	TerminalId          string             `json:"terminalId,omitempty"`
	RequestedAt         int                `json:"requestedAt,omitempty"`
	RedirectUrl         string             `json:"redirectUrl,omitempty"`
	RedirectType        string             `json:"redirectType,omitempty"`
	UserAgent           string             `json:"userAgent,omitempty"`
	IsAuthorization     bool               `json:"isAuthorization,omitempty"`
	AuthorizationExpiry int                `json:"authorizationExpiry,omitempty"`
}

type createCodeResponse struct {
	ResultInfo resultInfo             `json:"resultInfo"`
	Data       createCodeResponseData `json:"data"`
}

type createCodeResponseData struct {
	CodeID              string      `json:"codeId"`
	Url                 string      `json:"url"`
	Deeplink            string      `json:"deeplink"`
	ExpiryDate          int         `json:"expiryDate"`
	MerchantPaymentId   string      `json:"merchantPaymentId"`
	Amount              amount      `json:"amount"`
	OrderDescription    string      `json:"orderDescription"`
	OrderItems          []orderItem `json:"orderItems"`
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
