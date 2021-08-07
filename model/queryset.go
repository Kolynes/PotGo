package model

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

type QuerySet struct {
	tableName string
	values    []*Model
	index     int64
	operation string
	operands  []interface{}
}

type IQuerySet interface {
	First() *Model
	Last() *Model
	Next() *Model
	Get() *Model
	Create() bool
	All() *QuerySet
	Filter(...map[string]interface{}) *QuerySet
	OrderBy(...map[string]interface{}) *QuerySet
	Execute()
	New() *Model
}

func (querySet *QuerySet) All() *QuerySet {
	if querySet.operation == "" {
		querySet.operation = fmt.Sprintf("SELECT * FROM %s", querySet.tableName)
	} else {
		querySet.operation = fmt.Sprintf("SELECT * FROM (%s) as %s", querySet.operation, querySet.tableName)
	}
	return querySet
}

func (querySet *QuerySet) Filter(filters ...map[string]interface{}) *QuerySet {
	whereClause, operands := createWhereClause(filters...)
	fields := getFields(querySet.New())
	if querySet.operation == "" {
		querySet.operation = fmt.Sprintf("SELECT %s FROM %s WHERE %s", strings.Join(fields, ", "), querySet.tableName, whereClause)
		querySet.operands = operands
	} else {
		querySet.operation = fmt.Sprintf("SELECT %s FROM (%s) as %s WHERE %s", strings.Join(fields, ", "), querySet.operation, querySet.tableName, whereClause)
		querySet.operands = append(querySet.operands, operands...)
	}
	return querySet
}

func (querySet *QuerySet) OrderBy(ordering ...string) *QuerySet {
	if querySet.operation == "" {
		panic("invalid query")
	} else {
		orderingClauses := []string{}
		for i := 0; i < len(ordering); i++ {
			if len(ordering[i]) == 0 {
				panic("invalid ordering")
			} else if ordering[i][0] == '-' {
				orderingClauses = append(orderingClauses, fmt.Sprintf("%s DESC", strings.TrimPrefix(ordering[i], "-")))
			} else if ordering[i][0] == '+' {
				orderingClauses = append(orderingClauses, fmt.Sprintf("%s ASC", strings.TrimPrefix(ordering[i], "+")))
			} else {
				orderingClauses = append(orderingClauses, fmt.Sprintf("%s ASC", ordering[i]))
			}
		}
		querySet.operation = fmt.Sprintf("%s ORDER BY %s", querySet.operation, strings.Join(orderingClauses, ", "))
	}
	return querySet
}

func (querySet *QuerySet) Get(filters ...map[string]interface{}) *Model {
	whereClause, operands := createWhereClause(filters...)
	fields := getFields(querySet.New())
	if querySet.operation == "" {
		querySet.operation = fmt.Sprintf("SELECT DISTINCT %s FROM %s WHERE %s", strings.Join(fields, ", "), querySet.tableName, whereClause)
		querySet.operands = operands
	} else {
		querySet.operation = fmt.Sprintf("SELECT DISTINCT %s FROM (%s) as %s WHERE %s", strings.Join(fields, ", "), querySet.operation, querySet.tableName, whereClause)
		querySet.operands = append(querySet.operands, operands...)
	}
	return querySet.First()
}

func (querySet *QuerySet) First() *Model {
	querySet.Execute()
	return querySet.values[0]
}

func (querySet *QuerySet) Next() {

}

func (querySet *QuerySet) Execute() {
	db, err := sql.Open()
	results, err := db.Query(querySet.operation)
	models := []*Model{}
	for results.Next() {
		model := querySet.New()
		modelReflection := reflect.ValueOf(model)
		var modelFieldAddresses []interface{} = make([]interface{}, 0)
		for i := 0; i < modelReflection.NumField(); i++ {
			modelFieldAddresses = append(modelFieldAddresses, modelReflection.Field(i).Addr().Interface())
		}
		results.Scan(modelFieldAddresses...)
		models = append(models, model)
	}
	querySet.values = models
}

func (querySet *QuerySet) New() *Model

func toModel(modelReflection *reflect.Value, values map[string]interface{}) {

}

func createWhereClause(filters ...map[string]interface{}) (string, []interface{}) {
	var elements []interface{}
	clause := ""
	orOperations := []string{}
	for filter := range filters {
		andOperations := []string{}
		for key, element := range filters[filter] {
			fieldOperatorSlice := strings.Split(key, "__")
			if len(fieldOperatorSlice) > 2 {
				panic(fmt.Sprintf("invalid lookup '%s'", key))
			} else {
				field := fieldOperatorSlice[0]
				var operation string
				var placeholder string
				if len(fieldOperatorSlice) == 1 {
					operation = "="
					placeholder = "?"
					elements = append(elements, element)
				} else {
					switch fieldOperatorSlice[1] {
					case "ne":
						operation = "!="
						placeholder = "?"
						elements = append(elements, element)
					case "in":
						operation = "IN"
						holders := []string{}
						tempElements := element.([]interface{})
						for range tempElements {
							holders = append(holders, "?")
						}
						placeholder = "(" + strings.Join(holders, ", ") + ")"
						elements = append(elements, tempElements...)
					case "gte":
						operation = ">="
						placeholder = "?"
						elements = append(elements, element)
					case "lte":
						operation = "<="
						placeholder = "?"
						elements = append(elements, element)
					case "gt":
						operation = ">"
						placeholder = "?"
						elements = append(elements, element)
					case "lt":
						operation = "<"
						placeholder = "?"
						elements = append(elements, element)
					case "contains":
						operation = "LIKE"
						placeholder = "%?%"
						elements = append(elements, element)
					default:
						panic(fmt.Sprintf("invalid lookup operator '%s'", fieldOperatorSlice[1]))
					}
				}
				andOperations = append(andOperations, fmt.Sprintf("%s %s %s", field, operation, placeholder))
			}
			orOperations = append(orOperations, strings.Join(andOperations, " and "))
		}
		clause = strings.Join(orOperations, " or ")
	}
	return clause, elements
}

func getFields(model *Model) []string {
	modelReflection := reflect.ValueOf(*model)
	fields := []string{}
	for i := 0; i < modelReflection.Type().NumField(); i++ {
		fields = append(fields, modelReflection.Type().Field(i).Name)
	}
	return fields
}
