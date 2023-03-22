package utils

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

func QueryFormatPagination(values []any) (pagination string, err error) {
	pagination = " LIMIT "

	for _, val := range values {
		t := reflect.TypeOf(val).Kind()
		if t != reflect.Int {
			err = errors.New("pagination value is not an integer")
			return
		} else {
			pagination += fmt.Sprintf(`%v,`, val)
		}
	}

	pagination = pagination[:len(pagination)-1]

	return pagination, err
}

func QueryFormatIn(column []string, values []any, uniqueConstraint bool) string {
	var in string
	if uniqueConstraint {
		in += fmt.Sprintf("WHERE %v IN (", column[0])
	} else {
		in += fmt.Sprintf("AND %v IN (", column[0])
	}

	if len(column) == 1 {
		for range column {
			for _, val := range values {
				t := reflect.TypeOf(val).Kind()
				if t == reflect.String {
					in += fmt.Sprintf(`"%v",`, val)
				} else {
					in += fmt.Sprintf(`%v,`, val)
				}
			}

		}

		in = in[:len(in)-1]
		in += ")"

	} else {
		for i := range column {
			for j, val := range values {
				if i == j {
					t := reflect.TypeOf(val).Kind()
					if t == reflect.String {
						in += fmt.Sprintf(`"%v,"`, val)
					} else {
						in += fmt.Sprintf(`%v,`, val)
					}
				}
			}
			in = in[:len(in)-1]
			in += ")"
			if i+1 < len(column) {
				in += fmt.Sprintf(" %v IN (", column[i+1])
			}
		}

	}

	return in
}

func QueryFormatWheres(column []string, values []any) string {
	var wheres string = " WHERE "
	for i, v := range column {
		for j, val := range values {
			if i == j {
				if i != 0 {
					wheres += " AND "
				}
				t := reflect.TypeOf(val).Kind()
				if t == reflect.String {
					wheres += fmt.Sprintf(`%v = "%v"`, v, val)
				} else {
					wheres += fmt.Sprintf(`%v = %v`, v, val)
				}
			}
		}
	}

	// wheres = wheres[:len(wheres)-1]

	return wheres
}

func QueryFormatUpdates(field []string, fieldValue []any) string {
	var updates string
	for i, v := range field {
		for j, val := range fieldValue {
			if i == j {
				t := reflect.TypeOf(val).Kind()
				if t == reflect.String {
					updates += fmt.Sprintf(`%v = "%v",`, v, val)
				} else {
					updates += fmt.Sprintf(`%v = %v,`, v, val)
				}
			}
		}
	}

	updates = updates[:len(updates)-1]

	return updates
}

func QueryFormatter(field []string, fieldValue []any, columns []string, values []any, pageValues []any, inColumns []string, inValues []any, order string) (paginations string, wheres string, updates string, in string) {
	var errList []error
	var err error
	if len(pageValues) > 0 {
		paginations, err = QueryFormatPagination(pageValues)

		if err != nil {
			errList = append(errList, err)
		}
	}

	if len(inColumns) > 0 && len(inValues) > 0 && len(field) > 0 {
		in = QueryFormatIn(inColumns, inValues, true)
	} else if len(inColumns) > 0 && len(inValues) > 0 && len(field) == 0 {
		in = QueryFormatIn(inColumns, inValues, false)
	}

	if len(columns) > 0 && len(values) > 0 && len(columns) == len(values) {
		wheres = QueryFormatWheres(columns, values)
	} else if len(columns) != len(values) {
		errList = append(errList, errors.New("wrong parameters at query wheres formatter"))
	}

	if len(field) > 0 && len(fieldValue) > 0 && len(field) == len(fieldValue) {
		updates = QueryFormatUpdates(field, fieldValue)
	} else if len(field) != len(fieldValue) {
		errList = append(errList, errors.New("wrong parameters at query updates formatter"))
	}

	if len(errList) > 0 {
		log.Printf("QF 01: %v", errList)
	}

	return
}
