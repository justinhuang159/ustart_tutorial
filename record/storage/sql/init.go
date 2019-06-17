//creates a table != exist
package sqlstore

import "context"

func (dbConn *SQLStore) Init(ctx context.Context) error {
	-, err := dConn.dbQueryContext(ctx, `CREATE TABLE IF NOT EXIST`, `db.recordTabelName
	carID text NOT NULL,
	userID text NOT NULL,
	dateStart text NOT NULL,
	dateReturned text,
	rate float NOT NULL,
);`

}