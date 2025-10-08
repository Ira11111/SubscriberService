package domains

import (
	"database/sql"
	"time"
)

type Subscription struct {
	Price       int64  `db:"price"`
	ServiceName string `db:"name"`
	Id          int64  `db:"id"`
}

type SubscriptionUser struct {
	Price       int64        `db:"price"`
	ServiceName string       `db:"name"`
	SubId       int64        `db:"id_sub"`
	UserId      string       `db:"id_user"`
	StartDate   time.Time    `db:"start_date"`
	EndDate     sql.NullTime `db:"end_date"`
}
