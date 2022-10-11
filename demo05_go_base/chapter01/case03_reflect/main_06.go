//我们不需要go.mod，那么我们文件与文件之间是独立的，不能相互调用

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Movie struct {
	Title    string
	SubTitle string
	Year     int
	Color    bool
	Actor    map[string]string
	Oscars   []string
}

func main() {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		SubTitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	var v interface{}
	v = strangelove
	display("strangelove", reflect.ValueOf(&strangelove))
	println("-------------------------------------------")
	display("strangelove", reflect.ValueOf(v)) //这个v的动态类型还是Movie
	println("-------------------------------------------")
	display("strangelove", reflect.ValueOf(&v)) //接口的指针的动态类型的接口
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		// Len returns v's length.
		// It panics if v's Kind is not Array, Chan, Map, Slice, or String.

		// Index returns v's i'th element.
		// It panics if v's Kind is not Array, Slice, or String or i is out of range.
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s %s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatAtom2(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		// Elem returns the value that the interface v contains
		// or that the pointer v points to.
		// It panics if v's Kind is not Interface or Pointer.
		// It returns the zero Value if v is nil.
		if v.IsNil() {
			fmt.Printf("%s = nil", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: //基本类型、函数、通道
		fmt.Printf("%s = %s\n", path, formatAtom2(v))
	}
}

func formatAtom2(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.Invalid:
		return "invalid"
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default: //reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
