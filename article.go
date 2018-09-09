package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
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

type articleList struct {
	Article []article
}

var list articleList

// 初始文章
func init() {
	list.Article = []article{
		article{ID: 1, Title: "Title 1", Content: "body 1", CreatedAt: now()},
	}
}

// 首頁
func showIndex(c *gin.Context) {
	reade(c, "index.html", gin.H{
		"title":   "go-article-practice",
		"article": list.Article,
	})
}

// 顯示出單一文章
func showArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	article, err := list.getArticle(id)

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
func showCreateArticle(c *gin.Context) {
	reade(c, "create.html", gin.H{
		"title": "新增文章",
	})
}

// 新增文章頁
func createArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	list.addArticle(title, content)

	c.Redirect(http.StatusFound, "/")
}

// 取文章
func (a articleList) getArticle(id int) (*article, error) {
	for _, a := range a.Article {
		if a.ID == id {
			return &a, nil
		}
	}

	return nil, errors.New("找不到文章")
}

// 新增文章
func (a articleList) addArticle(title, content string) {
	list.Article = append(a.Article, article{
		ID:        len(a.Article) + 1,
		Title:     title,
		Content:   content,
		CreatedAt: now(),
	})
}

// 取得現在時間
func now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
