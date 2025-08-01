package models

type Element struct{
	Id int
	English string
	Polish string
	ExampleSentence string
	Synonym string
	SetId int
}

func NewElement(id, setId int, english, polish, exmapleSentence, synonym string) *Element{
	return &Element{
		Id:id,
		English: english,
		Polish: polish,
		ExampleSentence: exmapleSentence,
		Synonym: synonym,
		SetId: setId,
	}
}