package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), false); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func MarshalJson(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), true); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// encode writes to buf an S-expression representation of v.
func encode(buf *bytes.Buffer, v reflect.Value, json bool) error {
	begin := byte('(')
	end := byte(')')
	sep := byte(' ')
	invalid := "nil"
	if json {
		begin = byte('{')
		end = byte('}')
		sep = byte(',')
		invalid = "null"
	}
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString(invalid)

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		if v.Int() != 0 {
			fmt.Fprintf(buf, "%d", v.Int())
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if v.Uint() != 0 {
			fmt.Fprintf(buf, "%d", v.Uint())
		}

	case reflect.String:
		if v.String() != "" {
			fmt.Fprintf(buf, "%q", v.String())
		}

	case reflect.Ptr:
		return encode(buf, v.Elem(), json)

	case reflect.Array, reflect.Slice: // (value ...)
		if json {
			buf.WriteByte('[')
		} else {
			buf.WriteByte(begin)
		}
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(sep)
			}
			if err := encode(buf, v.Index(i), json); err != nil {
				return err
			}
		}
		if json {
			buf.WriteByte(']')
		} else {
			buf.WriteByte(end)
		}

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte(begin)
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(sep)
			}
			if json {
				fmt.Fprintf(buf, "\"%s\":", v.Type().Field(i).Name)
			} else {
				fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			}
			if err := encode(buf, v.Field(i), json); err != nil {
				return err
			}
			if !json {
				buf.WriteByte(end)
			}
		}
		buf.WriteByte(end)

	case reflect.Map: // ((key value) ...)
		buf.WriteByte(begin)
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(sep)
			}
			if !json {
				buf.WriteByte(begin)
			}
			if err := encode(buf, key, json); err != nil {
				return err
			}
			if json {
				buf.WriteByte(':')
			} else {
				buf.WriteByte(sep)
			}
			if err := encode(buf, v.MapIndex(key), json); err != nil {
				return err
			}
			if !json {
				buf.WriteByte(end)
			}
		}
		buf.WriteByte(end)

	case reflect.Bool:
		if v.Bool() {
			buf.WriteString("t")
		}

	case reflect.Float32, reflect.Float64:
		if v.Float() != 0 {
			fmt.Fprintf(buf, "%f", v.Float())
		}

	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		if c != complex(0, 0) {
			fmt.Fprintf(buf, "#C(%f %f)", real(c), imag(c))
		}

	default: // chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
