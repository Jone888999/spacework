package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "goods",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "news",
		})
	})
	r.GET("/news", func(c *gin.Context) {
		news := &Article{
			Title:   "新闻标题2",
			Content: "新闻内容",
		}
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title":   "新闻标题",
			"content": news,
		})
	})
	r.Run(":8080")
}
