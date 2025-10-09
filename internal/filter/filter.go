package filter

import "time"

type QueryOperation string

const (
	limitDefault  = 20
	offsetDefault = 0
)

var (
	ILikeOP QueryOperation = "ILIKE"
	LikeOP  QueryOperation = "LIKE"
	EqOP    QueryOperation = "="
	GrOP    QueryOperation = ">"
	LessOp  QueryOperation = "<"
)

type condition struct {
	Field     string
	Value     interface{}
	Operation QueryOperation // "=", "ILIKE"
}

type FilterOptions struct {
	// Пагинация
	Limit  int64
	Offset int64

	// Фильтры по дате
	StartDate *time.Time
	EndDate   *time.Time

	// Динамические фильтры
	conditions []*condition
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

func (b *FilterBuilder) WithEqualCondition(field string, value interface{}) *FilterBuilder {
	if value != nil && value != "" && value != 0 {
		b.filter.conditions = append(b.filter.conditions, &condition{
			Field:     field,
			Value:     value,
			Operation: EqOP,
		})
	}
	return b
}

func (b *FilterBuilder) WithGreaterCondition(field string, value interface{}) *FilterBuilder {
	if value != nil && value != "" && value != 0 {
		b.filter.conditions = append(b.filter.conditions, &condition{
			Field:     field,
			Value:     value,
			Operation: GrOP,
		})
	}
	return b
}

func (b *FilterBuilder) WithLessCondition(field string, value interface{}) *FilterBuilder {
	if value != nil && value != "" && value != 0 {
		b.filter.conditions = append(b.filter.conditions, &condition{
			Field:     field,
			Value:     value,
			Operation: LessOp,
		})
	}
	return b
}

func (b *FilterBuilder) WithILikeCondition(field string, value interface{}) *FilterBuilder {
	if value != nil && value != "" && value != 0 {
		b.filter.conditions = append(b.filter.conditions, &condition{
			Field:     field,
			Value:     value,
			Operation: ILikeOP,
		})
	}
	return b
}

func (b *FilterBuilder) WithLikeCondition(field string, value interface{}) *FilterBuilder {
	if value != nil && value != "" && value != 0 {
		b.filter.conditions = append(b.filter.conditions, &condition{
			Field:     field,
			Value:     value,
			Operation: LikeOP,
		})
	}
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
