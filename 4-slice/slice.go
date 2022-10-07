package main

import "fmt"

func changeSlice(a []int) {
	for i, v := range a {
		// slice 是指针传递，这将改变slice里的值
		a[i] = v + 1
	}
}

func main() {
	// --- 指针引用 ---
	var a = []int{1, 2, 3}
	changeSlice(a)
	for index, value := range a {
		fmt.Println("index = ", index, ", value = ", value)
	}

	// --- 声明方式 ---
	b := make([]int, 3)
	fmt.Printf("type of b: %T, slice %v\n", b, b)

	var c []int
	//c[0] = 1  //不能在开辟空间前赋值
	if c == nil {
		fmt.Println("slice是空")
	} else {
		fmt.Println("slice不是空")
	}

	c = make([]int, 3) //开辟空间
	c[0] = 1
	fmt.Println("c = ", c)

	// --- append ---
	d := make([]int, 3, 5)
	fmt.Printf("length of d : %d, cap of d: %d, slice d = %v\n", len(d), cap(d), d)

	d = []int{1, 2, 3}
	d = append(d, 4)
	fmt.Printf("length of d : %d, cap of d: %d, slice d = %v\n", len(d), cap(d), d)

	// -- intercept --
	e := d[0:2]
	fmt.Println("e = ", e)

	e[0] = 100
	fmt.Println("d = ", d, "e = ", e)

	f := make([]int, 3)
	copy(f, d) //将d中的值copy到f中
	f[0] = 200
	fmt.Println("d = ", d, "f = ", f)
}
