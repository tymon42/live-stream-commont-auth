// Code generated by goctl. DO NOT EDIT.
package types

type ApplyRequest struct {
	Buid        uint64 `json:"buid"`
	Client_uuid string `json:"client_uuid"`
}

type ApplyResponse struct {
	Vcode string `json:"vcode"`
}

type SubmitRequest struct {
	Buid  uint64 `json:"buid"`
	Vcode string `json:"vcode"`
}

type SubmitResponse struct {
}

type StatusRequest struct {
	Buid        uint64 `form:"buid"`
	Client_uuid string `form:"client_uuid"`
	Vcode       string `form:"vcode"`
	Count       uint64 `form:"count"`
}

type StatusResponse struct {
	Verify_count int64 `json:"verify_count"`
}
