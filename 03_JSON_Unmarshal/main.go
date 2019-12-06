package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`

	bs := []byte(s)

	fmt.Println(bs)

	// people := []person{}
	// or
	var people []person

	e := json.Unmarshal(bs, &people)

	if e != nil {
		fmt.Println("Error ", e)
	}

	for _, v := range people {
		fmt.Println(v.First)
	}

	fmt.Println(people)
}
