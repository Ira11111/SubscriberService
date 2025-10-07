package converter

import (
	"SubscriberService/api/generated"
	"SubscriberService/internal/domains"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func ToAPISubscription(domainSub *domains.Subscription) *generated.Subscription {
	return &generated.Subscription{
		Price:       domainSub.Price,
		ServiceName: domainSub.ServiceName,
		SubId:       &domainSub.Id,
	}
}

func ToAPISubscriptionUser(domainUser *domains.SubscriptionUser) *generated.SubscriptionUser {
	return &generated.SubscriptionUser{
		Price:       domainUser.Price,
		ServiceName: domainUser.ServiceName,
		SubId:       &domainUser.SubId,
	}
}

func ToAPISubSum(domainSum *domains.SubSum) *generated.SubSum {
	return &generated.SubSum{
		UserId:    &domainSum.UserId,
		TotalSum:  &domainSum.TotalSum,
		StartDate: &openapi_types.Date{domainSum.StartDate},
		EndDate:   &openapi_types.Date{domainSum.EndDate},
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
