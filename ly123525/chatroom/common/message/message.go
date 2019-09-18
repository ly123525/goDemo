package message

const (
	// LoginMesType 消息类型，登录
	LoginMesType = "LoginMes"
	// LoginResMesType 消息类型，登录返回
	LoginResMesType = "LoginResMes"
)

// Message 传递的消息
type Message struct {
	Type string `json:"type"` // 消息的类型
	Data string `json:"date"` // 消息的内容
}

// LoginMes 发送的时候的date
type LoginMes struct {
	UserID   int    `json:"userId"`   // 用户id
	UserPwd  string `json:"userPwd"`  // 用户密码
	UserName string `json:"userName"` // 用户名
}

// LoginResMes 返回的时候的date
type LoginResMes struct {
	Code  int    `json:"code"` // 返回状态码 500 表示该用户未注册 200 表示登录成功
	Error string `json:"error"`
}
