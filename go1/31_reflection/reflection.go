package main

import (
	"fmt"
	"reflect"
)

type order struct {
	id         int
	customerID int
}

type employee struct {
	id      int
	name    string
	address string
	salary  int
	country string
}

// reflect.Type: concrete type
// reflect.Value: represented value
// reflect.Kind: specific kind of the type
// NumField(): returns number of fields
// Field(index): returns the reflect.Value of the i.field
// Int/String(): extract the reflect value as <type>
func createQuery(q interface{}) string {
	typeOf := reflect.TypeOf(q)
	kind := typeOf.Kind()
	value := reflect.ValueOf(q)
	fmt.Println("Type ", typeOf)
	fmt.Println("Kind ", kind)
	fmt.Println("Value ", value)

	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		v := reflect.ValueOf(q)
		query := fmt.Sprintf("insert into %s(", t)

		for i := 0; i < v.NumField(); i++ {
			if i == 0 {
				query = fmt.Sprintf("%s%s", query, typeOf.Field(i).Name)
			} else {
				query = fmt.Sprintf("%s, %s", query, typeOf.Field(i).Name)
			}
		}

		query = fmt.Sprintf("%s) values(", query)

		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			}
		}

		return fmt.Sprintf("%s)", query)
	}

	return ""
}

// ability of a program to inspect its variables and values at runtime and find their type
// reflect: runtime reflections
func main() {
	// SIMPLE EXAMPLE - type of a variable at compile time
	i := 10
	fmt.Printf("%d %T\n", i, i)

	// EXAMPLE - type of variable at runtime
	o := order{1234, 567}
	fmt.Println(createQuery(o))

	// ATTENTION: "CLEAR IS BETTER THAN CLEVER. REFLECTION IS NEVER CLEAR."
}
