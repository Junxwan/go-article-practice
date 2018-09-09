package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func showLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}
