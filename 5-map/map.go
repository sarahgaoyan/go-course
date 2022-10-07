package main

import "fmt"

func printKV(map1 map[string]string) {
	for k, v := range map1 {
		fmt.Println("key is ", k, " value is ", v)
	}
}

func main() {
	map1 := make(map[int]string)
	map1[1] = "hello"
	fmt.Println(map1)

	map2 := map[string]string{
		"hello": "word",
		"foo":   "bar",
	}
	fmt.Println(map2)
	printKV(map2)

	fmt.Println("--------")

	delete(map2, "hello")
	map2["foo"] = "bbb"
	printKV(map2)
}
