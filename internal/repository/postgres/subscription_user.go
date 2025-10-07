package postgres

import (
	d "SubscriberService/internal/domains"
	"SubscriberService/internal/filter"
	"context"
)

func (s *Storage) SaveUserSub(ctx context.Context, userSub *d.SubscriptionUserCreate) (*d.SubscriptionUser, error) {
	return nil, nil
}

func (s *Storage) GetUserSubs(ctx context.Context, options *filter.FilterOptions) ([]d.SubscriptionUser, error) {
	return nil, nil
}
