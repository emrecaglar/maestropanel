package maestropanel

import (
    "reflect"
    "strconv"
)

func toString(field reflect.Value) string  {
		var v string
		switch field.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(field.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(field.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(field.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(field.Float(), 'f', 4, 64)
		case []byte:
			v = string(field.Bytes())
		case string:
			v = field.String()
		}

		return v
}