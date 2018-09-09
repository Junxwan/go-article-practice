package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// 文章結構
type article struct {
	ID        int
	Title     string
	Content   string
	CreatedAt string
}

// 文章列表
var articleList = []article{
	article{ID: 1, Title: "Title 1", Content: "body 1", CreatedAt: now()},
}

// 首頁
func showIndex(c *gin.Context) {
	reade(c, "index.html", gin.H{
		"title":   "go-article-practice",
		"article": articleList,
	})
}

// 顯示出單一文章
func showArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	article, err := getArticleByID(id)

	if (err != nil) {
		reade(c, "result.html", gin.H{
			"message": err.Error(),
		})
	} else {
		reade(c, "article.html", gin.H{
			"title":   article.Title,
			"article": article,
		})
	}
}

// 新增文章頁
func showCreateArticle(c *gin.Context)  {
	reade(c, "create.html", gin.H{
		"title":   "新增文章",
	})
}

// 取文章
func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}

	return nil, errors.New("找不到文章")
}

// 取得現在時間
func now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
