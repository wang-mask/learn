# Go 语言教程

## 结构
    Go 语言的基础组成有以下几个部分：
    package main

    import "fmt"

    func main() {
      /* 这是我的第一个简单的程序 */
      fmt.Println("Hello, World!")
    }
1. 包声明,你必须在源文件中非注释的第一行指明这个文件属于哪个包,package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
2. 引入包
3. 函数,main 函数是每一个可执行程序所必须包含的
4. 变量
5. 语句 & 表达式
6. 注释
#
    当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。
    !!需要注意的是 { 不能单独放在一行

## Go 语言基础语法
    在 Go 程序中，一行代表一个语句结束。每个语句不需要像 C 家族中的其它语言一样以分号 ; 结尾，因为这些工作都将由 Go 编译器自动完成。如果你打算将多个语句写在同一行，它们则必须使用 ; 

    fmt.Println("Google" + "Runoob") //  字符串连接

    格式化输出
    // %d 表示整型数字，%s 表示字符串
    var stockcode=123
    var enddate="2020-12-31"
    var url="Code=%d&endDate=%s"
    var target_url=fmt.Sprintf(url,stockcode,enddate)
    fmt.Println(target_url)

## 数据类型
    1. 布尔 bool（true，false）
    2. 数字类型 int、 float32、 float64
    3. 字符串 Go 的字符串是由单个字节连接起来的 string

## 变量
    变量声明方式：
    1. var identifier type、 var a string = "Runoob"
    2. var v_name = value  // 声明变量时也可以不带类型，自动判断
    3. v_name := value 
    intVal := 1 相等于：var intVal int  intVal =1 

    vname1, vname2, vname3 = v1, v2, v3

    你可以通过 &i 来获取变量 i 的内存地址
    *a; 是一个指针变量
    空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。
## 常量
    const identifier [type] = value

## 语句
    if 布尔表达式 {
      /* 在布尔表达式为 true 时执行 */
    }

    for 布尔表达式 {

    }
    for i := 0; i <= 10; i++ {
        sum += i
    }
    // 数组遍历
    strings := []string{"google", "runoob"}
    for i, s := range strings {
        fmt.Println(i, s)
    }

## 函数
    func function_name( [parameter list] ) [return_types] {
      函数体
    }
    func max(num1, num2 int) int 

    go语言函数也可以赋值给变量

## 数组
    var variable_name [SIZE] variable_type
    初始化：
    var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
    balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}   //自动计算数组长度

## 指针
    var var_name *var-type
    &取地址
    *取数
    nil空
    var ptr [MAX]*int;


## 结构体
    type struct_variable_type struct {
      member definition
      member definition
      ...
      member definition
    }
    type Books struct {
      title string
      author string
      subject string
      book_id int
    }
    var Book1 Books 
    // 创建一个新的结构体
    fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

    // 也可以使用 key => value 格式
    fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

    访问成员：
    结构体.成员名
    结构体指针.成员名

## 切片 
    Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
    var identifier []type
    var slice1 []type = make([]type, len)

    s := arr[:]    // 初始化切片 s，是数组 arr 的引用。

    len(x) // 获取切片长度

    numbers[1:4]  // 截取 左闭右开

    append(numbers, num)  // 增加元素


## Range 
    range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。

    for _, num := range nums {
        sum += num
    }