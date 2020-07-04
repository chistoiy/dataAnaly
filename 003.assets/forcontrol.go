package main

import (
	"fmt"
)

func main() {
	//forif()
	// forfor()
	//forswitch()
	// forgoto()
	forcalculateflag()
}

func forif() {
	age := 19
	if age > 18 {
		fmt.Printf("lucky ,Ur age is %d", age)
	} else if age > 15 {
		fmt.Print("U can listen to the music")
	} else {
		fmt.Println("sorry, Ur so younger")
	}
}

func forfor() {
	// 格式1
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		if i == 1 {
			continue
		}
		if i == 2 {
			break
		}
	}
	fmt.Println("==============")
	//格式2,初始参数可以不定义在for中
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	//格式3，变更省略
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//格式4 省略，死循环
	// for {
	// 	fmt.("for run !")
	// }

	// for range 对数字，字符串切片等适用
	for _, v := range "bei京" {
		fmt.Printf("%c\n", v)
	}

}

func forswitch() {
	n := 5
	switch n {
	case 1, 2:
		fmt.Println("get one")
	case 5:
		fmt.Println("get five")
	default:
		fmt.Println("get nothing")
	}

	switch { // 加逻辑判断
	case n > 3 && n < 15:
		fmt.Println("n is get in 3-4")
	}
	t := 1
	switch t {
	/*fallthrough,如果一个case中又fallthrough时，
	此case运行后，还会继续运行下一个case，为兼容c语言设计
	1.其下为default，那么就运行default，
	2.当其为最后一个时，会报错
	*/
	case 1:
		fmt.Println(n)
		fallthrough
	case 2:
		fmt.Println("w")
	default:
		fmt.Println("nothin")

	}

}

func forgoto() {
	p := 12
	if p > 10 {
		goto gototag
	}

	for i := 0; i < p; i++ {
		for j := 0; j < 10; j++ {
			if i > 3 && j == 5 {
				goto gototag2
			}
		}
		fmt.Println(i)
	}
gototag:
	fmt.Println("do something")
gototag2:
	fmt.Println("wer")
}

func forcalculateflag() { //算数运算符
}
