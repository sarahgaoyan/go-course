package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	mv := Movie{"snow white", 10}
	jsonStr, err := json.Marshal(mv)
	if err != nil {
		fmt.Println("json marshall error: ", err)
		return
	}
	fmt.Printf("json marshall string: %s\n", jsonStr)

	m := Movie{}
	err = json.Unmarshal(jsonStr, &m)
	if err != nil {
		fmt.Println("json unMarshall error: ", err)
		return
	}
	fmt.Printf("json unMarshall : %v\n", m)
}
