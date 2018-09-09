package main

func webRoute() {
	router.GET("/", showIndex)
	router.GET("/login", showLogin)
}
