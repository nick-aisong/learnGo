// 9-1 reflect
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4

	fmt.Println("type:", reflect.TypeOf(x))

	v := reflect.ValueOf(x) // 获得的是x的副本
	fmt.Println("type:", v.Type())

	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	// v.Set(4.1)
	// panic: reflect.Value.SetFloat using unaddressable value

	p := reflect.ValueOf(&x) // 注意：得到X的地址
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	v1 := p.Elem()
	fmt.Println("settability of v1:", v1.CanSet())
	v1.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)

	type T struct {
		A int
		B string
	}
	t := T{203, "mh203"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
