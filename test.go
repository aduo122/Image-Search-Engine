package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type Arc struct {
	Head     string
	Modifier string
}

func main() {
	arc := Arc{"saw", "He"}
	fmt.Printf("%v\n", arc)
	fmt.Printf("%+v\n", arc)
	fmt.Printf("%#v\n", arc)

	// Convert structs to JSON.
	data, err := json.Marshal(arc)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", data)
	res := string(data)
	fmt.Println(reflect.TypeOf(data[0]))
		fmt.Println(reflect.TypeOf(res))
	fmt.Println(res)
}
