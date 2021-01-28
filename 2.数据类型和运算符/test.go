package main

import "fmt"
import "unsafe"

// 这种因式分解关键字的写法一般用于声明全局变量
var (
	A int
	B bool
)

func main() {
	// 布尔型的值只可以是常量 true 或者 false
	fmt.Println("布尔值")
	var b bool = true
	var b1 = true // 根据值自行判定变量类型
	var b2 bool // 默认为false
	fmt.Println(b, b1, b2)
	fmt.Println("")

	// 数字类型 整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。
	/*
	1	uint8
	无符号 8 位整型 (0 到 255)
	2	uint16
	无符号 16 位整型 (0 到 65535)
	3	uint32
	无符号 32 位整型 (0 到 4294967295)
	4	uint64
	无符号 64 位整型 (0 到 18446744073709551615)
	5	int8
	有符号 8 位整型 (-128 到 127)
	6	int16
	有符号 16 位整型 (-32768 到 32767)
	7	int32
	有符号 32 位整型 (-2147483648 到 2147483647)
	8	int64
	有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
	===============
	1	float32
	IEEE-754 32位浮点型数
	2	float64
	IEEE-754 64位浮点型数
	3	complex64
	32 位实数和虚数
	4	complex128
	64 位实数和虚数
	===============
	1	byte
	类似 uint8
	2	rune
	类似 int32
	3	uint
	32 或 64 位
	4	int
	与 uint 一样大小
	5	uintptr
	无符号整型，用于存放一个指针
	*/
	fmt.Println("数字类型")
	// var i1, i4 int // 如果没有初始化赋值 值默认为0
	// i1, i4 = 10, 20 // 多变量可以在同一行进行赋值 类似js的解构赋值
	var i1, i4 int = 10, -40 // 声明变量的一般形式是使用 var 关键字, 可以一次声明多个变量
	// i1 = 10
	// i4 = -40
	var i2 uint8 = 100
	var i3 int8 = -10
	fmt.Println(i1, i2, i3, i4)
	var f1 float32 = 10.1
	var f2 float64 = 10131.1121
	var f3 complex64 = 10131.1121
	fmt.Println(f1, f2, f3)
	var by1 byte = 12
	fmt.Println(by1)
	fmt.Println("")
	
	// 字符串类型 字符串就是一串固定长度的字符连接起来的字符序列
	fmt.Println("字符串类型")
	var s string // 默认为 ""
	var s1 string = "sss"
	fmt.Println(s, s1)
	fmt.Println("")



	fmt.Println("默认值为 nil")
	var an1 *int
	var an2 []int
	var an3 map[string] int
	fmt.Println(an1, an2, an3)
	fmt.Println("")



	fmt.Println(":= 初始化声明")
	// 这种不带声明格式的只能在函数体(如main）中出现
	v_name := "vName"
	var vv_name string = "go"
	// v_name := "vName" // error 必须是未定义变量
	// vv_name := "vName" // error 必须是未定义变量
	fmt.Println(v_name, vv_name)
	fmt.Println("")


	fmt.Println("因式分解关键字")
	// var (  // 这种因式分解关键字的写法一般用于声明全局变量
  //   A int
  //   B bool
	// )
	
	fmt.Println(A, B)
	
	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("值类型和引用类型")
	fmt.Println("值类型")
	
	// 所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值
	// 当使用等号 = 将一个变量的值赋值给另一个变量时，如：j = i，实际上是在内存中将 i 的值进行了拷贝
	// 值类型的变量的值存储在栈中

	fmt.Println("通过 &i 来获取变量 i 的内存地址，例如：0xf840000040（每次的地址都可能不一样）。值类型的变量的值存储在栈中")
	fmt.Println(&A, &v_name)


	fmt.Println("")
	fmt.Println("引用类型")
	// 更复杂的数据通常会需要使用多个字，这些数据一般使用引用类型保存。
	// 一个引用类型的变量 r1 存储的是 r1 的值所在的内存地址（数字），或内存地址中第一个字所在的位置。
	/*

	这个内存地址称之为指针，这个指针实际上也被存在另外的某一个值中。

	同一个引用类型的指针指向的多个字可以是在连续的内存地址中（内存布局是连续的），这也是计算效率最高的一种存储形式；也可以将这些字分散存放在内存中，每个字都指示了下一个字所在的内存地址。

	当使用赋值语句 r2 = r1 时，只有引用（地址）被复制。

	如果 r1 的值被改变了，那么这个值的所有引用都会指向被修改后的内容
	*/
	
	fmt.Println("")

	fmt.Println("不同类型的值不能相互赋值")
	var ab = 1
	var ba = false
	// ab = ba // error 不同类型的值不能相互赋值
	fmt.Println(ab, ba)
	fmt.Println("")


	fmt.Println("常量")
	const LENGTH int = 10
	const WIDTH int = 5
	// LENGTH = 100 // error
	var area int = LENGTH * WIDTH
	// 常量用作枚举
	// todo 常量表达式中，函数必须是内置函数，否则编译不过
	const (
    A_5 = "abc"
    B_5 = len(A_5) // len 长度内置函数
    C_5 = unsafe.Sizeof(A_5)
	)
	fmt.Println(area, A_5, B_5, C_5)
	fmt.Println("")


	fmt.Println("iota")
	/*
		iota，特殊常量，可以认为是一个可以被编译器修改的常量。
		iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
		iota 可以被用作枚举值
	*/
	const (
		iota_a = iota   //0
		iota_b = iota   //1 等价于: iota_b
		iota_c          //2
		iota_d = "ha"   //独立值，iota += 1
		iota_e          //"ha"   iota += 1
		iota_f = 100    //iota +=1
		iota_g          //100  iota +=1
		iota_h = iota   //7,恢复计数
		iota_i          //8
	)
	fmt.Println(iota_a, iota_b, iota_c, iota_d, iota_e, iota_f, iota_g, iota_h, iota_i)
	fmt.Println("")



	fmt.Println("运算符")
	// 同于大部分预研的算数运算符：+ - * / % ++ -- (注意只能用作a++，a--, ++和--不能用作前置--a)
	var y_1 = 10
	var y_2 = false
	var y_4 = 11
	// 关系运算符 == >= <= != > <
	// var y_3 bool = y_1 == y_2 非同类型变量 无法比较
	var y_5 bool = y_1 == y_4
	// 逻辑运算符 && || ! 只能对条件表达式操作
	// var y_6 = !y_1 // error
	var y_7 = !y_2
	var y_8 = !(10 > 1)
	fmt.Println(y_1, y_2, y_5, y_7, y_8)
	fmt.Println()
}

// v_name1 := 100 // := 不带声明格式的声明必须是函数内部变量

// 注意：
// 1、如果你声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误【 xxx declared and not used 】
// 2、全局变量是允许声明但不使用的
