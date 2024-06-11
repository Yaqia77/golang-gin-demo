package graph

import (
	"context"
	"gin-graphgl/graph/generated"
)

type Resolver struct{}
type queryResolver struct{ *Resolver }

//解释一下
//queryResolver 是一个结构体，它包含一个指向 Resolver 结构体的指针，这样就可以访问到 Resolver 中的方法了。
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
//Hello 方法是 QueryResolver 接口的一个方法，它返回一个字符串和一个错误。
func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return "Hello, world!", nil
}
