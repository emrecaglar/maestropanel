package maestropanel

import (
    "reflect"
    "strconv"
)

func toString(field reflect.Value) []string  {
		var v = make([]string,0)
		switch field.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = append(v,strconv.FormatInt(field.Int(), 10))
		case uint, uint8, uint16, uint32, uint64:
			v = append(v,strconv.FormatUint(field.Uint(), 10))
		case float32:
			v = append(v,strconv.FormatFloat(field.Float(), 'f', 4, 32))
		case float64:
			v = append(v,strconv.FormatFloat(field.Float(), 'f', 4, 64))
		case []byte:
			v = append(v,string(field.Bytes()))
		case []string:
			arr := reflect.ValueOf(field.Interface())

			for i := 0; i<arr.Len(); i++ {
				v = append(v, arr.Index(i).String())
			}
		case string:
			v = append(v,field.String())
		}

		return v
}