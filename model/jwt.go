package model

// JWT基本信息
type JwtInfo struct {
	Exp      int64  `json:"exp"`      //过期时间
	Iat      int64  `json:"iat"`      //颁发时间
	Token    string `json:"token"`    //token
	UserId   int    `json:"userId"`   //用户编号
	UserName string `json:"userName"` //用户姓名
	IP       string `json:"ip"`       //IP地址
}
