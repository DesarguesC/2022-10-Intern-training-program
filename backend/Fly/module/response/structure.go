package response

type err_meg struct {
	Msg string `json:msg`
	Code int64 `json:code`
}