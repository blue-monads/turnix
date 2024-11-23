package autoquery

import (
	"fmt"
	"strings"
)

type Builder struct {
	Table  string
	Params LoaderParams
}

func NewBuilder(table string, params LoaderParams) *Builder {
	return &Builder{
		Table:  table,
		Params: params,
	}
}

func (qb *Builder) Build() (string, []any) {
	query := fmt.Sprintf("SELECT %s FROM %s", qb.buildColumns(), qb.Table)

	whereClause, whereArgs := qb.buildWhereClause()
	if whereClause != "" {
		query += " WHERE " + whereClause
	}

	if qb.Params.OrderBy != "" {
		query += qb.buildOrderBy()
	}

	query += qb.buildPagination()

	return query, whereArgs
}

func (qb *Builder) buildColumns() string {
	if len(qb.Params.ActiveColumns) > 0 {
		return strings.Join(qb.Params.ActiveColumns, ", ")
	}
	return "*"
}

func (qb *Builder) buildWhereClause() (string, []any) {
	var conditions []string
	var args []any

	for colKey, filters := range qb.Params.FilterModels {
		for _, filter := range filters {
			condition, filterArgs := qb.buildFilterCondition(colKey, filter)
			if condition != "" {
				conditions = append(conditions, condition)
				args = append(args, filterArgs...)
			}

		}

	}

	if qb.Params.OrderBy == "" {
		if qb.Params.LoadType != "previous" {
			conditions = append(conditions, "id > ?")
			args = append(args, qb.Params.MaxId)
		} else {
			conditions = append(conditions, "id < ?")
			args = append(args, qb.Params.MaxId)
		}
	}

	return strings.Join(conditions, " AND "), args
}

func (qb *Builder) buildFilterCondition(column string, filter FilterModel) (string, []any) {
	switch filter.Operator {
	case OperatorEqual:
		return fmt.Sprintf("%s = ?", column), []any{filter.Value}
	case OperatorNotEqual:
		return fmt.Sprintf("%s != ?", column), []any{filter.Value}
	case OperatorContains:
		return fmt.Sprintf("%s LIKE ?", column), []any{"%" + filter.Value[0] + "%"}
	case OperatorNotContains:
		return fmt.Sprintf("%s NOT LIKE ?", column), []any{"%" + filter.Value[0] + "%"}
	case OperatorStartsWith:
		return fmt.Sprintf("%s LIKE ?", column), []any{filter.Value[0] + "%"}
	case OperatorEndsWith:
		return fmt.Sprintf("%s LIKE ?", column), []any{"%" + filter.Value[0]}
	case OperatorIsNull:
		return fmt.Sprintf("%s IS NULL", column), nil
	case OperatorIsNotNull:
		return fmt.Sprintf("%s IS NOT NULL", column), nil
	case OperatorBetween:
		return fmt.Sprintf("%s BETWEEN ? AND ?", column), []any{filter.Value[0], filter.Value[1]}
	case OperatorNotBetween:
		return fmt.Sprintf("%s NOT BETWEEN ? AND ?", column), []any{filter.Value[0], filter.Value[1]}
	case OperatorGreaterThan:
		return fmt.Sprintf("%s > ?", column), []any{filter.Value}
	case OperatorLessThan:
		return fmt.Sprintf("%s < ?", column), []any{filter.Value}
	case OperatorGreaterThanOrEqual:
		return fmt.Sprintf("%s >= ?", column), []any{filter.Value}
	case OperatorLessThanOrEqual:
		return fmt.Sprintf("%s <= ?", column), []any{filter.Value}
	}

	return "", nil
}

func (qb *Builder) buildOrderBy() string {
	direction := "ASC"
	if strings.ToUpper(qb.Params.OrderDirection) == "DESC" {
		direction = "DESC"
	}
	return fmt.Sprintf(" ORDER BY %s %s", qb.Params.OrderBy, direction)
}

func (qb *Builder) buildPagination() string {
	if qb.Params.OrderBy == "" {
		// ID-based offset pagination
		if qb.Params.PageSize != 0 {
			return fmt.Sprintf(" LIMIT %d", qb.Params.PageSize)
		}
	} else {
		// Page-based pagination
		if qb.Params.PageId != 0 && qb.Params.PageSize != 0 {
			offset := (qb.Params.PageId - 1) * qb.Params.PageSize
			return fmt.Sprintf(" LIMIT %d OFFSET %d", qb.Params.PageSize, offset)
		}
	}
	return ""
}
