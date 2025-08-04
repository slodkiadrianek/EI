package schema

import z "github.com/Oudwins/zog"

type CreateCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *CreateCategory) Validate() (z.ZogIssueMap, error) {
	errMap := CreateCategorySchema.Validate(c)
	if errMap != nil {
		return errMap, nil
	}
	return nil, nil
}

var CreateCategorySchema = z.Struct(z.Shape{
	"name":        z.String().Required(),
	"description": z.String().Required(),
})

type GetCategory struct {
	CategoryId int `json:"categoryId" uri:"categoryId"`
}	

func (g *GetCategory) Validate() (z.ZogIssueMap, error) {
	errMap := GetCategorySchema.Validate(g)
	if errMap != nil {
		return errMap, nil
	}
	return nil, nil
}	
var GetCategorySchema = z.Struct(z.Shape{
	"categoryId": z.Int().Required(),
})
