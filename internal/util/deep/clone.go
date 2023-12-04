package deep

import (
	"fmt"
	"reflect"
)

func Clone[T any](in *T) *T {
	if in == nil {
		return nil
	}

	var out T
	pIn := reflect.ValueOf(in)
	pOut := reflect.ValueOf(&out)

	seen := map[uintptr]reflect.Value{pIn.Pointer(): pOut}
	pOut.Elem().Set(cloneValue(pIn.Elem(), seen))

	return &out
}

func cloneValue(in reflect.Value, seen map[uintptr]reflect.Value) reflect.Value {
	t := in.Type()
	switch t.Kind() {
	case reflect.Bool, reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128:
		return in

	case reflect.Array:
		l := in.Len()
		out := reflect.New(t).Elem()
		for i := 0; i < l; i++ {
			out.Index(i).Set(cloneValue(in.Index(i), seen))
		}
		return out

	case reflect.Map:
		out := reflect.MakeMapWithSize(t, in.Len())
		iter := in.MapRange()
		for iter.Next() {
			out.SetMapIndex(cloneValue(iter.Key(), seen), cloneValue(iter.Value(), seen))
		}
		return out

	case reflect.Pointer:
		if in.IsNil() {
			return reflect.Zero(t)
		}

		ptr := in.Pointer()
		if out, ok := seen[ptr]; ok {
			return out
		} else {
			out = reflect.New(t.Elem())
			seen[ptr] = out
			out.Elem().Set(cloneValue(in.Elem(), seen))
			return out
		}

	case reflect.Slice:
		l := in.Len()
		out := reflect.MakeSlice(t, l, l)
		for i := 0; i < l; i++ {
			out.Index(i).Set(cloneValue(in.Index(i), seen))
		}
		return out

	case reflect.Struct:
		l := t.NumField()
		out := reflect.New(t).Elem()
		for i := 0; i < l; i++ {
			sf := t.Field(i)
			if sf.Name == "_" && sf.Type.Kind() == reflect.Struct && sf.Type.NumField() == 0 {
				// allow blank identifier fields with zero size struct types
				continue
			}
			if !sf.IsExported() {
				panic(fmt.Sprintf("unexported field: field %d `%s` in `%s`", i, sf.Name, t.Name()))
			}

			out.Field(i).Set(cloneValue(in.Field(i), seen))
		}
		return out

	default:
		panic(fmt.Sprintf("unsupported type: %v", t))
	}
}
