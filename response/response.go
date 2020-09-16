package response

type Responser interface {
	GetHTTPStatusCode() int
	GetResponse() interface{}
}

type Response struct {
	HTTPStatusCode int         `json:"-"`
	Code           string      `json:"code" example:"0"`
	Description    string      `json:"description" example:"success"`
	Data           interface{} `json:"data,omitempty"`
}

type ErrResponse struct {
	HTTPStatusCode int    `json:"-"`
	Code           string `json:"code" example:"4001"`
	Description    string `json:"description" example:"Cannot bind request."`
}

func NewResponseError(httpStatusCode int, code, desc string) *ErrResponse {
	return &ErrResponse{
		HTTPStatusCode: httpStatusCode,
		Code:           code,
		Description:    desc,
	}
}

func NewResponse(httpStatusCode int, code, desc string, data interface{}) *Response {
	return &Response{
		HTTPStatusCode: httpStatusCode,
		Code:           code,
		Description:    desc,
		Data:           data,
	}
}

func (r *Response) GetHTTPStatusCode() int {
	return r.HTTPStatusCode
}

func (r *Response) GetResponse() interface{} {
	return r
}

func (e *ErrResponse) GetHTTPStatusCode() int {
	return e.HTTPStatusCode
}

func (e *ErrResponse) GetResponse() interface{} {
	return e
}
