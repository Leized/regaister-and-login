package router

import (
	"Project/app/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		//group.ALL("/hello", api.Hello)
		group.ALL("/register", api.Register)
		group.ALL("/login", api.Login)
	})
}
