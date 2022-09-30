package filter

import "reflect"

func Filter(slice interface{}, filter interface{}) interface{} {
	sv := reflect.ValueOf(slice)
	fv := reflect.ValueOf(filter)

	sliceLen := sv.Len()
	out := reflect.MakeSlice(sv.Type(), 0, sliceLen)
	for i := 0; i < sliceLen; i++ {
		curVal := sv.Index(i)
		values := fv.Call([]reflect.Value{curVal})
		if values[0].Bool() {
			out = reflect.Append(out, curVal)
		}
	}

	return out.Interface()
}

func FilterString(slice []string, filter func(string) bool) []string {
	out := make([]string, 0)

	for _, elem := range slice {
		if filter(elem) {
			out = append(out, elem)
		}
	}

	return out
}
