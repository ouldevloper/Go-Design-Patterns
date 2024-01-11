package querybuilder

import "fmt"

type QueryBuilder struct {
	tableName    string
	selectClause string
	updateClause string
	deleteClause string
	whereClause  string
	sql          string
}

func (builder *QueryBuilder) Table(table string) *QueryBuilder {
	builder.tableName = table
	return builder
}
func (builder *QueryBuilder) Select(s string) *QueryBuilder {
	builder.selectClause = s
	return builder
}
func (builder *QueryBuilder) Update(s string) *QueryBuilder {
	builder.updateClause = s
	return builder
}

func (builder *QueryBuilder) Delete(s string) *QueryBuilder {
	builder.deleteClause = s
	return builder
}

func (builder *QueryBuilder) Where(s string) *QueryBuilder {
	builder.whereClause = s
	return builder
}

func (builder *QueryBuilder) Find() string {
	builder.sql = fmt.Sprintf("%s%s%s %s WHERE %s limit 1", builder.selectClause, builder.updateClause, builder.deleteClause, builder.tableName, builder.whereClause)
	return builder.sql
}

func (builder *QueryBuilder) Get() string {
	builder.sql = fmt.Sprintf("%s%s%s %s WHERE %s ", builder.selectClause, builder.updateClause, builder.deleteClause, builder.tableName, builder.whereClause)
	return builder.sql
}
