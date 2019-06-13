package car

import (
	"context"

	"github.com/justinhuang159/ustart_tutorial/car/carpb"
)

// Register is a generic register function that registers a user in a database
func (car *Car) Register(ctx context.Context, req *carpb.RegisterRequest) (*carpb.RegisterResponse, error) {

	uuid := randString(32)

	_, err := car.strg.Lookup(ctx, uuid)
	if err != nil && err != car.strg.ErrUserDoesNotExist() {
		return nil, err
	}
	if err == nil {
		return nil, errCarExists
	}

	err = car.strg.Register(ctx, uuid, req.FirstName, req.LastName, req.DOB)
	if err != nil {
		return nil, err
	}

	return &carpb.RegisterResponse{}, nil

}
