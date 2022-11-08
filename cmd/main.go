package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Capstain/coinsmarketcup_fetcher/pkg/currency"
	"github.com/Capstain/coinsmarketcup_fetcher/pkg/fetcher"
	"github.com/Capstain/coinsmarketcup_fetcher/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	API_KEY := os.Getenv("API_KEY")
	if "" == API_KEY {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Fail on init ENVs", err)
		}
	}

	r := gin.Default()

	r.GET("/api/v1/tokens", func(ctx *gin.Context) {
		ctx.JSONP(http.StatusOK, token.Fake())
	})

	r.GET("/api/v1/tokens/:code", func(ctx *gin.Context) {
		code := ctx.Param("code")
		token := token.Find(code)
		if token == nil {
			ctx.AbortWithStatus(http.StatusNotFound)
		}

		ctx.JSON(http.StatusOK, token)
	})

	r.GET("/api/v1/currencies", func(ctx *gin.Context) {
		ctx.JSONP(http.StatusOK, currency.Fake())
	})

	r.GET("/api/v1/currencies/:code", func(ctx *gin.Context) {
		code := ctx.Param("code")
		currency := currency.Find(code)
		if currency == nil {
			ctx.AbortWithStatus(http.StatusNotFound)
		}

		ctx.JSON(http.StatusOK, currency)
	})

	r.GET("/api/v1/rates/:token/:currency", func(ctx *gin.Context) {
		tokenCode := ctx.Param("token")
		currencyCode := ctx.Param("currency")
		token := token.Find(tokenCode)
		currency := currency.Find(currencyCode)

		rate := fetcher.FetchRate(token, currency)

		ctx.JSON(http.StatusOK, gin.H{
			"rate": rate,
		})

	})

	r.Run()
}
