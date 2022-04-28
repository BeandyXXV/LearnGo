package basicSyntax

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
		fifthVar   int
		sixthVar   string
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

	// 打印上述变量的值
	println(firstVar, secondVar, thirdVar, fourthVar, fifthVar, sixthVar,
		seventhVar.name, seventhVar.age, eighthVar, ninthVar, tenthVar, '\n')
	println(damage, '\n')
	println(conn, err, conn1)
}
