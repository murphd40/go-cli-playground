package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

func main() {
	port := pflag.Int16P("port", "p", 8080, "the port to use")
	pflag.Parse()

	router := gin.Default()
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any {
			"status": "OK",
		})
	})
	router.POST("/echo", func(ctx *gin.Context) {
		var body map[string]any
		if err := ctx.BindJSON(&body); err != nil {
			log.Println(err.Error())
			return
		}

		bs, _ := json.MarshalIndent(body, "", " ")
		log.Println(string(bs))

		ctx.JSON(200, body)
	})

	server := http.Server{
		Handler: router,
		Addr: fmt.Sprint(":", *port),
	}

	server.ListenAndServe()
}
