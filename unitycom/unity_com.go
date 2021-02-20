package unitycom

import "net/http"

func GetSendClientComDefault() SendClientCom {

	com := SendClientCom{}
	com.Code = http.StatusOK

	return com
}

//s-c返回给客户端的统一组件
type SendClientCom struct {
	Code int         `json:"status_code"` //1 操作状态ID
	Text string      `json:"status_text"` //2 操作状态描述
	Data interface{} `json:"data"`        //3 返回数据
}
