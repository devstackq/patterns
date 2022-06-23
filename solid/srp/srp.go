package srp

import "log"

/*
2 example:
1 func = 1 responsibility

It means that your struct must only have one responsibility.
It also means that when you create your struct, you should look at 3 things :

    Is my method a behavior of my struct ? (Is this behavior the Single Responsability of my struct ?)
    Is my struct tightly coupled with another struct ?
    Could my method be applied differently ?
*/

// any struct - own - realize;
type Drone interface {
	Fly()
}

type DroneSt struct {
	name string
}

func NewDrone(name string) *DroneSt {
	return &DroneSt{name: name}
}

//------------------------------------
/*maybe use - own struct*/
func (dr *DroneSt) checkBattery() {
	log.Println("check battery state...")
}

func (dr *DroneSt) checkPropeller() {
	log.Println("check propeller state...")
}

func (dr *DroneSt) healthCheck() {
	dr.checkBattery()
	dr.checkPropeller()
}

//------------------------------------

/*maybe use - own struct*/
func (dr *DroneSt) prepare() {
	dr.fillOil()
	dr.turnOnEngine()
}

func (dr *DroneSt) turnOnEngine() {
	log.Println("turn on engine at ", dr.name)
}

func (dr *DroneSt) fillOil() {
	log.Println("filling oil...")
}

//------------------------------------

func (dr *DroneSt) takeOff() {
	dr.turnOffEngine()
}

func (dr *DroneSt) turnOffEngine() {
	log.Println("turn off engine..")
}

/*---------------------------------
another function - call here/ composition
*/
func (dr *DroneSt) Fly() {
	dr.healthCheck()
	dr.prepare()
	log.Println("flying drone ... ", dr.name)
	dr.takeOff()
}
