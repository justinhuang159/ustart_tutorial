package car

import "errors"

var (
	//errCustomerExists This user already has a customer
	errCarExists = errors.New("This car already exists")

	//errProblemLoadingCustomer This user already has a customer
	errProblemLoadingCar = errors.New("There was a problem loading one or more cars in this list")
)
