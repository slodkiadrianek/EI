package models

type Category struct{
	Id int
	Name string
	Description string
}

func NewCategory(id int, name, description string) *Category{
	return &Category{
		Id:id,
		Name: name,
		Description: description,
	}
}