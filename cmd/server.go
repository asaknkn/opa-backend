package main

import (
	"log"
	"opa-backend/configs"
	"opa-backend/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := configs.NewConfig()
	if err != nil {
		log.Fatal("Can't Set Config with Environment: %v", err)
	}

	r := gin.Default()
	r.POST("/v2/codes", handler.CreateCode(config))
	r.GET("/v2/codes/payments/:merchantPaymentId", handler.GetCode(config))
	r.POST("/v1/qr/sessions", handler.CreateAccountLinkQRCode(config))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
