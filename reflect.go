package main

import (
	"fmt"
	"reflect"
)

/**
	Reflect adalah package untuk mengetahui struktur dari coding
	contoh :
	Mau tau apa saja field dari struct
	Mau tau apa saja tag field dan isinya (eg. required = true)
	Bisa digunakan untuk validasi seperti dibawah
*/

type Sample struct {
	Name string `required:"true" max:"10"`
}

type ContohLagi struct {
	Name string `required:"true"`
	Description string `required:"true"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func callReflect() {
	sample := Sample{"Eko"}

	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	sample.Name = ""
	fmt.Println(IsValid(sample))

	contoh := ContohLagi{"Eko", "Oke"}
	fmt.Println(IsValid(contoh))
}