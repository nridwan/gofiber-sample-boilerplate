package response

import "gopkg.in/guregu/null.v3"

//Response : Base response
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

//Meta : Meta response
type Meta struct {
	Code    null.Int    `json:"code"`
	Message null.String `json:"message"`
}

func CreateMetaResponse(code int, message string) Response {
	return Response{
		Meta: Meta{
			Code:    null.IntFrom(int64(code)),
			Message: null.StringFrom(message),
		},
	}
}

func CreateResponse(code int, message string, data interface{}) Response {
	return Response{
		Meta: Meta{
			Code:    null.IntFrom(int64(code)),
			Message: null.StringFrom(message),
		},
		Data: data,
	}
}
