package car

import (

	"github.com/justinhuang159/ustart_tutorial/car/carpb"
)

// Toggles a car's availability
func (car *Car) ToggleAvailable(req *carpb.ToggleRequest) {
	//search for car, if it exists, update availability 

	bool response = CheckCarAvailability(req)

	if response == true {
		UpdateCarAvailability(false)
	} else {
		UpdateCarAvailability(true)
	}
}

