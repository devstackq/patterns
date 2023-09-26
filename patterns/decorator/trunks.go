package decorator

import "fmt"

type Trunks struct {
	Base IHuman
}

func NewTrunks() *Trunks {
	return &Trunks{}
}

func (t *Trunks) PutOn() error {
	fmt.Println("now, put on - trunks ")
	return t.Base.PutOn()
}
