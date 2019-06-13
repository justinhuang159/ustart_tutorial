package car

import (
	"context"

	"github.com/justinhuang159/ustart_tutorial/car/carpb"
)

// Toggles a car's availability
func (car *Car) ToggleAvailable(ctx context.Context, req *carpb.ToggleRequest) (*carpb.ToggleResponse, error) {

	cust, err := car.strg.Lookup(ctx, req.CID)

	// cust, err := customer.strg.Lookup(ctx, req.UUID)
	// if err != nil {
	// 	return nil, err
	// }

	// return &customerpb.PullResponse{CustomerProfile: &cust}, nil
	return &carpb.ToggleResponse{}, nil

}