package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
)

type user struct {
	Username string `form:"username" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

var account = []user{}

func showLogin(c *gin.Context) {
	reade(c, "login.html", gin.H{
		"title": "Login",
	})
}

func showRegister(c *gin.Context) {
	reade(c, "register.html", gin.H{
		"title": "Register",
	})
}

func login(c *gin.Context) {
	username, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")

	if (isUser(username, password)) {
		c.SetCookie("login", strconv.FormatInt(rand.Int63(), 20), 3600, "", "", false, true)

		c.Set("isLogin", true)

		reade(c, "result.html", gin.H{
			"message": "登入成功",
		})
	} else {
		reade(c, "result.html", gin.H{
			"message": "登入失敗",
		})
	}
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

	reade(c, "result.html", gin.H{
		"message": message,
	})
}

func isUser(username, password string) bool {
	for _, u := range account {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func checkLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("login"); err == nil || token != "" {
			c.Set("isLogin", true)
		} else {
			c.Set("isLogin", false)
		}
	}
}