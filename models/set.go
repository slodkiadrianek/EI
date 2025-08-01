package models

type Set struct{
	Id int
	Name string
	Description string
	CategoryId  int
}

func NewSet(name , description string, categoryId, id int) *Set{
	return &Set{
		Id:id,
		Name: name,
		Description: description,
		CategoryId: categoryId,
	}
}