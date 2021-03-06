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
	fmt.Println(firstVar, secondVar, thirdVar, fourthVar, fifthVar, sixthVar,
		seventhVar.name, seventhVar.age, eighthVar, ninthVar, tenthVar)
	fmt.Println(damage)
	fmt.Println(conn, err, conn1)
	fmt.Println(num_1, num_2)
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

	fmt.Println("浮点数部分：")
	fmt.Println(a, b)
	fmt.Println(Avogadro, Planck)
	fmt.Println(math.Pi)

	/*
		// 复数变量
		// 计算机中，复数变量是由实部和虚部组成的，实部和虚部都是浮点型变量。
		// Go 语言中复数类型有两种：complex64（32 位实部和虚部） 和 complex128（64 位实部和虚部）。
		// complex64 和 complex128 的大小与机器的实现有关，在 32 位机器上，complex64 的大小为 16 位，而 complex128 的大小为 32 位。
		// complex128 是复数的默认类型。
	*/
	// 复数的声明（c 为复数变量名， complex64 为复数的类型，x，y 为构成该复数的两个 float 类型的数值，x 为实部，y 为虚部）
	x := 1.22
	y := 2.33
	var comp complex128 = complex(x, y)
	// 也可以简易声明
	comp_easy := complex(x, y)
	// 复数之间可以使用 == 判断相等，也可以使用 != 判断不相等，当且仅当实部和虚部都相等时，两个复数才相等。

	fmt.Println("复数部分：")
	fmt.Println(comp, comp_easy)

	// 对于一个复数，可以使用 real 和 imag 方法来获取它的实部和虚部
	fmt.Println("实部：", real(comp), "\t", "虚部：", imag(comp))

	/*
		// 布尔型变量
		// Go 语言中，布尔型变量的类型为 bool，它的值只能是 true 或者 false。
		// if 和 for 循环中的条件语句，只能使用布尔型变量。同时 == 和 != 等比较操作符也会返回布尔型变量。
		// 一元操作符 ! 可以将布尔型变量转换为相反的值。
	*/
	fmt.Println("布尔型变量：")
	fmt.Println((true == false) == false)

	// 布尔值可以用 &&（AND） 和 ||（OR） 操作符来组合，如果运算符左侧的值已经可以确定整个布尔表达式的值，那么右侧的运算符就不会被执行。
	// 下例中，s[0] 操作如果用于空字符串，会导致 panic 异常。但运算符左侧已经确定了整个表达式的值，右侧运算符不执行，所以不会报错。
	s := ""
	fmt.Println(s != "" && s[0] == 'x')

	// 因为 && 的优先级高于 ||（&& 对应逻辑乘法，|| 对应逻辑加法），可以视优先级来决定是否加上括号进行比较。
	c := "b"
	fmt.Println("a" <= c && c <= "z" || "A" <= c && c <= "Z" || "0" <= c && c <= "9")

	// Go 语言中不允许将整形变量强制转换为布尔形变量。
	// var n bool
	// fmt.Println(int(n) * 2)

	/*
		// 字符串变量
		// 一个字符串是一个不可改变的 UTF-8 字符序列。
		// 字符串是一种值类型，且值不可变，即创建某个文本后将无法再次修改这个文本的内容，更深入地讲，字符串是字节的定长数组。
	*/

	// 可以使用双引号""来定义字符串，字符串中可以使用转义字符来实现换行、缩进等效果，常用的转义字符包括：
	// \n 换行 \r 回车 \t 水平制表符 \v 垂直制表符 \\ 反斜杠 \' 单引号 \" 双引号 \u 字符的 Unicode 编码
	fmt.Println("字符串变量：")
	var str = "C语言中文网\nGo语言教程"
	fmt.Println(str)

	// 比较运算符是通过在内存中按字节比较来实现字符串的比较，所以比较的结果是字符串自然编码的顺序。
	var str_01 string = "str_test_01"
	var str_02 string = "str_test_02"
	fmt.Println(str_01 == str_02, len(str_01), len(str_02), len(str))

	// 字符串的内容（纯字节）可以通过标准索引法来获取，在方括号 [] 中写入索引，从 0 开始计数。此方法只适用于纯ASCII字符串。
	fmt.Println(str_01[0])
	// 获取字符串中某个字节的地址属于非法行为，例如 &srt[0]，因为字符串是一个纯字节数组，不能获取字符串中的字节地址。

	// 两个字符串可以通过 + 操作符进行拼接，拼接的结果是一个新的字符串，而不会改变原来的字符串。
	str_03 := "Beginning of the string " + "second part of the string"
	str_04 := "hel" + "lo,"
	str_04 += "world!"
	fmt.Println(str_03, str_04)

	// 在Go语言中，使用双引号书写字符串的方式是字符串常见表达方式之一，被称为字符串字面量（string literal），这种双引号字面量不能跨行.
	// 如果想要在源码中嵌入一个多行字符串时，就必须使用 ` 反引号
	// 在这种方式下，反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。
	// 在 ` 间的所有代码均不会被编译器识别，而只是作为字符串的一部分。
	const str_05 = `第一行
第二行
第三行
\r\n
`
	fmt.Print(str_05)

	/*
		// Go 语言字符类型
		// Go 语言的字符有以下两种：uint8 和 rune。
		// uint8 也可以称作 byte 型，可以表示 ASCII 码的一个字符
		// rune 代表一个 UTF-8 字符，当需要处理中文、日文等复合字符时，就需要用到 rune 类型，等价于 int32。
	*/
	fmt.Println("字符类型变量")
	var ch_byte byte = 65
	var ch_byte_01 int8 = 'a'
	fmt.Print(ch_byte, ch_byte_01)

	// Go 语言同样支持 Unicode 编码，因此字符同样称为 Unicode 代码点或者 runes，并在内存中用 int 表示。
	// 在文档中，一般使用格式 U+hhhh 表示，其中 h 表示一个 16 进制数。
	// 在书写 Unicode 字符时，需要在 16 进制前面加上前缀 \u 或者 \U。
	// 因为 Unicode 至少占用 2 个字节，所以我们使用 int16 或者 int 类型来表示。
	// 如果需要使用到 4 字节，则使用 \u 前缀，如果需要使用到 8 个字节，则使用 \U 前缀。
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U\n", ch, ch2, ch3) // UTF-8 code point

	// Unicode 包中内置了一些用于测试字符的函数
	// 判断是否为字母：unicode.IsLetter(ch)
	// 判断是否为数字：unicode.IsDigit(ch)
	// 判断是否为空白符号：unicode.IsSpace(ch)

	/*
		// 数据类型的转换
	*/
	// 在必要及可行的情况下，一个类型的值可以被转换成另一个类型。Go 语言不存在隐式类型转换，所有类型转换必须显式声明
	a_float := 10.0
	a_int := int(a_float)
	fmt.Println(a_int)
	// 类型转换只能在定义正确的情况下转换成功，当从一个取值范围较大的类型转换到一个取值范围较小的类型的时候，会出现精度丢失的问题。
	// 只有相同底层类型的变量之间可以进行相互转换，如 int16 -> int32。不同底层类型的变量相互转换时会发生编译错误。如 bool -> int
	// 浮点数在转换为整数时，会去掉所有小数部分只保留整数部分。
}
