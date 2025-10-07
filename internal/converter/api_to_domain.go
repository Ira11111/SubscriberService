package converter

import (
	"SubscriberService/api/generated"
	"SubscriberService/internal/domains"
	"database/sql"
)

func ToDomainSubscription(apiSub *generated.Subscription) *domains.Subscription {
	return &domains.Subscription{
		Price:       apiSub.Price,
		ServiceName: apiSub.ServiceName,
		Id:          int64(*apiSub.SubId),
	}
}

func ToDomainSubscriptionUser(apiUser *generated.SubscriptionUser) *domains.SubscriptionUser {
	var endDate sql.NullTime
	if apiUser.EndDate != nil {
		endDate = sql.NullTime{
			Time:  apiUser.EndDate.Time,
			Valid: true,
		}
	} else {
		endDate = sql.NullTime{Valid: false}
	}
	return &domains.SubscriptionUser{
		Price:       apiUser.Price,
		ServiceName: apiUser.ServiceName,
		UserId:      apiUser.UserId,
		StartDate:   apiUser.StartDate.Time,
		EndDate:     endDate,
		SubId:       *apiUser.SubId,
	}
}

func ToDomainSubscriptionUserCreate(apiCreate *generated.SubscriptionUserCreate) *domains.SubscriptionUserCreate {
	return &domains.SubscriptionUserCreate{
		SubId:     int64(apiCreate.SubId),
		UserId:    apiCreate.UserId,
		StartDate: apiCreate.StartDate.Time,
	}
}
