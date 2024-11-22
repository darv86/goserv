package utils

import (
	"reflect"
)

func GetStructTypeOf[T any](source interface{}) T {
	valSource := reflect.ValueOf(source)
	if valSource.Kind() == reflect.Pointer {
		valSource = valSource.Elem()
	}
	typSource := valSource.Type()
	var target T
	valTarget := reflect.ValueOf(&target).Elem()
	typTarget := valTarget.Type()
	fieldsTarget := reflect.VisibleFields(typTarget)
	for i := range fieldsTarget {
		fieldTarget := valTarget.Field(i)
		fieldSource := valSource.Field(i)
		if typSource.Field(i).Name == "ID" {
			fieldTarget.SetInt(fieldSource.FieldByName("Int64").Int())
			continue
		}
		fieldTarget.Set(fieldSource)
	}
	return target
}
