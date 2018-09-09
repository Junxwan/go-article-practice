package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	Username string `form:"username" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

var account = []user{}

func showLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

func showRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register",
	})
}

func register(c *gin.Context) {
	var form user
	message := ""

	username, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")

	if err := c.ShouldBind(&form); err == nil {
		account = append(account, user{
			Username: username,
			Password: password,
		})

		message = "恭喜你註冊成功，請前往登入頁做登入"
	} else {
		message = "帳號密碼輸入有誤請重新填寫"
	}

	c.HTML(http.StatusOK, "result.html", gin.H{
		"message": message,
	})
}
