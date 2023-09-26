package decorator

import "fmt"

/*
idea: set to handler, service etc -> Many functionality; - without use Inherit;
пример  с одеждой - челоек = Base -> conreceteDecorotaro -люая одежда - метод Одеваться -
1.можно надеть плавки для плавания, также 2.плащь от защиты дождя, etc
*/

// base class
type IHuman interface {
	PutOn() error
}

// default entiry
type Client struct {
	Base IHuman
}

func (c *Client) PutOn() error {
	fmt.Println("default Golyi")
	return nil
}

func Test() {

	defaultHuman := &Client{}

	//prepare matreshka, conrete decorator; can skip
	//trunks := &Trunks{
	//	Base: defaultHuman,
	//}

	cloak := &Cloak{
		Base: defaultHuman,
	}

	sock := NewSock(cloak)
	//called like Chain - 1 method; FILO
	if err := sock.PutOn(); err != nil {
		fmt.Printf(" err decorator: %v \n ", err)
	}

}
