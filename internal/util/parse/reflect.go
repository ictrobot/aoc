package parse

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var NotEnoughStrings = errors.New("not enough strings provided")
var TooManyStrings = errors.New("too many strings provided")

func Reflect[T any](s []string) (T, error) {
	var zero, result T
	pv := reflect.ValueOf(&result)

	s, ignored, err := reflectValue(pv.Elem(), s, nil)
	if errors.Is(err, NotEnoughStrings) {
		// add parsed result for easier debugging
		return zero, fmt.Errorf("%w\n\nresult:\n%+v", err, result)
	} else if err != nil {
		return zero, err
	} else if len(s) > 0 {
		// add ignored errors and parsed result for easier debugging
		ignoredMsg := strings.Builder{}
		if len(ignored) > 0 {
			ignoredMsg.WriteString("\n\nthe following errors were ignored during parsing:")
			for _, i := range ignored {
				ignoredMsg.WriteString("\n- ")
				ignoredMsg.WriteString(i.Error())
			}
		}

		return zero, fmt.Errorf("%w: %d left%s\n\nresult:\n%+v", TooManyStrings, len(s), ignoredMsg.String(), result)
	}

	return result, nil
}

func MustReflect[T any](s []string) T {
	r, err := Reflect[T](s)
	if err != nil {
		panic(fmt.Sprintf("failed to parse via reflection: %v", err))
	}
	return r
}

func reflectValue(v reflect.Value, s []string, ignored []error) ([]string, []error, error) {
	t := v.Type()
	switch t.Kind() {

	// Invalid

	case reflect.Bool:
		if len(s) == 0 {
			return nil, ignored, NotEnoughStrings
		}

		if b, err := Bool(s[0]); err != nil {
			return nil, ignored, err
		} else {
			v.SetBool(b)
			return s[1:], ignored, nil
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if len(s) == 0 {
			return nil, ignored, NotEnoughStrings
		}

		// fast path for single digits
		if len(s[0]) == 1 && s[0][0] >= '0' && s[0][0] <= '9' {
			v.SetInt(int64(s[0][0] - '0'))
			return s[1:], ignored, nil
		}

		if i, err := strconv.ParseInt(s[0], 0, t.Bits()); err != nil {
			return nil, ignored, err
		} else {
			v.SetInt(i)
			return s[1:], ignored, nil
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if len(s) == 0 {
			return nil, ignored, NotEnoughStrings
		}

		// fast path for single digits
		if len(s[0]) == 1 && s[0][0] >= '0' && s[0][0] <= '9' {
			v.SetUint(uint64(s[0][0] - '0'))
			return s[1:], ignored, nil
		}

		if i, err := strconv.ParseUint(s[0], 0, t.Bits()); err != nil {
			return nil, ignored, err
		} else {
			v.SetUint(i)
			return s[1:], ignored, nil
		}

	// Uintptr

	case reflect.Float32, reflect.Float64:
		if len(s) == 0 {
			return nil, ignored, NotEnoughStrings
		}

		if f, err := strconv.ParseFloat(s[0], t.Bits()); err != nil {
			return nil, ignored, err
		} else {
			v.SetFloat(f)
			return s[1:], ignored, nil
		}

	// Complex64
	// Complex128

	case reflect.Array:
		var err error
		for i := 0; i < v.Len(); i++ {
			s, ignored, err = reflectValue(v.Index(i), s, ignored)
			if err != nil {
				return nil, ignored, fmt.Errorf("reading array element %d: %w", i, err)
			}
		}
		return s, ignored, nil

	// Chan
	// Func
	// Interface
	// Map

	case reflect.Pointer:
		// treat pointers as optional fields
		if len(s) == 0 {
			// no strings left, must be nil
			v.SetZero()
			return nil, ignored, nil
		}
		if v.IsZero() {
			v.Set(reflect.New(v.Type().Elem()))
		}

		s2, ignored, err := reflectValue(v.Elem(), s, ignored)
		if err != nil {
			// append error to ignored, set pointer to nil and continue
			// this allows using e.g. `[]struct { Opcode string, Operand *int }` to represent
			// one or more opcodes, which may or may not be followed by an operand
			v.SetZero()
			ignored = append(ignored, err)
			return s, ignored, nil
		}
		return s2, ignored, nil

	case reflect.Slice:
		// greedily consuming matching strings until error

		// N strings makes at most N elements
		sliceV := reflect.MakeSlice(t, len(s), len(s))

		var i int
		for len(s) > 0 {
			var s2 []string
			var err error

			s2, ignored, err = reflectValue(sliceV.Index(i), s, ignored)
			if err != nil {
				ignored = append(ignored, err)
				break
			} else if len(s) == len(s2) {
				// prevent infinite loop from trying to parse []*int with []string{"a"}
				// panic("reading slice element consumed no strings")
				ignored = append(ignored, fmt.Errorf("reading slice element %d consumed no strings", i))
				break
			}

			i++
			s = s2
		}

		if i == 0 {
			v.SetZero()
			return s, ignored, nil
		}

		if i < sliceV.Len() {
			// handle case where each element took >1 string, or the loop ended early
			sliceV = sliceV.Slice(0, i)
		}

		v.Set(sliceV)
		return s, ignored, nil

	case reflect.String:
		if len(s) == 0 {
			return nil, ignored, NotEnoughStrings
		}
		v.Set(reflect.ValueOf(s[0]))
		return s[1:], ignored, nil

	case reflect.Struct:
		var err error
		for i := 0; i < t.NumField(); i++ {
			sf := t.Field(i)

			if sf.Name == "_" || sf.Type == placeholderType {
				s, err = handlePlaceholder(sf, s)
				if err != nil {
					return nil, ignored, fmt.Errorf("reading struct field %d `%s` in `%s`: %w", i, sf.Name, t.Name(), err)
				}
				continue
			}

			if !sf.IsExported() {
				panic(fmt.Sprintf("unexported field: field %d `%s` in `%s`", i, sf.Name, t.Name()))
			}

			s, ignored, err = reflectValue(v.Field(i), s, ignored)
			if err != nil {
				return nil, ignored, fmt.Errorf("reading struct field %d `%s` in `%s`: %w", i, sf.Name, t.Name(), err)
			}
		}
		return s, ignored, nil

	// UnsafePointer

	default:
		panic(fmt.Sprintf("unsupported type: %v", t))
	}
}
