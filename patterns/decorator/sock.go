package decorator

import "fmt"

type Sock struct {
	Base IHuman
}

func NewSock(human IHuman) *Sock {
	return &Sock{
		Base: human,
	}
}
func (s *Sock) PutOn() error {
	fmt.Println(" now, put on - sock ")
	return s.Base.PutOn()
}
