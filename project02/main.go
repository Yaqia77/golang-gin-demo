package main

import (
	"project02/database"
	"project02/router"
)

func main() {
	database.InitDB()

	r := router.Router()
	r.Run(":8080")

}
