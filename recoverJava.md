# 复习 JAVA

## 全局知识
1. 所有的 Java 程序由 public static void main(String[] args) 方法开始执行。
2. 源文件名必须和类名相同。当保存文件的时候，你应该使用类名作为文件名保存（切记 Java 是大小写敏感的），文件名的后缀为 .java。（如果文件名和类名不相同则会导致编译错误）。
3. 构造方法，构造方法名与类名一致，public classname{}
4. 一个源文件中只能有一个 public 类，源文件的名称应该和 public 类的类名保持一致
5. 先 package 语句，然后 import 语句，再类定义语句

## java 基本数据类型
    Java 的两大数据类型:
      内置数据类型
      引用数据类型
    Java语言提供了八种基本类型。六种数字类型（四个整数型，两个浮点型），一种字符类型，还有一种布尔型
    数字类型：byte八位、short十六位、int三十二位、long六十四位、float单精度32位浮点和double 数据类型是双精度、64 位
    布尔类型：boolean
    字符类型：char

    引用类型：引用类型的变量非常类似于C/C++的指针。引用类型指向一个对象，指向对象的变量是引用变量。这些变量在声明时被指定为一个特定的类型，比如 Employee、Puppy 等。变量一旦声明后，类型就不能被改变了。

    常量：使用 final 关键字来修饰常量

    类型转换：byte,short,char—> int —> long—> float —> double 
    运算中，不同类型的数据先转化为同一类型，然后进行运算。
    自动转换：必须满足转换前的数据类型的位数要低于转换后的数据类型
    强制转换：(type)变量名

## java 访问修饰符
    default (即默认，什么也不写）: 在同一包内可见，不使用任何修饰符。使用对象：类、接口、变量、方法。
    private : 在同一类内可见。使用对象：变量、方法。 注意：不能修饰类（外部类），只能通过类中公共的 getter 方法被外部类访问
    public : 对所有类可见。使用对象：类、接口、变量、方法
    protected : 对同一包内的类和所有子类可见。使用对象：变量、方法。 注意：不能修饰类（外部类）。

    父类中声明为 public 的方法在子类中也必须为 public。
    父类中声明为 protected 的方法在子类中要么声明为 protected，要么声明为 public，不能声明为 private。
    父类中声明为 private 的方法，不能够被继承
  
## 改进 for 循环
    for(声明语句 : 表达式)
    {
      //代码句子
    }
    for( String name : names ) {
      System.out.print( name );
      System.out.print(",");
    }

## 包装类
    在实际开发过程中，我们经常会遇到需要使用对象，而不是内置数据类型的情形。为了解决这个问题，Java 语言为每一个内置数据类型提供了对应的包装类。

    Java Number类
    所有的包装类（Integer、Long、Byte、Double、Float、Short）都是抽象类 Number 的子类。
    toString(),Integer.parseInt()

    Math 类
      Math.min(),max(),pow(x,n),sqrt(),random(),abs(),round(),ceil(),floor()

    Character 类
    Character.isLetter()，isDigit()，isUpperCase()，isLowerCase()，toUpperCase()，toLowerCase()，toString()

## String 类
    String一旦创建不可更改，更改只是变化指针
    String 创建的字符串存储在公共池中，而 new 创建的字符串对象在堆上
    str.length()
    str1.concat(str2); # 或者直接使用 +
    格式化字符串
    System.out.printf("浮点型变量的值为 " +
                  "%f, 整型变量的值为 " +
                  " %d, 字符串变量的值为 " +
                  "is %s", floatVar, intVar, stringVar);
    判断两个字符串内容是否相等尽量使用str1.equal(str2)
    indexOf(),返回匹配到的第一个的索引
    replace(oldstr, newstr) 替换操作，返回新串
    split()
    subString(int beginIndex, int endIndex) 获取子串
    toLowerCase(),toUpperCase()
    contains(CharSequence chars)
    isEmpty()
    toCharArray() 将字符串变成字符数组

## Java StringBuffer 
    当对字符串进行修改的时候，需要使用 StringBuffer 和 StringBuilder 类。在使用 StringBuffer 类时，每次都会对 StringBuffer 对象本身进行操作，而不是生成新的对象
    StringBuilder 类在 Java 5 中被提出，它和 StringBuffer 之间的最大不同在于 StringBuilder 的方法不是线程安全的（不能同步访问）。
    由于 StringBuilder 相较于 StringBuffer 有速度优势，所以多数情况下建议使用 StringBuilder 类。
    StringBuilder sb = new StringBuilder(10); 相当于声明了长度为10的buffer
    append(String s)
    reverse()
    delete(int start, int end) 前闭后开
    insert(int offset, String str)
    replace(int start, int end, String str)
    etCharAt(int index, char ch)
    其他方法大致跟String一致
  
## 数组
    声明：ArrayType[] name = new ArrayType[size],ArrayType[] name = {}
    属性：.length
    方法：
    Arrays.fill(arr, value)
    Arrays.sort(arr)
    Arrays.equals(arr1, arr2)
    数组作为参数
    public static void printArray(int[] array)

## 可变参数
    typeName... parameterName
    public static void printMax( double... numbers) {}
    访问 numbers 就是一个可变数组

## IO
    输入流表示从一个源读取数据，输出流表示向一个目标写数据。
    读取单个字符：
      // 使用 System.in 创建 BufferedReader
      BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
      // 读取字符
      do {
          c = (char) br.read();
      } while (c != 'q');

      读取字符串：
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        do {
            str = br.readLine();
        } while (!str.equals("end"));
      
      文件操作：
      InputStream f = new FileInputStream("C:/java/hello");
      File f = new File("C:/java/hello");
      InputStream in = new FileInputStream(f);

      OutputStream f = new FileOutputStream("C:/java/hello")
      File f = new File("C:/java/hello");
      OutputStream fOut = new FileOutputStream(f);