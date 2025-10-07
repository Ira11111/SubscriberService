package converter

import (
	"SubscriberService/api/generated"
	"SubscriberService/internal/domains"
)

func ToDomainSubscription(apiSub *generated.Subscription) *domains.Subscription {
	if apiSub == nil {
		return nil
	}
	domainSub := &domains.Subscription{
		Price:       apiSub.Price,
		ServiceName: apiSub.ServiceName,
	}

	if apiSub.SubId != nil {
		domainSub.Id = int64(*apiSub.SubId)
	}

	return domainSub
}

func ToDomainSubscriptionUser(apiUser *generated.SubscriptionUser) *domains.SubscriptionUser {
	if apiUser == nil {
		return nil
	}

	domainUser := &domains.SubscriptionUser{
		Price:       apiUser.Price,
		ServiceName: apiUser.ServiceName,
	}

	if apiUser.SubId != nil {
		domainUser.SubId = int64(*apiUser.SubId)
	}

	return domainUser
}

//// ToDomainSubscriptionSlice преобразует slice API Subscription в slice доменных Subscription
//func ToDomainSubscriptionSlice(apiSubs []*generated.Subscription) []*domains.Subscription {
//	if apiSubs == nil {
//		return nil
//	}
//
//	domainSubs := make([]*domains.Subscription, len(apiSubs))
//	for i, apiSub := range apiSubs {
//		domainSubs[i] = ToDomainSubscription(apiSub)
//	}
//	return domainSubs
//}
//
//// ToDomainSubscriptionUserSlice преобразует slice API SubscriptionUser в slice доменных SubscriptionUser
//func ToDomainSubscriptionUserSlice(apiUsers []*generated.SubscriptionUser) []*domains.SubscriptionUser {
//	if apiUsers == nil {
//		return nil
//	}
//
//	domainUsers := make([]*domains.SubscriptionUser, len(apiUsers))
//	for i, apiUser := range apiUsers {
//		domainUsers[i] = ToDomainSubscriptionUser(apiUser)
//	}
//	return domainUsers
//}
//
//// ToDomainSubscriptionFromRequest преобразует тело запроса создания в доменную модель
//func ToDomainSubscriptionFromRequest(apiReq *generated.PostSubscriptionsJSONRequestBody) *domains.Subscription {
//	if apiReq == nil {
//		return nil
//	}
//
//	apiSub := generated.Subscription(*apiReq)
//	return ToDomainSubscription(&apiSub)
//}
//
//// ToDomainSubscriptionUserFromRequest преобразует тело запроса создания пользовательской подписки
//func ToDomainSubscriptionUserFromRequest(apiReq *generated.PostSubscriptionsUsersJSONRequestBody) *domains.SubscriptionUserCreate {
//	if apiReq == nil {
//		return nil
//	}
//
//	apiCreate := generated.SubscriptionUserCreate(*apiReq)
//	return ToDomainSubscriptionUserCreate(&apiCreate)
//}
//
//// ToDomainSubscriptionUserCreate преобразует API SubscriptionUserCreate в доменную SubscriptionUserCreate
//func ToDomainSubscriptionUserCreate(apiCreate *generated.SubscriptionUserCreate) *domains.SubscriptionUserCreate {
//	if apiCreate == nil {
//		return nil
//	}
//
//	return &domains.SubscriptionUserCreate{
//		UserID:    apiCreate.UserId,
//		SubID:     int64(apiCreate.SubId),
//		StartDate: apiCreate.StartDate.Time(), // преобразование openapi_types.Date в time.Time
//	}
//}
