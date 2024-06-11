package controllers

import (
	"gin-graphgl/graph"
	"gin-graphgl/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

// GraphqlHandler是GraphQL的处理函数，它将GraphQL的请求转发给GraphQL的服务端，并返回响应。
func GraphqlHandler() gin.HandlerFunc {
	// 创建GraphQL服务端

	// 这里的generated.NewExecutableSchema函数会根据schema.graphqls文件生成GraphQL的schema，并将其绑定到Resolver
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		// 处理GraphQL请求
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler是GraphQL的Playground处理函数，它将Playground的请求转发给GraphQL的Playground，并返回响应。
func PlaygroundHandler() gin.HandlerFunc {
	// 创建GraphQL的Playground
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
