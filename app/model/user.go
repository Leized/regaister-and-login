package model

type Users struct {
	ID       int
	Username string `p:"username"  v:"required|length:6,30#请输入账号|账号长度为:min到:max位"`
	Password string `p:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
}

type RegisterReq struct {
	Username  string `p:"username"  v:"required|length:6,30#请输入账号|账号长度为:min到:max位"`
	Password  string `p:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
	Password2 string `p:"password2" v:"required|length:6,30|same:password#请确认密码|密码长度不够|两次密码不一致"`
}

type UserRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
