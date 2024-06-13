package main

import (
	"gin-blogs/config"
	"gin-blogs/router"
)

func main() {
	config.InitDB()

	var router = router.SetupRouter()

	router.Run()
}
