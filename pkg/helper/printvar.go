package helper

import (
	"fmt"
	"reflect"
	"strings"
)

func printStruct(t reflect.Type, v reflect.Value, pc int) {
	fmt.Println("")
	for i := 0; i < t.NumField(); i++ {
		fmt.Print(strings.Repeat(" ", pc), t.Field(i).Name, ":")
		value := v.Field(i)
		PrintVar(value.Interface(), pc+2)
		fmt.Println("")
	}
}

func printArraySlice(v reflect.Value, pc int) {
	for j := 0; j < v.Len(); j++ {
		PrintVar(v.Index(j).Interface(), pc+2)
	}
}
func printMap(v reflect.Value, pc int) {
	for _, k := range v.MapKeys() {
		fmt.Println(v.MapIndex(k).Kind())
		PrintVar(k.Interface(), pc)
		PrintVar(v.MapIndex(k).Interface(), pc)
	}
}

func PrintVar(i interface{}, ident int) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {

		v = reflect.ValueOf(i).Elem()
		t = v.Type()
	}
	switch v.Kind() {
	case reflect.Array:
		printArraySlice(v, ident)
	case reflect.Chan:
		fmt.Println("Chan")
	case reflect.Func:
		fmt.Println("Func")
	case reflect.Interface:
		fmt.Println("Interface")
	case reflect.Map:
		printMap(v, ident)
	case reflect.Slice:
		printArraySlice(v, ident)
	case reflect.Struct:
		printStruct(t, v, ident)
	case reflect.UnsafePointer:
		fmt.Println("UnsafePointer")
	default:
		fmt.Print(strings.Repeat(" ", ident), v.Interface())
	}
}
