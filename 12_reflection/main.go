package main

import "reflect"

func main() {
	// entry point
}
func walk(x interface{}, fn func(input string)) {
	val := getValue(x) // struct
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}
	/*
		switch val.Kind() {
		case reflect.Struct:
			for i := 0; i < val.NumField(); i++ {
				walk(val.Field(i).Interface(), fn)
			}
		case reflect.Slice:
			for i := 0; i < val.Len(); i++ {
				walk(val.Index(i).Interface(), fn)
			}
		case reflect.String:
			fn(val.String())
		}

	*/

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}

	case reflect.Chan:
		// `for { ... }` is  an infinite loop until we break
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, result := range valFnResult {
			walkValue(result)
		}

	}

	for i := 0; i < numberOfValues; i++ {
		walkValue(getField(i))
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
