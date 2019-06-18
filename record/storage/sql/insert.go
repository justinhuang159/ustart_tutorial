package sqlstore

import {
	"context"
	"fmt"
	"time"
}

func (dbConn *SQLStore) Insert(ctx context.Context, carID, userID. startDate string, rate int) error{
	queryString :+ fmt.Sprintf(
		"INSERT INTO %s (carID, userID, startDate, rate) VALUES ('%s','%s','%s','+strconv.Itoa(rate)');",
		dbConn.recordTableName, carId, userID, startDate)
	
	_, err := dbConn.db.QueryContext(ctx, queryString)
}