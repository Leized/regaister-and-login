package api

import (
	"Project/app/model"
	"Project/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

func Login(r *ghttp.Request) {
	var req model.Users
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
		encryption := service.Encryption(req.Password)
		rows := service.QueryName(req.Username)
		if rows.ID > 0 {
			if rows.Password == encryption.Password {
				r.Response.WriteJsonP(g.Map{
					"code":    0,
					"message": "登陆成功",
				})
			} else {
				r.Response.WriteJsonP(g.Map{
					"code":    1,
					"message": "密码错误",
				})
			}
		} else {
			r.Response.WriteJsonP(g.Map{
				"code":    1,
				"message": "用户不存在，请注册后再进行登陆",
			})
		}
	}
}
