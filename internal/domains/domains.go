package domains

import (
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
	SubId       int64  `json:"id_sub"`
}

// SubscriptionUser defines model for SubscriptionUser.
type SubscriptionUser struct {
	Price       int64  `db:"price"`
	ServiceName string `db:"name"`
	SubId       int64  `db:"id_sub"`
}

// SubscriptionUserCreate defines model for SubscriptionUserCreate.
type SubscriptionUserCreate struct {
	StartDate time.Time `db:"start_date"`
	SubId     int64     `db:"id_sub"`
	UserId    string    `db:"id_user"`
}
