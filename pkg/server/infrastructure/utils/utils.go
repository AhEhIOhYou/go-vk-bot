package utils

import (
	"log"
	"reflect"
)

func ConvertMapToStruct(m map[string]interface{}, s interface{}) error {
	stValue := reflect.ValueOf(s).Elem()
	sType := stValue.Type()
	for i := 0; i < sType.NumField(); i++ {
		field := sType.Field(i)
		if value, ok := m[field.Name]; ok {
			log.Println(reflect.ValueOf(value))
			stValue.Field(i).Set(reflect.ValueOf(value))
		}
	}
	return nil
}
