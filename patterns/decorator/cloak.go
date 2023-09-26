package decorator

import "fmt"

type Cloak struct {
	Base IHuman
}

func NewCloak() *Cloak {
	return &Cloak{}
}

func (c *Cloak) PutOn() error {
	fmt.Println(" now, put on - cloak ")
	return c.Base.PutOn()
}
