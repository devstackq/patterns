package liskov

import "log"

/*
наследник, дополняет, а не замещает базовый класс
*/

type Drone interface {
	fly()
}

type DroneXSt struct {
	name, model string
}

func NewDroneX(name, model string) *DroneXSt {
	return &DroneXSt{name, model}
}

func (d *DroneXSt) prepare() {
	log.Println("prepare drone ")
}

func (d *DroneXSt) check() {
	log.Println("check oil, engine, etc")
}

func (d *DroneXSt) fly() {
	d.prepare()
	d.check()
	log.Println("drone x ", d.name, d.model)
}

type DroneYSt struct {
	DroneXSt
}

func NewDroneY(name, model string) *DroneYSt {
	return &DroneYSt{
		DroneXSt{
			name, model,
		},
	}
}

func (d *DroneYSt) runAway() {
	log.Println("get out")
}

func (d *DroneYSt) fly() {
	d.prepare()
	d.check()
	log.Println("drone y ", d.model, d.name)
	// дополняет
	d.runAway()
}

func test() {
	dY := NewDroneY("airba", "x87")
	dY.fly()
}
