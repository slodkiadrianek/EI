package models

type Element struct{
	Id int `json:"id" sql:"id"`
	English string `json:"english" sql:"english"`
	Polish string 	`json:"polish" sql:"polish"`
	ExampleSentence string `json:"exampleSentence" sql:"example_sentence"`
	Synonym string `json:"synonym" sql:"synonym"`
	SetId int `json:"setId" sql:"set_id"`
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