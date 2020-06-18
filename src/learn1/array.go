package learn1

import "fmt"

func f1(a [3]int) {
	a[1] = 1
	fmt.Printf("a的类型：%T\n", a)
}

func f2(a *[3]int) {
	a[2] = 2
	fmt.Printf("a的类型：%T\n", a)
}

func Array1() {
	var a [3]int
	f1(a)
	fmt.Println("f1(a)之后a的值", a)
	f2(&a)
	fmt.Println("f2(&a)之后a的值", a)
}

func Array2() {
	var arr1 [6]int
	// 注意 绝对不要用指针指向 slice。切片本身已经是一个引用类型,所以它本身就是一个指针!!
	var slice1 []int = arr1[2:5]
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice1[%d]=%d\n", i, slice1[i])
	}
	fmt.Printf("%d %d %d\n", len(arr1), len(slice1), cap(slice1))

	slice1 = slice1[:4] // 这里注意下
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice1[%d]=%d\n", i, slice1[i])
	}
	fmt.Printf("%d %d %d\n", len(arr1), len(slice1), cap(slice1))
}

func MyMap() {
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int)
		items[i][i] = i
	}
	fmt.Println(items)

	items2 := make([]map[int]int, 5)
	for i, item := range items2 {
		// 这里的item只是map值的拷贝，修改不会被应用到原数据
		// 可以看到两个的指针地址不一样
		fmt.Printf("%p %p\n", &items2[i], &item)
		item = make(map[int]int, 1)
		item[i] = i
	}
	fmt.Println(items2)
}