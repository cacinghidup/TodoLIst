package result

type SuccessResult struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type ErrorResult struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
