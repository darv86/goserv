package utils

import (
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func GetUrlLastParam(url *url.URL) (int64, error) {
	path := strings.Split(url.Path, "/")
	urlParam := path[len(path)-1]
	param, err := strconv.ParseInt(urlParam, 10, 64)
	if err != nil {
		return 0, nil
	}
	return param, nil
}

func GetApiKey(headers http.Header) (string, error) {
	header := headers.Get("Authorization")
	if header == "" {
		return "", errors.New("there is no authorization header")
	}
	if !strings.HasPrefix(header, "ApiKey") {
		return "", errors.New("there is a wrong format for authorization header")
	}
	apiKey := strings.Split(header, " ")[1]
	return apiKey, nil
}

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
