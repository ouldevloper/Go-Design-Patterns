package querybuilder

type IQueryBuilder interface {
	Select() IQueryBuilder
	Where(s string) IQueryBuilder
	Get() IQueryBuilder
	Find() IQueryBuilder
	Update()
	Delete()
}
