package response

type Data any

type Response struct {
	Code    int    `json:"code"`    // 业务响应状态码
	Message string `json:"message"` // 提示信息
	Data    Data   `json:"data"`    // 数据
}
