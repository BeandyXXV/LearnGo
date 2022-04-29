package basicSyntax

import (
	"fmt"
	"math"
)

// 函数首字母大写时，可以被包外的代码调用，但是函数首字母小写时，只能被包内的代码调用
func Var_learn() {
	// 当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil 等。

	/*
		标准声明方式(var 变量名 类型 = 值)
		var ：声明变量的关键字
		变量名：变量名必须是大小写字母、数字、下划线组成，且不能以数字开头
		类型：变量的类型，可以是任何类型，包括数组、结构体、接口等
	*/
	var firstVar int = 100                            // 声明一个变量，并赋值
	var secondVar, thirdVar string = "hello", "world" // 声明多个变量, 并赋值

	// 声明一个变量，不进行赋值
	var fourthVar int

	// 批量声明变量(不可进行赋值)
	var (
		fifthVar   int    = 100
		sixthVar   string = "hello"
		seventhVar struct {
			name string
			age  int
		}
	)

	// 简短赋值方式
	// 定义变量，同时显式初始化。
	// 不能提供数据类型。
	// 只能用在函数内部。
	// 使用时左侧的变量必须是未被声明的。
	eighthVar := "hello"
	ninthVar, tenthVar := "hello", "world" // 简短声明多个变量，并赋值

	// 编译器根据右值自动推断变量类型
	var attack = 40                // 右值为整形，编译器自动推断为int类型
	var defense = 20               // 右值为整形，编译器自动推断为int类型
	var damage_rate float32 = 0.17 // 如果不指定精度，编译器会默认使用最高精度（float64），当不许需要高精度时，可以指定较低精度节省空间。
	var damage = float32(attack-defense) * damage_rate

	// 多个短变量声明中，要求至少一个新声明的变量出现在左值中，此时编译器不会报错，重复声明的变量最终会采用最后赋的值。
	conn, err := 1, 2
	conn1, err := 2, 3

	// 多重赋值,变量的左值和右值按从左到右的顺序赋值。
	var num_1 int = 100
	var num_2 int = 200
	num_1, num_2 = num_2, num_1 // 交换变量的值

	// 匿名变量 _
	// 匿名变量是一个临时变量，只能在函数内部使用，不能用于函数外部。
	// 匿名变量不占用内存空间，也不会分配内存，匿名变量之间也不会因为多次声明而无法使用

	/*
		根据变量定义位置的不同，可以分为以下三个类型：
			函数内定义的变量称为局部变量
			函数外定义的变量称为全局变量
			函数定义中的变量称为形式参数
	*/

	// 打印上述变量的值
	println(firstVar, secondVar, thirdVar, fourthVar, fifthVar, sixthVar,
		seventhVar.name, seventhVar.age, eighthVar, ninthVar, tenthVar)
	println(damage)
	println(conn, err, conn1)
	println(num_1, num_2)
}

func Data_type() {
	// 数据类型

	/*
		// 整形变量
		// Go 语言提供了有符号和无符号两种整形变量，其中有符号整形变量可以表示负数，无符号整形变量不能表示负数。
		// 无符号整型变量：int8，int16，int32，int64
		// 无符号整型变量：uint8，uint16，uint32，uint64
		// 分别对应：8bit，16bit，32bit，64bit（二进制位）大小；
		// 此外还有两种整数类型：int，uint，随着计算机硬件与编译器的不同，这两种整数类型的大小也会不同（32bit / 64bit）。
		// 还有一种整数类型 uintptr，用于指针类型，只有在底层编程时才需要。
		// 尽管在特定在特定运行环境下 int、uint、uintptr 的大小可能相等，但依然是不同的类型，必须显示的对数据类型进行转换。
	*/

	/*
		// 浮点型变量
		// Go 语言提供了两种浮点型变量：float32 和 float64。
		// float32 和 float64 的大小与机器的实现有关，在 32 位机器上，float32 的大小为 32 位，而 float64 的大小为 64 位。
		// 在 Go 语言中，浮点型变量的精度与机器的实现有关，在 32 位机器上，float32 的精度为 6 位小数，而 float64 的精度为 15 位小数。
		// 通常应该优先使用 float64 类型，因为它的精度更高，因为 float32 类型的累计计算误差很容易扩散
	*/
	// 浮点数的生命可以只写整数部分或者小数部分
	var a float32 = .23
	var b float64 = 1.

	// 当生命很大或者很小的数时，可以使用科学计数法表示，通过 e 或 E 来指定指数部分
	const Avogadro = 6.022140857e23
	const Planck = 6.626070040e-34

	// 使用 printf 函数对浮点数进行打印的时候，可以使用 %f 或 %g 格式，%f 格式会输出最精确的结果，%g 格式会输出最简短的结果
	fmt.Printf("%f, %.2f, %g\n", math.Pi, math.Pi, math.Pi)

	println(a, b)
	println(Avogadro, Planck)
	println(math.Pi)

	/*

	 */

}
