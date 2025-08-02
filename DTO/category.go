package DTO

type Category struct{
	Name string `json:"name"`
	Description string `json:"description"`
}

func NewCategory(name string, description string) *Category{
	return &Category{
		Name:name,
		Description: description,
	}
}