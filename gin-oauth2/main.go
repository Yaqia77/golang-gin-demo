package main

import (
    "gin-oauth2/router"
)

func main() {
    r := router.SetupRouter()
    r.Run(":8080")
}
