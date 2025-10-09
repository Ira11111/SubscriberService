package filter

import (
	"fmt"
	"strings"
	"time"
)

type QueryBuilder struct {
	baseQuery  string
	conditions []string
	args       []interface{}
}

func BuildQuery(baseQuery string, filter *FilterOptions) (string, []interface{}) {
	builder := &QueryBuilder{
		baseQuery: baseQuery,
		args:      make([]interface{}, 0),
	}

	for _, cond := range filter.conditions {
		condStr := fmt.Sprintf("%s %s", cond.Field, cond.Operation)
		if cond.Operation == ILikeOP || cond.Operation == LikeOP {
			//отдельный метод чтобы валидировать аргумент и обернуть в %%
			builder.addTextSearchCondition(condStr+" $%d", cond.Value)
		} else {
			builder.addCondition(condStr+" $%d", cond.Value)
		}
	}

	// добавляем даты если нужно и пагинацию
	builder.addCondition("su.start_date >= $%d", filter.StartDate)
	builder.addCondition("su.end_date <= $%d", filter.EndDate)

	// Собираем полный запрос
	query := builder.baseQuery

	if len(builder.conditions) > 0 {
		query += " WHERE " + strings.Join(builder.conditions, " AND ")
	}

	if filter.Limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", len(builder.args)+1)
		builder.args = append(builder.args, filter.Limit)
	}

	if filter.Offset > 0 {
		query += fmt.Sprintf(" OFFSET $%d", len(builder.args)+1)
		builder.args = append(builder.args, filter.Offset)
	}
	fmt.Println(query, builder.args)
	return query, builder.args
}

func (b *QueryBuilder) addCondition(condition string, value interface{}) {
	if b.isValidValue(value) {
		b.conditions = append(b.conditions, fmt.Sprintf(condition, len(b.args)+1))
		b.args = append(b.args, b.formatValue(value))
	}
}

func (b *QueryBuilder) addTextSearchCondition(condition string, value interface{}) {
	if b.isValidValue(value) {
		formattedValue := b.formatValue(value)
		if str, ok := formattedValue.(string); ok {
			formattedValue = "%" + str + "%"
		}
		b.conditions = append(b.conditions, fmt.Sprintf(condition, len(b.args)+1))
		b.args = append(b.args, formattedValue)
	}
}

// isValidValue - проверяет что значение валидно для фильтрации
func (b *QueryBuilder) isValidValue(value interface{}) bool {
	switch v := value.(type) {
	case *int64:
		return v != nil && *v > 0
	case *string:
		return v != nil && *v != ""
	case *time.Time:
		return v != nil && !v.IsZero()
	case int64:
		return v > 0
	case string:
		return v != ""
	case time.Time:
		return !v.IsZero()
	default:
		return false
	}
}

// formatValue - форматирует значение для SQL
func (b *QueryBuilder) formatValue(value interface{}) interface{} {
	switch v := value.(type) {
	case *string:
		return *v
	case *int64:
		return *v
	case *time.Time:
		return *v
	default:
		return v
	}
}
