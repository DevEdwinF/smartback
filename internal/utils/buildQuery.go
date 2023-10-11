package utils

import "fmt"

func BuildFilters(field, value, op string, where *string) {
	if len(*where) == 0 && len(value) > 0 {
		*where = fmt.Sprintf("%s ILIKE '%%%s%%'", field, value)
	} else if len(*where) > 0 && op == "AND" && len(value) > 0 {
		*where = fmt.Sprintf(" %s AND %s ILIKE '%%%s%%'", *where, field, value)
	} else if len(*where) > 0 && op == "OR" && len(value) > 0 {
		*where = fmt.Sprintf(" %s OR %s ILIKE '%%%s%%'", *where, field, value)
	}
}
