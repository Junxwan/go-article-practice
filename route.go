package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func webRoute() {
	router.Use(checkLogin())

	router.GET("/", showIndex)
	router.GET("/login", showLogin)
	router.POST("/login", attempt)
	router.GET("/logout", logout)
	router.GET("/register", showRegister)
	router.POST("/register", register)
	router.GET("/article/view/:id", showArticle)
	router.GET("/create", showCreateArticle)
}

// 讀取view
func reade(c *gin.Context, view string, data gin.H) {
	isLogin, _ := c.Get("isLogin")

	data["isLogin"] = isLogin.(bool)

	c.HTML(http.StatusOK, view, data)
}
