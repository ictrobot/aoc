package parse

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"sync"
)

var PlaceholderDoesNotMatch = errors.New("string does not match placeholder match or regex tag")

type Placeholder struct{}

type placeholderAttrs struct {
	optional, multiple bool
	match              *string
	regexp             *regexp.Regexp
}

var placeholderType = reflect.TypeOf(Placeholder{})
var placeholderCache = sync.Map{}

// called by reflectValue
func handlePlaceholder(sf reflect.StructField, s []string) ([]string, error) {
	if sf.Name != "_" {
		panic(fmt.Sprintf("parse.Placeholder fields in structs must use the blank identifier"))
	}
	if sf.Type != placeholderType {
		panic(fmt.Sprintf("blank identifiers in structs must have the parse.Placeholder type"))
	}

	attrs := parsePlaceholderTag(sf.Tag)

	if len(s) == 0 {
		if attrs.optional {
			return nil, nil
		}
		return nil, NotEnoughStrings
	}

	consumed := 0
	for i := 0; i < len(s) && (i == 0 || attrs.multiple); i++ {
		if attrs.match != nil {
			if s[i] == *attrs.match {
				consumed++
				continue
			}
			break
		}

		if attrs.regexp != nil {
			if attrs.regexp.MatchString(s[i]) {
				consumed++
				continue
			}
			break
		}

		// neither match nor regexp set, consume anything
		consumed++
	}

	if consumed > 0 {
		return s[consumed:], nil
	} else if attrs.optional {
		return s, nil
	} else {
		return nil, PlaceholderDoesNotMatch
	}
}

func parsePlaceholderTag(tag reflect.StructTag) placeholderAttrs {
	if v, ok := placeholderCache.Load(tag); ok {
		return v.(placeholderAttrs)
	}

	var attrs placeholderAttrs
	if flags := tag.Get("flags"); flags != "" {
		for _, flag := range strings.Split(flags, ",") {
			switch flag {
			case "optional":
				attrs.optional = true
			case "multiple":
				attrs.multiple = true
			default:
				panic(fmt.Sprintf("unknown parse.Placeholder flag `%s`", flag))
			}
		}
	}

	if match, ok := tag.Lookup("match"); ok {
		attrs.match = &match
	}

	if regex, ok := tag.Lookup("regex"); ok {
		r, err := regexp.Compile(regex)
		if err != nil {
			panic(fmt.Sprintf("invalid placeholder regex tag `%s`: %v", regex, err))
		}
		attrs.regexp = r
	}

	if attrs.match != nil && attrs.regexp != nil {
		panic("parse.Placeholder tag may only contain one of match & regexp")
	}

	placeholderCache.Store(tag, attrs)
	return attrs
}
