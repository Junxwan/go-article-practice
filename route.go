package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func webRoute() {
	router.Use(checkLogin())

	router.GET("/", showIndex)
	router.GET("/login", showLogin)
	router.POST("/login", login)
	router.GET("/register", showRegister)
	router.POST("/register", register)
}

// 讀取view
func reade(c *gin.Context, view string, data gin.H) {
	isLogin, _ := c.Get("isLogin")

	data["isLogin"] = isLogin.(bool)

	c.HTML(http.StatusOK, view, data)
}
