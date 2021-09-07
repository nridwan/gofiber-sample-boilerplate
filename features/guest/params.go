package guest

import "gopkg.in/guregu/null.v3"

type paramApps struct {
	Alias  null.String `json:"alias"`
	Appkey null.String `json:"appkey"`
}
