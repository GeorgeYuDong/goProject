package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `json:"name,omitempty"`
	Sex  string `json:"sex,omitempty"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagstring := t.Field(i).Tag.Get("json")
		fmt.Println("json:", tagstring)
	}

}

func main() {
	var re resume
	findTag(&re)

}
