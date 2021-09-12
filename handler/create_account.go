package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"opa-backend/configs"
	"opa-backend/utils"

	"github.com/gin-gonic/gin"
)

func CreateAccountLinkQRCode(config configs.ApiConfig) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req createAccountLinkQRCodeRequest
		err := c.ShouldBindJSON(&req)
		if err != nil {
			fmt.Println("#### erro bind json ####")
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		statusCode, code, err := creteAccountLinkQRCode(&req, config)
		if err != nil {
			c.JSON(http.StatusBadRequest, code)
		}
		c.JSON(statusCode, code)
	}
}

func creteAccountLinkQRCode(req *createAccountLinkQRCodeRequest, config configs.ApiConfig) (int, *createCodeResponse, error) {
	method := "POST"
	path := "/v1/qr/sessions"
	url := config.BASEURL + path
	data, err := json.Marshal(req)
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

type createAccountLinkQRCodeRequest struct {
	Scopes       string `json:"scopes" validate:"required"`
	Nonce        string `json:"nonce" validate:"required"`
	RedirectType string `json:"redirectType,omitempty"`
	RedirectUrl  string `json:"redirectUrl" validate:"required"`
	ReferenceID  string `json:"referenceId" validate:"required"`
	PhoneNumber  string `json:"phoneNumber,omitempty"`
	DeviceId     string `json:"deviceId,omitempty"`
	UserAgent    string `json:"userAgent,omitempty"`
}

type createAccountLinkQRCodeResponse struct {
	ResultInfo resultInfo                          `json:"resultInfo"`
	Data       createAccountLinkQRCodeResponseData `json:"data"`
}

type createAccountLinkQRCodeResponseData struct {
	LinkQRCodeURL string `json:"linkQRCodeURL"`
}
