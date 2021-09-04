package response

import "gopkg.in/guregu/null.v3"

//Response : Base response
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

//Error : error detail
type Error struct {
	Code   string `json:"code"`
	Reason string `json:"reason"`
}

//Meta : Meta response
type Meta struct {
	Code    null.Int    `json:"code"`
	Message null.String `json:"message"`
	Errors  []Error     `json:"errors"`
}

func CreateMetaResponse(code int, message string, errors []Error) Response {
	return Response{
		Meta: Meta{
			Code:    null.IntFrom(int64(code)),
			Message: null.StringFrom(message),
			Errors:  errors,
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
