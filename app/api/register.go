package api

import (
	"Project/app/model"
	"Project/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

// Register 注册
func Register(r *ghttp.Request) {
	var req model.RegisterReq
	if err := r.Parse(&req); err != nil {
		// Validation error.
		if v, ok := err.(gvalid.Error); ok {
			r.Response.WriteJsonExit(model.UserRes{
				Code:    1,
				Message: v.FirstString(),
			})
		}
		// Other error.
		r.Response.WriteJsonExit(model.UserRes{
			Code:    1,
			Message: err.Error(),
		})
	} else {
		user := service.QueryName(req.Username)
		if user.ID > 0 {
			r.Response.WriteJsonP(g.Map{
				"code":    1,
				"message": "用户已存在，请进行登陆",
			})
		} else {
			encryption := service.Encryption1(req.Password)
			service.InsertData(req.Username, encryption.Password)
			r.Response.WriteJsonP(g.Map{
				"code":    0,
				"message": "注册成功",
			})
		}
	}
}
