package maestropanel

import (
    "reflect"
    "strings"
)

func convertToQueryString(obj interface{}) string  {
	data := make([]string, 0)

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

		if len(v) > 0 {
			for _, item := range v {
				if item != "" {
					data = append(data, tagName + "=" + item)
				}
			}
		}	
	}

	return strings.Join(data,"&")
}