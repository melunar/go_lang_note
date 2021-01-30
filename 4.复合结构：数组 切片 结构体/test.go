package main

import "fmt"

// 数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整型、字符串或者自定义类型
// 数组的定义：var variable_name [SIZE] variable_type，如 var arr [10] int
// 多维数组 var arr [10][10] int

func arrFoo1() {
	balance := [3]float32{100, 101.1, 103.4}
	// 数组长度不确定 [...]表示
	var balanceNotSureLength = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	// 下标制定元素(未制定赋值则为默认值)
	balanceWithIndex := [5]float32{1: 2.0 ,3: 7.0}
	balanceWithIndex[0] = 9.1
	fmt.Println(balance, balanceNotSureLength, balanceWithIndex)
}

func arrFoo2() {
	var i int
	balance := [5]float32{1000, 2.1, 2.0, 100, 50.1}
	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}
}


// 一个指针变量指向了一个值的内存地址
// var var_name *var-type
// 在指针类型前面加上 * 号（前缀）来获取指针所指向的内容
// 取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
func pointerFoo1 () {
	var a int = 20
	var ip *int

	ip = &a
	fmt.Println(a, ip, &a, *ip)


	/*
		空指针:
		当一个指针被定义后没有分配到任何变量时，它的值为 nil
		nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
		一个指针变量通常缩写为 ptr
	*/
	var  ptr *int
	// 判断空指针 ptr == nil
	var isNilPointer bool = ptr == nil
	fmt.Println(ptr, isNilPointer)
}

// 结构体是由一系列具有相同类型或不同类型的数据构成的数据集合
/*
	type struct_variable_type struct {
		member definition
		...
	}
*/
// 结构体赋值：
// variable_name := structure_variable_type {value1, value2...valuen}
// variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}

func structFoo1 () {
	type Books struct {
		title string
		author string
		price float32
	}
	fmt.Println(Books{"Go 语言", "Jack", 32})
	// key - value 赋值方式 没有赋值的为默认值
	fmt.Println(Books{title: "Go 语言", author: "Jack"})

	var book1 Books // 不可以直接整体赋值 = {title: "Go 语言1", author: "Jack"}
	book1.title = "Go lang"
	var book2 = Books{title: "Go 语言", author: "Jack2"}
	fmt.Println(book1.title, book2.author)
}

// 切片是对数组的抽象
// Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大
func sliceFoo1 () {
	// 声明一个未指定大小的数组来定义切片： var identifier []type
	// make函数声明切片： make([]T, length, capacity) //  capacity 指定容量 为可选参数
	s := []int {1,2,3}

	arr1 := []int {1,2,3,4,5,6}
	s1 := arr1[:] // 引用一个数组的方式初始化切片

	// s := arr[startIndex:endIndex]
	// arr 中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片
	// 引用冒号缺少的一侧代表从该侧头切
	s2 := arr1[2:5]
	s3 := arr1[2:]

	var s4 = make([]int, 4)
	s4[1] = 100
	var s4L = len(s4) //  len() 方法获取长度

	
	fmt.Println(s, s1, s2, s3, s4, s4L)

	var numbers = make([]int, 4, 5) // Golang 不允许创建容量小于长度的切片  make([]int, 3, 2)
	numbers[2] = 2
	numbers[1] = 1
	numbers[3] = 3
	// numbers[4] = 4
	// numbers[6] = 6
	// 计算容量的方法 cap() 可以测量切片最长可以达到多少
	// 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数
	fmt.Printf("numbers len=%d cap=%d slice=%v\n",len(numbers),cap(numbers),numbers)

	// 空切片
	// var numbersEmpty []int
	// numbersEmpty == nil // true

	// 设置下限及上限来设置截取切片 [lower-bound:upper-bound]
	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)
	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:2] ==", numbers[1:2])
	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
func sliceFoo2()  {
	// append copy 函数
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2,3,4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1,numbers)
	printSlice(numbers1)  

}

func main() {
	fmt.Println("数组")
	arrFoo1()
	fmt.Println("")
	arrFoo2()
	fmt.Println("")
	fmt.Println("简单的指针")
	pointerFoo1()
	fmt.Println("")
	fmt.Println("结构体")
	structFoo1()
	fmt.Println("")
	fmt.Println("切片")
	sliceFoo1()
	fmt.Println("")
	sliceFoo2()
	fmt.Println("")
	fmt.Println("")
}