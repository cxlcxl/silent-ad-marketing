package res_code

import "ad-marketing/api/common/response"

func ResErr(code int32, msg string) *response.ResMsg {
	return &response.ResMsg{Code: code, Msg: msg}
}

func ResOk() *response.ResMsg {
	return &response.ResMsg{Code: 0, Msg: "OK"}
}
