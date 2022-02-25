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
			err = r.Response.WriteJsonExit(model.UserRes{
				Code:    1,
				Message: v.FirstString(),
			})
			if err != nil {
				return
			}
		}
		// Other error.
		err = r.Response.WriteJsonExit(model.UserRes{
			Code:    1,
			Message: err.Error(),
		})
		if err != nil {
			return
		}
	} else {
		user := service.QueryName(req.Username)
		if user.ID > 0 {
			err = r.Response.WriteJsonP(g.Map{
				"code":    1,
				"message": "用户已存在，请进行登陆",
			})
			if err != nil {
				return
			}
		} else {
			encryption := service.Encryption1(req.Password)
			service.InsertData(req.Username, encryption.Password)
			err = r.Response.WriteJsonP(g.Map{
				"code":    0,
				"message": "注册成功",
			})
			if err != nil {
				return
			}
		}
	}
}
