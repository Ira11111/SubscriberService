package domains

import (
	"database/sql"
	"time"
)

type SubSum struct {
	EndDate   time.Time `db:"end_date"`
	StartDate time.Time `db:"start_date"`
	TotalSum  int64     `db:"total"`
	UserId    string    `db:"id_user"`
}

type Subscription struct {
	Price       int64  `db:"price"`
	ServiceName string `db:"name"`
	Id          int64  `db:"id"`
}

type SubscriptionUser struct {
	Id          int64        `db:"id"`
	Price       int64        `db:"price"`
	ServiceName string       `db:"name"`
	SubId       int64        `db:"id_sub"`
	UserId      string       `db:"id_user"`
	StartDate   time.Time    `db:"start_date"`
	EndDate     sql.NullTime `db:"end_date"`
}

type SubscriptionUserCreate struct {
	StartDate time.Time `db:"start_date"`
	SubId     int64     `db:"id_sub"`
	UserId    string    `db:"id_user"`
}
