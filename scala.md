# Scala 教程
## Scala 是一门多范式（multi-paradigm）的编程语言，设计初衷是要集成面向对象编程和函数式编程的各种特性。Scala 运行在 Java 虚拟机上，并兼容现有的 Java 程序。Scala 与 Java 的最大区别是：Scala 语句末尾的分号 ; 是可选的。

## 可以认为 Scala 程序是对象的集合，通过调用彼此的方法来实现消息传递

    类名 - 对于所有的类名的第一个字母要大写。如果需要使用几个单词来构成一个类的名称，每个单词的第一个字母要大写。

    方法名称 - 所有的方法名称的第一个字母用小写。如果若干单词被用于构成方法的名称，则每个单词的第一个字母应大写。

    程序文件名 - 程序文件的名称应该与对象名称完全匹配

## 数据类型
    scala没有java中的原生类型， 都是对象 Int等
    使用关键词 "var" 声明变量，使用关键词 "val" 声明常量
    var VariableName : DataType [=  Initial Value]
    var myVar : String = "Foo"
    var myVar = 10; 也可以不给出变量类型，但是此时必须给出初始值

## 访问修饰符
    用 private 关键字修饰，带有此标记的成员仅在包含了成员定义的类或对象内部可见
    （Protected）成员的访问比 java 更严格一些。因为它只允许保护成员在定义了该成员的的类的子类中被访问
    
    作用域保护
    private[x] 或 protected[x]
    这里的x指代某个所属的包、类或单例对象。如果写成private[x],读作"这个成员除了对[…]中的类或[…]中的包中的类及它们的伴生对像可见外，对其它所有类都是private。

## 方法与函数
    Scala 方法是类的一部分，而函数是一个对象可以赋值给一个变量。换句话来说在类中定义的函数即是方法。
    Scala 中使用 val 语句可以定义函数，def 语句定义方法。
    方法：
    def functionName ([参数列表]) : [return type]
    def functionName ([参数列表]) : [return type] = {
      function body
      return [expr]
    }
    
## 字符串
    在 Scala 中，字符串的类型实际上是 Java String，它本身没有 String 类。
    java 中 String是不可以修改的
    创建可修改的字符串 StringBuilder
    val buf = new StringBuilder;
    str.length()
    str.charAt()
    str.indexOf()
    str.replace()
    str.split()
    str.substring(int beginIndex, int endIndex)
    str.toCharArray()
    格式化输出：
    var fs = printf("浮点型变量为 " +
                   "%f, 整型变量为 %d, 字符串为 " +
                   " %s", floatVar, intVar, stringVar)
    println(fs)

## 数组
    Scala 语言中提供的数组是用来存储固定大小的同类型元素
    var z:Array[String] = new Array[String](3)
    var z = new Array[String](3)
    访问：z(n)！！！
    遍历：
    for ( x <- myList ) {
        println( x )
    }
    for ( i <- 0 to (myList.length - 1)) {
        total += myList(i);
    }

## 迭代器
    Scala Iterator（迭代器）不是一个集合，它是一种用于访问集合的方法。
    迭代器 it 的两个基本操作是 next 和 hasNext。
    val it = Iterator("Baidu", "Google", "Runoob", "Taobao")
    while (it.hasNext){
        println(it.next())
    }
    println("最大元素是：" + ita.max )
    println("最小元素是：" + itb.min )
    ita.size

## 类和对象
    类是对象的抽象，而对象是类的具体实例。类是抽象的，不占用内存，而对象是具体的，占用存储空间
    类似于java

## 模式匹配
    类似于switch
    def matchTest(x: Int): String = x match {
      case 1 => "one"
      case 2 => "two"
      case _ => "many"
   }

## I/O
    Scala 进行文件写操作，直接用的都是 java中 的 I/O 类 （java.io.File)
    写文件：
    val writer = new PrintWriter(new File("test.txt" ))
    writer.write("菜鸟教程")
    writer.close()

    读屏幕：
    val line = StdIn.readLine()

    读文件：
    Source.fromFile("test.txt" ).foreach{
        print
    }