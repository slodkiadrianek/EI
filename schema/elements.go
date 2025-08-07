package schema

import z "github.com/Oudwins/zog"

type ElementById struct {
	SetId int `json:"set_id" uri:"setId"`
}

func (e *ElementById) Validate() (z.ZogIssueMap, error) {
	errMap := ElementsByIdSchema.Validate(e)
	if errMap != nil {
		return errMap, nil
	}
	return nil, nil
}

var ElementsByIdSchema = z.Struct(z.Shape{
	"setId": z.Int().Required(),
})
