package main

import (
	"builder/querybuilder"
	"fmt"
)

func main() {
	var (
		xx = querybuilder.QueryBuilder{}
	)
	x := xx.Select("select * from").Table("Table1").Where("id=1").Find()
	y := xx.Update("Update ").Table("Table1").Where("id=1").Find()
	fmt.Println("Query Result: ", x)
}
