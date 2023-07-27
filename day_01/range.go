package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	// range第一个值是索引，第二个值是value
	for i, num := range nums {
		sum += num
		if num == 2 {
			fmt.Println("index", i, "num", num)
		}
	}
	fmt.Println(sum)

	m := map[string]string{"a": "A", "b": "B"}
	// range第一个值是key，第二个值是value
	for k, v := range m {
		fmt.Println(k, v)

	}
	for k := range m {
		fmt.Println("key ", k)

	}
}
