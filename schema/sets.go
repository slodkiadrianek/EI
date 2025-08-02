package schema

import z "github.com/Oudwins/zog"

type CreateSet struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryId  int    `json:"categoryId" binding:"required"`
}
func (c *CreateSet) Validate() (z.ZogIssueMap, error) {
	errMap := CreateSetSchema.Validate(c)
	if errMap != nil {
		return errMap, nil
	}
	return nil, nil
}
var CreateSetSchema = z.Struct(z.Shape{
	"name":        z.String().Required(),
	"description": z.String().Required(),
	"categoryId":  z.Int().Required(),
})