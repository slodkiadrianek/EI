package DTO

type Category struct{
	Name string
	Description string
}

func NewCategory(name string, description string) *Category{
	return &Category{
		Name:name,
		Description: description,
	}
}