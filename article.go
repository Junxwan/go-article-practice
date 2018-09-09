package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type article struct {
	ID        int
	Title     string
	Content   string
	CreatedAt string
}

var articleList = []article{
	article{ID: 1, Title: "Title 1", Content: "body 1", CreatedAt: now()},
}

func showIndex(c *gin.Context) {
	reade(c, "index.html", gin.H{
		"title":   "go-article-practice",
		"article": articleList,
	})
}

func now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
