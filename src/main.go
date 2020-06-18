package main

import (
	"./learn1"
	"fmt"
)

func init() {
	fmt.Println("main.init")
}

func Test() {
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
	}
	fmt.Println(items)
}

func main() {
	fmt.Println("-------begin main---------")
	// learn1.MyPrint()
	// learn1.GoOs()
	// learn1.Complex()
	// learn1.MyRand()
	// learn1.MyChar()
	// learn1.MyStrconv()
	// learn1.MyTime()
	// learn1.MyAddr()
	// learn1.MyStr()
	// learn1.MyDefer1()
	// learn1.MyDefer3("Go")
// Test()
	// learn1.MyMap()
	fmt.Println("----------------")
	learn1.MyRegexp()
	fmt.Println("-------end main---------")
}// learn1.MyClosure1()
// Test()
	// learn1.MyMap()
	fmt.Println("----------------")
	learn1.MyRegexp()
	fmt.Println("-------end main---------")
}// learn1.MyClosure2()
// Test()
	// learn1.MyMap()
	fmt.Println("----------------")
	learn1.MyRegexp()
	fmt.Println("-------end main---------")
}// learn1.ClosureDebug()
// Test()
	// learn1.MyMap()
	fmt.Println("----------------")
	learn1.MyRegexp()
	fmt.Println("-------end main---------")
}// learn1.Array1()
	// learn1.Array2()
	// Test()
	// learn1.MyMap()
	fmt.Println("----------------")
	learn1.MyRegexp()
	fmt.Println("-------end main---------")
}
