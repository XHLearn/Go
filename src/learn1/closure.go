package learn1

import (
	"fmt"
	"log"
	"runtime"
)

func init() {
	fmt.Println("closure.init")
}

func MyClosure1() {
	p2 := Add2()
	fmt.Printf("Add2()(3)=%v\n", p2(3))
	p3 := Adder(2)
	fmt.Printf("Adder(2)(3)=%v\n", p3(3))
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func MyClosure2() {
	f := Adder_2()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300), " - ")
	fmt.Println()
}

func Adder_2() func(int) int {
	var x int // 闭包数据
	return func(delta int) int {
		x += delta
		return x
	}
}

// ClosureDebug:使用闭包调试
func ClosureDebug() {
	where := func() {
		a, file, line, d := runtime.Caller(1)
		log.Printf("%v %v %v %v\n", a, file, line, d)
	}
	where()
	where()
}
