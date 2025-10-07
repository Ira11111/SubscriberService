package filter

import "time"

const (
	limitDefault  = 20
	offsetDefault = 0
)

type FilterOptions struct {
	// Пагинация
	Limit  int64
	Offset int64

	// Основные фильтры
	SubID   *int64
	UserID  *string
	SubName *string

	// Фильтры по дате
	StartDate *time.Time
	EndDate   *time.Time
}

type FilterBuilder struct {
	filter FilterOptions
}

func NewFilterBuilder() *FilterBuilder {
	return &FilterBuilder{
		filter: FilterOptions{},
	}
}

func (b *FilterBuilder) WithPagination(limit, offset int64) *FilterBuilder {
	if limit > 0 {
		b.filter.Limit = min(limit, 50) // защита от больших лимитов
	}
	if offset >= 0 {
		b.filter.Offset = offset
	}
	return b
}

func (b *FilterBuilder) WithSubID(subID int64) *FilterBuilder {
	b.filter.SubID = &subID
	return b
}

func (b *FilterBuilder) WithUserID(userID string) *FilterBuilder {
	b.filter.UserID = &userID
	return b
}

func (b *FilterBuilder) WithSubName(subName *string) *FilterBuilder {
	b.filter.SubName = subName
	return b
}

func (b *FilterBuilder) WithDateRange(startDate, endDate time.Time) *FilterBuilder {
	if !startDate.IsZero() {
		b.filter.StartDate = &startDate
	}
	if !endDate.IsZero() {
		b.filter.EndDate = &endDate
	}
	return b
}

func (b *FilterBuilder) Build() FilterOptions {
	return b.filter
}
