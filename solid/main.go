package main

import (
	"solid/open_closed"
	"solid/srp"
)

func main() {
	// srp 1 case: 1 func = 1 responsibility
	dr := srp.NewDrone("aero smith x2")
	dr.Fly()

	open_closed.Test()

	// open closed: - each struct - realize own; method Fly
}
