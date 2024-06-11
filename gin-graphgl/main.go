package main

import (
	"gin-graphgl/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}

//Gin 和 GraphQL 的结合可以让你创建强大且灵活的 API
//graphgl是go语言的GraphQL实现库，它可以帮助开发者快速构建GraphQL API。
//gin-graphgl是一个基于gin框架的GraphQL API的实现，它提供了GraphQL API的基本功能，包括查询、Mutation、订阅等。

//graphgl应用于什么场景
//graphgl可以应用于任何需要构建GraphQL API的场景，如：
