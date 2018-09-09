package main

func webRoute() {
	router.GET("/", showIndex)
	router.GET("/login", showLogin)
	router.GET("/register", showRegister)
	router.POST("/register", register)
}
