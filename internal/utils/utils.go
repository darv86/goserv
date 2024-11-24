package utils

import (
	"reflect"
)

func GetStructTypeOf[T any](source interface{}) T {
	valSource := reflect.ValueOf(source)
	if valSource.Kind() == reflect.Pointer {
		valSource = valSource.Elem()
	}
	var target T
	valTarget := reflect.ValueOf(&target).Elem()
	typTarget := valTarget.Type()
	fieldsTarget := reflect.VisibleFields(typTarget)
	for i := range fieldsTarget {
		fieldTarget := valTarget.Field(i)
		fieldSource := valSource.Field(i)
		fieldTarget.Set(fieldSource)
	}
	return target
}
