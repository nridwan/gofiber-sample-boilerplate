package baseresponse

import "gopkg.in/guregu/null.v3"

//Response : Base response
type Response struct {
	Meta Meta `json:"meta"`
}

//Meta : Meta response
type Meta struct {
	Code null.Int `json:"code"`
}
