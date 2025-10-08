package converter

import (
	"SubscriberService/api/generated"
	"SubscriberService/internal/domains"

	"github.com/oapi-codegen/runtime/types"
)

func ToAPISubscription(domainSub *domains.Subscription) *generated.Subscription {
	return &generated.Subscription{
		Price:       domainSub.Price,
		ServiceName: domainSub.ServiceName,
		SubId:       &domainSub.Id,
	}
}

func ToAPISubscriptionUser(domainUser *domains.SubscriptionUser) *generated.SubscriptionUser {
	var endDate *types.Date
	if domainUser.EndDate.Valid {
		endDate = &types.Date{Time: domainUser.EndDate.Time}
	}
	return &generated.SubscriptionUser{
		Price:       domainUser.Price,
		ServiceName: domainUser.ServiceName,
		UserId:      domainUser.UserId,
		StartDate:   types.Date{domainUser.StartDate},
		EndDate:     endDate,
		SubId:       domainUser.SubId,
	}

}

func ToAPISubscriptionSlice(domainSubs []domains.Subscription) []*generated.Subscription {
	apiSubs := make([]*generated.Subscription, len(domainSubs))
	for i, domainSub := range domainSubs {
		apiSubs[i] = ToAPISubscription(&domainSub)
	}
	return apiSubs
}

func ToAPISubscriptionUserSlice(domainUsers []domains.SubscriptionUser) []*generated.SubscriptionUser {
	apiUsers := make([]*generated.SubscriptionUser, len(domainUsers))
	for i, domainUser := range domainUsers {
		apiUsers[i] = ToAPISubscriptionUser(&domainUser)
	}
	return apiUsers
}
