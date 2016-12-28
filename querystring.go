package maestropanel

import (
    "reflect"
    "strings"
)

func convertToQueryString(obj interface{}) string  {
	data := make(map[string]string)

	objVal := reflect.ValueOf(obj)
	objType := reflect.TypeOf(obj)

	fieldCount := objVal.NumField()
	
	for i := 0; i<fieldCount; i++ {
		field := objVal.Field(i)
		
		tagName := objType.Field(i).Tag.Get("json")

		if tagName == "" {
			tagName = objType.Field(i).Name
		}

		v := toString(field)

		if v != "" {
			data[tagName] = v
		}	
	}

	str := make([]string,len(data))

	index := 0
	for key, val := range data {
		str[index] = key+"="+ val

		index++
	}

	return strings.Join(str,"&")
}