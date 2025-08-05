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

type SetWithElements struct {
	Id int `json:"id" sql:"id"`
	Name string `json:"name" sql:"name"`
	Description string `json:"description" sql:"description"`
	CategoryId int `json:"categoryId" sql:"categoryId"`
	English string `json:"english"`
	Polish string `json:"polish"`
	ExampleSentence string `json:"exampleSentence"`
	Synonym string `json:"synonym"`
}