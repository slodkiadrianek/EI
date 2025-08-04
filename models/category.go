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


type CategoryWithSet struct {
	CategoryId int
	CategoryName string
	CategoryDescription string
	SetId int
	SetName string
	SetDescription string
}

func NewCategoryWithSet(categoryId int, categoryName string, categoryDescription string, setId int, setName string, setDescription string) *CategoryWithSet{
	return &CategoryWithSet{
		CategoryId: categoryId,
		CategoryName: categoryName,
		CategoryDescription: categoryDescription,
		SetId: setId,
		SetName: setName,
		SetDescription: setDescription,
	}
}