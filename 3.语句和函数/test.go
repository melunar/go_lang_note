package main

import "fmt"
import "strconv"
import "os"

// 函数基本结构
/*
关键字 func
参数和返回类型可选
func function_name( [parameter list] ) [return_types] {
   函数体
}
*/
func scoreSearch() {
	var score int = 0
	var fenshu string = "A"
	fmt.Printf("欢迎进入成绩查询系统\n")
	ZHU:
	// for true {
			var xuanzhe int = 0
			fmt.Println("1.进入成绩查询 2.退出程序")
			fmt.Printf("请输入序号选择:")
			fmt.Scanln(&xuanzhe)
			fmt.Printf("\n")
			if xuanzhe == 1 {
						goto CHA
			}
			if xuanzhe == 2 {
					os.Exit(1)
			}

	// }

	CHA: for true {
			fmt.Printf("请输入一个学生的成绩:")
			fmt.Scanln(&score)

			switch {
					case score >= 90:fenshu = "A"
					break

					case score >= 80&&score < 90:fenshu = "B"
					break

					case score >= 60&&score < 80:fenshu = "C"
					break

					default: fenshu = "D"
			}

			//fmt.Println(fenshu)
				var c string  = strconv.Itoa(score) // 转成字符
			switch {
					case fenshu == "A":
							fmt.Printf("你考了%s分,评价为%s,成绩优秀\n",c,fenshu)
							break
					case fenshu == "B" || fenshu == "C":
							fmt.Printf("你考了%s分,评价为%s,成绩良好\n",c,fenshu)
							break
					case fenshu == "D":
							fmt.Printf("你考了%s分,评价为%s,成绩不合格\n",c,fenshu)
							break
			}
			fmt.Printf("\n")
			goto ZHU
		}
}

// range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对
func rangeFoo1() {
	var nums = []int {2,3,4}
	var sum = 0
	for i, num := range nums {
		fmt.Printf("nums[%d] = %d\n", i, num)
		sum += num
	}
	fmt.Println("sum, ", sum)

	var kvs = map[string]string {"a": "a1", "b": "bb"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	
}

// 函数支持返回多个参数
func swap(x int, y string) (string, int) {
	return y, x
}

// 支持匿名函数
func getSequence() func() int { // 返回一个 【func() int】类型的函数
	i := 0
	// 没有函数名的匿名函数
	return func() int {
		 i += 1
		return i
	}
}

// Go 语言最少有个 main() 函数
func main() {
	//	if 语句	if 语句 由一个布尔表达式后紧跟一个或多个语句组成。
	// if...else 语句	if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。
	// if 嵌套语句	你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。
	// switch 语句	switch 语句用于基于不同条件执行不同动作。
	// select 语句	select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
	// Go 没有三目运算符，所以不支持 ?: 形式的条件判断
	fmt.Println("条件语句")
	var res bool
	if 10 > 9 { res = false } else { res = true }
	fmt.Println(res)
	fmt.Println("")


	fmt.Println("循环语句")
	// scoreSearch()
	/*
	for 循环	重复执行语句块
	循环嵌套	在 for 循环中嵌套一个或多个 for 循环
	break 语句	经常用于中断当前 for 循环或跳出 switch 语句
	continue 语句	跳过当前循环的剩余语句，然后继续进行下一轮循环。
	goto 语句	将控制转移到被标记的语句。
	*/
	fmt.Println("")


	fmt.Println("函数")
	// scoreSearch
	var x1 = 10
	var x2 = "go"
	X1, X2 := swap(x1, x2)
	fmt.Println(X1, X2)
	fmt.Println("")
	fmt.Println("范围 range")
	rangeFoo1()
	fmt.Println("")
}