package models

type Elements struct{
	Id int
	English string
	Polish string
	ExampleSentence string
	Synonym string
	SetId int
}

func NewElement(id, setId int, english, polish, exmapleSentence, synonym string) *Element{
	return &Elements{
		Id:id,
		English: english,
		Polish: polish,
		ExampleSentence: exmapleSentence,
		Synonym: synonym,
		SetId: setId,
	}
}