package DTO

type Element struct {
	English string
	Polish string
	ExampleSentence string
	Synonym string
	SetId int
}

func NewElement(english string, polish string, exmapleSentence string, synonym string, setId int) *Element{
	return &Element{
		English: english,
		Polish: polish,
		Synonym: synonym,
		SetId: setId,
	}
}