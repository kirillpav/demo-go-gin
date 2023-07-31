package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main(){


// creating router
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", ShowIndexPage)
	router.GET("/article/view/:article_id", getArticle)

	router.Run()
}