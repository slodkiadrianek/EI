package DTO

type Set struct {
	Name string
	Description string
	CategoryId int
}

func NewSet( name , description string, categoryId int) *Set {
	return &Set{
		Name: name,
		Description: description,
		CategoryId: categoryId,
	}
}
