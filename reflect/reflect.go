package reflect

import "reflect"

func walk(x any, fn func(in string)) {
	val := reflect.ValueOf(x)
	in := val.Field(0)
	fn(in.String())
}
