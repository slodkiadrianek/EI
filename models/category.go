package models

type Category struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Description string 	`json:"description"` 
}

func NewCategory(id int, name, description string) *Category{
	return &Category{
		Id:id,
		Name: name,
		Description: description,
	}
}