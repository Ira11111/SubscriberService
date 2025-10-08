package converter

import (
	"SubscriberService/api/generated"
	"SubscriberService/internal/domains"
	"database/sql"
)

func ToDomainCreateSubscription(apiSub *generated.SubscriptionCreate) *domains.Subscription {
	return &domains.Subscription{
		ServiceName: apiSub.ServiceName,
		Price:       apiSub.Price,
	}
}

func ToDomainSubscriptionUserUpdate(apiSub *generated.SubscriptionUserUpdate) *domains.SubscriptionUser {
	endDate := sql.NullTime{
		Time:  apiSub.EndDate.Time,
		Valid: true,
	}
	return &domains.SubscriptionUser{
		EndDate:   endDate,
		StartDate: apiSub.StartDate.Time,
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
		SubId:       apiUser.SubId,
	}
}

func ToDomainCreateSubscriptionUser(apiUser *generated.SubscriptionUserCreate) *domains.SubscriptionUser {
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
		SubId:     apiUser.SubId,
		UserId:    apiUser.UserId,
		StartDate: apiUser.StartDate.Time,
		EndDate:   endDate,
	}
}
