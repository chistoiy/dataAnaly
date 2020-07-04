package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// forint()
	// forfloat()
	// forcomplex()
	// forbool()
	// forfmt()
	// forstring()
	// forrange()
	// forbyte()
	modystring()
	countchinese()
}

func forint() {
	var a int = 10
	fmt.Printf("十进制打印：%d\n", a)
	fmt.Printf("转为二进制：%b\n", a)
	fmt.Printf("转为八进制：%o\n", a)
	fmt.Printf("转为十六进制：%x\n", a)

	i2 := 077
	fmt.Printf("八进制转换为十进制打印： %d \n", i2)

	i3 := 0X1011
	fmt.Printf("十六进制转换为十进制打印： %d \n", i3)

}

func forfloat() {
	//float
	//math.MaxFloat32
	f1 := 1.2345

	fmt.Printf("%T ", f1)
	var f0 float64 = 1.567
	fmt.Println(f0)
	fmt.Println(int(f0))

}

func forcomplex() {
	// complex
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1, c2)

}
func forbool() {
	//bool
	var (
		b1 bool
		b2 bool = true
	)
	fmt.Println(b1, b2)
}
func forfmt() {
	s := "asc"

	fmt.Println("打印时默认带换行", s)
	fmt.Print("只打印，无换行", s)

	fmt.Printf("带格式打印%s", s)
	/*%T,取参数类型，
		%d参数转换位十进制，
	%o转八进制
	%x转十六进制
	%b 转二进制
	%v 只取值
	%#v 打印时增加类型的描述符，如string，会增加引号
	%s string类型占位
	%c 字符
	*/

}

func forstring() {

	s1 := "字符串"
	s2 := 'a'
	s3 := `
	多行字符串，使用的esc之下的点，不是引号
	`
	fmt.Printf("%T,%v,%#v\n", s1, s1, s1)
	fmt.Printf("%T,%v,%#v\n", s2, s2, s2)
	fmt.Printf("%T,%v,%#v\n", s3, s3, s3)
	/*
		转义字符
		\r回车  \n换行 \t制表 tab
		\'	    \"	  \\
		常用操作
		len(str)
		+ fmt.sprintf(str1,str2) 均可拼接
		strings.Split(str1, str2),以str2来分割str1
			   .Contains(str1,str2)
			   .HasPrefix
			   .HasSuffix
			   .Index 返回字串的字符左起位置
			   .LastIndex 最后一个的位置
			   .Join()
	*/
	fmt.Println("------------------------")
	fmt.Println("拼接1", s1+s3)
	fmt.Println("拼接2", fmt.Sprint("拼接", s1, s3))
	fmt.Println("分割", strings.Split(s1, "符"))
	fmt.Println("包含", strings.Contains(s1, "符"))
	fmt.Println("前缀判断", strings.HasPrefix(s1, "字"))
	fmt.Println("后缀判断", strings.HasSuffix(s1, "串"))
	fmt.Println("查询的第一字串", strings.Index(s1, "符"))
	fmt.Println("查询的最后字串", strings.LastIndex(s1, "符"))
	fmt.Println("拼接字串", strings.Join(strings.Split(s1, "符"), "+"))
}

func forrange() {
	s := "test_word"

	for a, c := range s {
		fmt.Printf("%#v:%c\n", a, c)
	}
}

func forbyte() {
	s := "test_你好word"
	/*go 中无法转换位ascii码类型的字符，定义为rune类型，
	如汉字或其他语言特殊符号等
	其实际为int32类型
	*/
	for a, c := range s {
		fmt.Printf("%#v:%c,and type is %T\n", a, c, c)
	}
}

func modystring() {
	/*
		string本身不可修改，因此需要转换为rune数组后修改完毕再转换为字符串
	*/
	s2 := "qwe"
	s3 := []rune(s2)
	s3[0] = 'T'
	fmt.Println(string(s3))
}

func countchinese() {
	str := "hello,你好"
	var count int
	for _, w := range str {
		if unicode.Is(unicode.Han, w) {
			count++
		}
	}
	fmt.Println("count:", count)

}
