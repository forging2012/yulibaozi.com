package transport

import "reflect"

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

func ToSlice(container interface{}) []interface{} {
	val := reflect.ValueOf(container)
	sInd := reflect.Indirect(val)
	if sInd.Kind() != reflect.Slice {
		return []interface{}{container}
	}
	l := sInd.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = sInd.Index(i).Interface()
	}
	return ret
}

func ExtractFieldValues(source interface{}, field string) ([]interface{}, error) {

	val := reflect.ValueOf(source)
	sInd := reflect.Indirect(val)
	if sInd.Kind() != reflect.Slice {
		return make([]interface{}, 0), errors.New("source must be slice type")
	}

	defer Recover("extract field values failure !")

	var slice = ToSlice(source)
	var filedValues = make([]interface{}, 0)
	for _, obj := range slice {
		var objVal = reflect.ValueOf(obj)
		var val = objVal.FieldByName(field).Interface()
		filedValues = append(filedValues, val)
	}

	return filedValues, nil

}

func SliceToString(a interface{}, sep string) string {
	var strSlice []string
	var slice = ToSlice(a)
	for _, v := range slice {
		switch v.(type) {
		case string:
			strSlice = append(strSlice, v.(string))
			continue
		case int, int8, int16, int32, int64:
			strV := strconv.FormatInt(reflect.ValueOf(v).Int(), 10)
			strSlice = append(strSlice, strV)
			continue
		default:
			panic("params type not supported")
		}
	}
	return strings.Join(strSlice, ",")
}

func IsSlice(a interface{}) bool {
	val := reflect.ValueOf(a)
	sInd := reflect.Indirect(val)

	if sInd.Kind() == reflect.Slice {
		return true
	}

	return false
}

func IsMap(a interface{}) bool {
	val := reflect.ValueOf(a)
	sInd := reflect.Indirect(val)
	if sInd.Kind() == reflect.Map {
		return true
	}
	return false
}

func IsEmpty(a interface{}) bool {
	if nil == a {
		return true
	}

	if val, ok := a.(string); ok && len(val) == 0 {
		return true
	}

	val := reflect.ValueOf(a)
	sInd := reflect.Indirect(val)

	if sInd.Kind() == reflect.Slice && len([]interface{}{a}) == 0 {
		return true
	}

	if sInd.Kind() == reflect.Map && len(sInd.MapKeys()) == 0 {
		return true
	}

	return false

}
