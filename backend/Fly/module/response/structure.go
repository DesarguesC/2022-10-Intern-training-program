package response

type Err_msg struct {
	Code int64       `json:code`
	Msg  string      `json:msg`
	Data interface{} `json:data`
}
