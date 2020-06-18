package learn1

import fm "fmt" // 包的别名

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func init() {
	fmt.Println("simple.init")
}

// MyPrint 的函数注释，这里注释需要以函数名为开头
func MyPrint() {
	str := "Καλημέρα κόσμε; or こんにちは 世界"
	fm.Println(str, len(str))
}

func GoOs() {
	goos := runtime.GOOS
	fmt.Printf("当前操作系统是:%s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("当前Path是:%s\n", path)
}

func Complex() {
	c1 := 5 + 10i
	re := real(c1)
	im := imag(c1)
	fmt.Printf("复数为%v,实数为%g,虚数为%g\n", c1, re, im)
}

func MyRand() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d ", a)
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		r := rand.Intn(8)
		fmt.Printf("%d ", r)
	}
	fmt.Println()
	// 这里会发现如果没有用随机种子，上面的每次运行结果都是一样的
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f ", 100*rand.Float32())
	}
}

func MyChar() {
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
	fmt.Println()
}

func MyStrconv() {
	oriStr := "666"
	var an int
	fmt.Printf("int size:%d\n", strconv.IntSize)
	an, _ = strconv.Atoi(oriStr)
	fmt.Printf("对应的整数为:%d\n", an)
	an += 6
	newStr := strconv.Itoa(an)
	fmt.Printf("新的字符串为:%s\n", newStr)
}

func MyTime() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%4d-%02d-%02d\n", t.Year(), t.Month(), t.Day())

	tu := t.UTC()
	fmt.Println(tu)
	fmt.Printf("%f\n", 1e9)
	week_from_now := t.Add(60 * 60 * 24 * 7 * 1e9) //Add的单位为纳秒
	fmt.Println(week_from_now)
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2008 16:02"))
}

func MyAddr() {
	var i1 = 5
	fmt.Printf("整数 %d, 内存地址为: %p\n", i1, &i1)
	var intP *int
	intP = &i1
	fmt.Printf("内存地址为 %p 的值是 %d\n", intP, *intP)

	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints same string
}

func MyStr() {
	str := "Go language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}
	fmt.Println()

	str2 := "Chinese:こんにちは"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
	fmt.Println()
	fmt.Println("index int(rune) runechar bytes")
	for index, rune := range str2 {
		fmt.Printf("%-2d %d %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}
}

func MyDefer1() {
	fmt.Println("in MyDefer1 begin")
	defer MyDefer2()
	for i := 0; i < 5; i++ {
		defer fmt.Println("in MyDefer1", i)
	}
	fmt.Println("in MyDefer1 end")
}

func MyDefer2() {
	fmt.Println("in MyDefer2")
}

func MyDefer3(s string) (n int, err error) {
	defer func() {
		log.Printf("func(%q) = %d, %v\n", s, n, err)
		fmt.Printf("func(%q) = %d, %v\n", s, n, err)
	}()
	return 7, io.EOF
}
