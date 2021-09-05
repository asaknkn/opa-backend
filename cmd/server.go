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
	r.GET("/ping", handler.CreateCode(config))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
