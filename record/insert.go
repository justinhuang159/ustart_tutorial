package record

import (
	"context"

	"github.com/sea350/ustart_tutorial/record/recordpb"
)

// Insert retreives all customer data based off of a username
func (record *Record) Insert(ctx context.Context, req *recordpb.InsertRequest) (*recordpb.InsertResponse, error) {

	intVal, err := strconv.Atoi(req.Rate)
	if err:=nil{
		return nil, err
	}

	err = rec.strg.Insert(ctx, req.CarID, req.UserID, req.DateStart, intVal)
	if err != nil{
		return nil, err
	}

	return recordpb.InsertResponse{}, nil

}
