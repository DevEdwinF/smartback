package utils

import (
	"fmt"
	"time"
)

func BuildFilters(field string, value interface{}, op string, where *string) {
	if v, ok := value.(string); ok {
		if len(*where) == 0 && len(v) > 0 {
			*where = fmt.Sprintf("%s ILIKE '%%%s%%'", field, value)
		} else if len(*where) > 0 && op == "AND" && len(v) > 0 {
			*where = fmt.Sprintf(" %s AND %s ILIKE '%%%s%%'", *where, field, value)
		} else if len(*where) > 0 && op == "OR" && len(v) > 0 {
			*where = fmt.Sprintf(" %s OR %s ILIKE '%%%s%%'", *where, field, value)
		}
	} else if v, ok := value.(time.Time); ok {
		if len(*where) == 0 && !v.IsZero() {
			*where = fmt.Sprintf("%s = '%v'", field, v.Format("2006-01-02"))
		} else if len(*where) > 0 && op == "AND" && !v.IsZero() {
			*where = fmt.Sprintf(" %s AND %s = '%v'", *where, field, v.Format("2006-01-02"))
		} else if len(*where) > 0 && op == "OR" && !v.IsZero() {
			*where = fmt.Sprintf(" %s OR %s = '%v'", *where, field, v.Format("2006-01-02"))
		}
	} else if v, ok := value.(*bool); ok {
		if len(*where) == 0 && v != nil {
			*where = fmt.Sprintf("%s = '%v'", field, *v)
		} else if len(*where) > 0 && op == "AND" && v != nil {
			*where = fmt.Sprintf(" %s AND %s = '%v'", *where, field, *v)
		} else if len(*where) > 0 && op == "OR" && v != nil {
			*where = fmt.Sprintf(" %s OR %s = '%v'", *where, field, *v)
		}
	}
}
