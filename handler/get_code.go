package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"opa-backend/configs"
	"opa-backend/utils"

	"github.com/gin-gonic/gin"
)

func GetCode(config configs.ApiConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		merchantPaymentId := c.Param("merchantPaymentId")
		statusCode, reason, err := getCode(merchantPaymentId, config)
		if err != nil {
			c.JSON(http.StatusBadRequest, reason)
		}
		c.JSON(statusCode, reason)

	}
}

func getCode(merchantPaymentId string, config configs.ApiConfig) (int, *getCodeResponse, error) {
	method := "GET"
	path := "/v2/codes/payments/" + merchantPaymentId
	url := config.BASEURL + path

	header, err := utils.GetHeader(method, path, nil, config)
	if err != nil {
		return 0, nil, err
	}

	query := utils.GetQuery(config.ASSUMEMERCHANT)

	statusCode, res, err := utils.DoHttpRequest(method, url, header, query, nil)
	if err != nil {
		return 0, nil, err
	}

	var response getCodeResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		fmt.Println("Unmarshal err")
		fmt.Println(err)
		return 0, nil, err
	}

	fmt.Println("#### respons ####")
	fmt.Println(string(res))
	return statusCode, &response, nil
}

type getCodeResponse struct {
	ResultInfo resultInfo          `json:"resultInfo"`
	Data       getCodeResponseData `json:"data"`
}

type getCodeResponseData struct {
	PaymentID          string      `json:"paymentId"`
	Status             string      `json:"status"`
	AcceptedAt         int         `json:"acceptedAt"`
	Refunds            refunds     `json:"refunds"`
	Captures           captures    `json:"captures"`
	Rrevert            revert      `json:"revert"`
	MerchantPaymentID  string      `json:"merchantPaymentId"`
	Amount             amount      `json:"amount"`
	RequestedAt        int         `json:"requestedAt"`
	ExpiresAt          int         `json:"expiresAt"`
	CanceledAt         int         `json:"canceledAt"`
	StoreId            string      `json:"storeId"`
	TerminalId         string      `json:"terminalId"`
	OrderReceiptNumber string      `json:"orderReceiptNumber"`
	OrderDescription   string      `json:"orderDescription"`
	OrderItems         []orderItem `json:"orderItems"`
	Metadata           interface{} `json:"metadata"`
}
