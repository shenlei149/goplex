package cycle

import (
	"reflect"
	"unsafe"
)

func HasCycle(x interface{}) bool {
	seen := make(map[ptrAndType]bool)
	return hasCycle(reflect.ValueOf(x), seen)
}

type ptrAndType struct {
	x unsafe.Pointer
	t reflect.Type
}

func hasCycle(x reflect.Value, seen map[ptrAndType]bool) bool {
	if x.CanAddr() {
		p := ptrAndType{unsafe.Pointer(x.UnsafeAddr()), x.Type()}
		if seen[p] {
			return true // already seen
		}
		seen[p] = true
	}
	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return hasCycle(x.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if hasCycle(x.Index(i), seen) {
				return true
			}
		}
		return false

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if hasCycle(x.Field(i), seen) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range x.MapKeys() {
			if hasCycle(x.MapIndex(k), seen) || hasCycle(k, seen) {
				return true
			}
		}
		return false

	default:
		return false
	}
}
