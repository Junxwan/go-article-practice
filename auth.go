package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

// 使用者結構
type user struct {
	Username string `form:"username" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

// 現已註冊的使用者
var account = []user{}

// 登入頁
func showLogin(c *gin.Context) {
	reade(c, "login.html", gin.H{
		"title": "Login",
	})
}

// 註冊頁
func showRegister(c *gin.Context) {
	reade(c, "register.html", gin.H{
		"title": "Register",
	})
}

// 嘗試登入
func attempt(c *gin.Context) {
	username, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")

	if (isUser(username, password)) {
		login(c)
	} else {
		reade(c, "result.html", gin.H{
			"message": "登入失敗",
		})
	}
}

// 登入
func login(c *gin.Context) {
	c.SetCookie("login", strconv.FormatInt(rand.Int63(), 20), 3600, "", "", false, true)

	c.Set("isLogin", true)

	reade(c, "result.html", gin.H{
		"message": "登入成功",
	})
}

// 註冊
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

// 登出
func logout(c *gin.Context) {
	c.SetCookie("login", "", -1, "", "", false, true)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

// 檢查帳號是否正確
func isUser(username, password string) bool {
	for _, u := range account {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

// 檢查是否已登入
func checkLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("login"); err == nil && token != "" {
			c.Set("isLogin", true)
		} else {
			c.Set("isLogin", false)
		}
	}
}
