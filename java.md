
## Java 简介
    一个Java源码只能定义一个public类型的class，并且class名称和文件名要完全一致；
    使用javac可以将.java源码编译成.class字节码；
    使用java可以运行一个已编译的Java程序，参数是类名。

## 基本数据类型
    对于float类型，需要加上f后缀。
    char用'', string用""
    final double PI = 3.14; 定义常量
    var sb = new StringBuilder(); 可以省略变量类型，自动判断
    在运算过程中，如果参与运算的两个数类型不一致，那么计算结果为较大类型的整型。
    强转 (type)varName
    浮点数常常无法精确表示
    由于浮点数存在运算误差，所以比较两个浮点数是否相等常常会出现错误的结果。正确的比较方法是判断两个浮点数之差的绝对值是否小于一个很小的数：
    可以将浮点数强制转型为整数。在转型时，浮点数的小数部分会被丢掉。如果转型后超过了整型能表示的最大范围，将返回整型的最大值。
    java中char使用Unicode表示字符，两个字节
    字符串可以使用 + 连接
    多行字符串"""...."""，多行字符串前面共同的空格会被去掉
    java的字符串是不可变的，一旦定义就不可修改
    字符串没有赋值时为 null
    数组：
    int[] ns = new int[5];
    int[] ns = { 68, 79, 91, 85, 62 };
    ns.length

## 输入输出
    格式化输出
    System.out.printf("%.2f\n", d);
    输入：
    Scanner scanner = new Scanner(System.in); 
    String name = scanner.nextLine();
    scanner.nextInt()

## 流程控制
    java中引用类型使用“==”代表引用是否指向同一个对象，判断引用对象的内容是否相等，使用equals()方法 s1.equals(s2)
    switch语句还可以匹配字符串。字符串匹配时，是比较“内容相等”

## 数组操作
    for (int n : ns) {
        System.out.println(n);
    }
    快速打印数组Arrays.toString(ns)
    排序 Arrays.sort()
    二维数组： 每个数组元素的长度并不要求相同
    int[][] ns = {
        { 1, 2, 3, 4 },
        { 5, 6, 7, 8 },
        { 9, 10, 11, 12 }
    };

## 面向对象
    一个Java源文件可以包含多个类的定义，但只能定义一个public类，且public类名必须与文件名一致。如果要定义多个public类，必须拆到多个Java源文件中。
    可变参数：
    void setNames(String... names) // names为一个string数组
    g.setNames("Xiao Ming", "Xiao Hong", "Xiao Jun"); 
    外部代码通过public方法操作实例，内部代码可以调用private方法；
    没有在构造方法中初始化字段时，引用类型的字段默认是null，数值类型的字段用默认值，int类型默认值是0，布尔类型默认值是false

### 构造函数
    构造函数中调用构造函数 this(name, 18); 

### 继承
    子类无法访问父类的private字段或者private方法
    protected关键字可以把字段和方法的访问权限控制在继承树内部，一个protected字段和方法可以被其子类，以及子类的子类所访问
    super关键字表示父类（超类）
    任何class的构造方法，第一行语句必须是调用父类的构造方法。如果没有明确地调用父类的构造方法，编译器会帮我们自动加一句super();
    所以父类构造方法一般要有一个无参构造方法
    如果父类没有默认的构造方法，子类就必须显式调用super()并给出参数以便让编译器定位到父类的一个合适的构造方法。即子类不会继承任何父类的构造方法
    只要某个class没有final修饰符，那么任何类都可以从该class继承。
    向上转型：把一个子类类型安全地变为父类类型的赋值，及将子类实例赋值给父类
    向下转型：不可以！把一个父类类型强制转型为子类类型，不能把父类变为子类，因为子类功能比父类多，多的功能无法凭空变出来。
    组合：一个类中有另一个类的实例，相当于学生有书的关系

### 多态
    覆写父类方法要加 @Override
    Java的实例方法调用是基于运行时的实际类型的动态调用，而非变量的声明类型。
    多态是指，针对某个类型的方法调用，其真正执行的方法取决于运行时期实际类型的方法
    多态用处：
    public double totalTax(Income... incomes) {
        double total = 0;
        for (Income income: incomes) {
            total = total + income.getTax();
        }
        return total;
    }
    传入其中的参数为各种扩展之后的子类，他们有不同的getTax的实现方式
    在子类的覆写方法中，如果要调用父类的被覆写的方法，可以通过super来调用
    final修饰的方法可以阻止被覆写；
    final修饰的class可以阻止被继承；
    final修饰的field必须在创建对象时初始化，随后不可修改。

### 抽象类
    如果一个class定义了方法，但没有具体执行代码，这个方法就是抽象方法，抽象方法用abstract修饰。
    抽象类无法实例话
    定义了抽象类Person，以及具体的Student、Teacher子类的时候，我们可以通过抽象类Person类型去引用具体的子类的实例：

### 接口
    如果一个抽象类没有字段，所有方法全部都是抽象方法，就可以把该抽象类改写为接口：interface。
    使用implements实现
    一个类可以实现多个interface

### 静态方法
    因为静态方法属于class而不属于实例，因此，静态方法内部，无法访问this变量，也无法访问实例字段，它只能访问静态字段。

### 包
    包没有父子关系。java.util和java.util.zip是不同的包，两者没有任何继承关系。
    package_sample
    └─ src
        ├─ hong
        │  └─ Person.java
        │  ming
        │  └─ Person.java
        └─ mr
        └─ jun
            └─ Arrays.java
    包可以是多层结构，用.隔开。例如：java.util
    Java文件对应的目录层次要和包的层次一致（一个点就是一个目录）
    class查找顺序：
    如果是简单类名，按下面的顺序依次查找：
        查找当前package是否存在这个class；
        查找import的包是否包含这个class；
        查找java.lang包是否包含这个class。
    默认自动import java.lang.*

### 访问修饰符
    定义为public的class、interface可以被其他任何类访问：
    定义为public的field、method可以被其他类访问，前提是首先有访问class的权限：

    定义为private的field、method无法被其他类访问：只能在类内访问

    protected作用于继承关系。定义为protected的字段和方法可以被子类访问，以及子类的子类：

    包作用域是指一个类允许访问同一个package的没有public、private修饰的class，以及没有public、protected、private修饰的字段和方法。

### 内部类
    Outer outer = new Outer("Nested"); // 实例化一个Outer
    Outer.Inner inner = outer.new Inner(); // 实例化一个Inner
    内部类访问外部类实例 Outer.this.name
    内部类必须依赖于外部类的实例而存在

### classpath和jar
    classpath是JVM用到的一个环境变量，它用来指示JVM如何搜索class。
    jar包相当于目录
    如果有很多.class文件，散落在各层目录中，肯定不便于管理。如果能把目录打一个包，变成一个文件，就方便多了。
    jar包就是用来干这个事的，它可以把package组织的目录层级，以及各个目录下的所有文件（包括.class文件和其他文件）都打成一个jar文件

## 核心类
### String 
    实际上字符串在String内部是通过一个char[]数组表示的，Java字符串的一个重要特点就是字符串不可变。这种不可变性是通过内部的private final char[]字段，以及没有任何修改char[]的方法实现的。
    判断内容是否相等：equals()方法
    Java编译器在编译期，会自动把所有相同的字符串当作一个对象放入常量池（String s1 = "hello";String s2 = "hello";s1和s2指向同一地址）
    方法：
    .contains()  .indexOf()
    .substring()
    .replace(old,new)
    .split()
    .join()
    String.join(分隔符, string数组);

    转换：
    String.valueOf(anyType);  // 将任意类型转为string
    Integer.parseInt("123"); 
    Boolean.parseBoolean("true");
    char[] cs = "Hello".toCharArray(); // String -> char[] 
    String s = new String(cs); // char[] -> String
### StringBuilder
    StringBuilder，它是一个可变对象，可以预分配缓冲区，这样，往StringBuilder中新增字符时，不会创建新的临时对象
    StringBuilder sb = new StringBuilder(1024);
    String s = sb.toString();
    .append()
    .delete(start, end)
    .insert(index, str)
    .reverse()
    .charAt()
    .setCharAt()
    继承所有String的方法

### 包装类
    直接把int变为Integer的赋值写法，称为自动装箱（Auto Boxing），反过来，把Integer变为int的赋值写法，称为自动拆箱（Auto Unboxing）。
    所有的包装类型都是不变类
    对包装类判等，绝对不能用==比较，因为Integer是引用类型，必须使用equals()比较：
    Integer n = Integer.valueOf(100);
    我们把能创建“新”对象的静态方法称为静态工厂方法。Integer.valueOf()就是静态工厂方法，它尽可能地返回缓存的实例以节省内存。

### JavaBean
    若干private实例字段；
    通过public方法来读写实例字段。通过set、get变量名来进行赋值和获取

### 枚举类
    普通定义常量：public static final int SUN = 0;
    以上定义常量有一个严重的问题就是，编译器无法检查每个值的合理性，不能检查有些值是否被枚举

    enum Weekday {
        SUN, MON, TUE, WED, THU, FRI, SAT;
    }
    Weekday day = Weekday.SUN;
    枚举类好处：
    1、enum常量本身带有类型信息
    2、不可能引用到非枚举的值，因为无法通过编译。
    枚举类型无论是否为非引用类型都可以使用“ == ”判等，因为enum类型的每个常量在JVM中只有一个唯一实例，所以可以直接用==比较：
    方法：
    String s = Weekday.SUN.name(); // "SUN"
    每个枚举的值都是class实例

### 工具类
    Math.abs(-100);
    Math.max(100, 99); // 100
    Math.min(1.2, 2.3); // 1.2
    Math.pow(2, 10); Math.sqrt(2);
    Math.random(); // 生成一个随机数x，x的范围是0 <= x < 1：
    Random用来创建伪随机数。所谓伪随机数，是指只要给定一个初始的种子，产生的随机数序列是完全一样的。
    在创建Random实例时指定一个种子，就会得到完全确定的随机数序列
    Random r = new Random();
    r.nextInt(); // 2071575453,每次都不一样
    r.nextInt(10); // 5,生成一个[0,10)之间的int

## 异常
    异常是一种class，因此它本身带有类型信息。异常可以在任何地方抛出，但只需要在上层捕获，这样就和方法调用分离了
    Throwable有两个体系：Error和Exception，
    Error表示严重的错误，程序对此一般无能为力
    而Exception则是运行时的错误，它可以被捕获并处理。
    在方法定义的时候，使用throws Xxx表示该方法可能抛出的异常类型。调用方在调用的时候，必须强制捕获这些异常，否则编译器会报错。
    如果在调用处不捕获该处理的异常，则必须在该用处所在的方法上throws出该异常，即要么内部处理，要么抛出到外部
    RuntimeException无需强制捕获，非RuntimeException（Checked Exception）需强制捕获，或者用throws声明
    存在多个catch的时候，catch的顺序非常重要：子类必须写在前面。
    finally语句
    无论是否有异常发生，如果我们都希望执行一些语句
    处理过程相同的异常的捕获：catch (IOException | NumberFormatException e)
    当某个方法抛出了异常时，如果当前方法没有捕获异常，异常就会被抛到上层调用方法，直到遇到某个try ... catch被捕获为止：

### 抛出异常
    1、创建某个Exception的实例；
    2、用throw语句抛出。
    捕获到异常并再次抛出时，一定要留住原始异常，否则很难定位第一案发现场！
    在catch中抛出异常，不会影响finally的执行。JVM会先执行finally，然后抛出异常。
    这说明finally抛出异常后，原来在catch中准备抛出的异常就“消失”了，因为只能抛出一个异常。没有被抛出的异常称为“被屏蔽”的异常（Suppressed Exception）。

### 断言
    断言（Assertion）是一种调试程序的方式。在Java中，使用assert关键字来实现断言。
    double x = Math.abs(-123.45);
    assert x >= 0;
    System.out.println(x);
    语句assert x >= 0;即为断言，断言条件x >= 0预期为true。如果计算结果为false，则断言失败，抛出AssertionError。

### JDK Logging
    方便打印调试信息，取代System.out.println()
        Logger logger = Logger.getGlobal();
        logger.info("start process...");
        logger.warning("memory is running out...");
        logger.fine("ignored.");
        logger.severe("process will be terminated...");
    JDK的Logging定义了7个日志级别，默认级别是INFO，因此，INFO级别以下的日志，不会被打印出来。使用日志级别的好处在于，调整级别，就可以屏蔽掉很多调试相关的日志输出。

### Commons Logging
    Commons Logging的特色是，它可以挂接不同的日志系统，并通过配置文件指定挂接的日志系统。默认情况下，Commons Loggin自动搜索并使用Log4j（Log4j是另一个流行的日志系统），如果没有找到Log4j，再使用JDK Logging。
        Log log = LogFactory.getLog(Main.class);
        log.info("start...");
        log.warn("end.");
    实例变量log的获取方式是LogFactory.getLog(getClass())，虽然也可以用LogFactory.getLog(Person.class)，但是前一种方式有个非常大的好处，就是子类可以直接使用该log实例。

## 反射
    反射就是Reflection，Java的反射是指程序在运行期可以拿到一个对象的所有信息。
    反射是为了解决在运行期，对某个实例一无所知的情况下，如何调用其方法。

### Class类
    class是由JVM在执行过程中动态加载的。JVM在第一次读取到一种class类型时，将其加载进内存。
    每加载一种class，JVM就为其创建一个Class类型的实例，并关联起来。注意：这里的Class类型是一个名叫Class的class
    以String类为例，当JVM加载String类时，它首先读取String.class文件到内存，然后，为String类创建一个Class实例并关联起来
    一个Class实例包含了该class的所有完整信息
    通过Class实例获取class信息的方法称为反射（Reflection）。
    String s = "Hello";
    Class cls = s.getClass();
    获取类对象：
    1，类实例.getClass()
    2,类.class

### 访问字段
    · Field getField(name)：根据字段名获取某个public的field（包括父类）
    · Field getDeclaredField(name)：根据字段名获取当前类的某个field（不包括父类）
    · Field[] getFields()：获取所有public的field（包括父类）
    · Field[] getDeclaredFields()：获取当前类的所有field（不包括父类）
    一个Field对象包含了一个字段的所有信息：public int Student.score
    · getName()：返回字段名称，例如，"name"；
    · getType()：返回字段类型，也是一个Class实例，例如，String.class；
    · getModifiers()：返回字段的修饰符，它是一个int，不同的bit表示不同的含义。
    获取到字段的实例之后就可以使用Field.get(Object)获取指定实例的指定字段的值、通过Field.set(Object, Object)设置字段值

### 调用方法
    通过Class实例获取所有Method信息
    · Method getMethod(name, Class...(方法参数的class))：获取某个public的Method（包括父类）
    · Method getDeclaredMethod(name, Class...)：获取当前类的某个Method（不包括父类）
    · Method[] getMethods()：获取所有public的Method（包括父类）
    · Method[] getDeclaredMethods()：获取当前类的所有Method（不包括父类）

    Method对象方法：
    · getName()：返回方法名称，例如："getScore"；
    · getReturnType()：返回方法返回值类型，也是一个Class实例，例如：String.class；
    · getParameterTypes()：返回方法的参数类型，是一个Class数组，例如：{String.class, int.class}；
    · getModifiers()：返回方法的修饰符，它是一个int，不同的bit表示不同的含义。

    调用方法：
    Method实例.invoke(对象实例, 参数...)

### 调用构造方法
    Constructor对象类似Method对象，它包含一个构造方法的所有信息，可以创建一个实例。
    · getConstructor(Class...)：获取某个public的Constructor；
    · getDeclaredConstructor(Class...)：获取某个Constructor；
    · getConstructors()：获取所有public的Constructor；
    · getDeclaredConstructors()：获取所有Constructor。
    注意Constructor总是当前类定义的构造方法，和父类无关，因此不存在多态的问题。

### 获取继承关系
    Class i = Integer.class;
    Class n = i.getSuperclass();
    获取实现的接口：
    Class[] is = s.getInterfaces();

## 注解
    注解是放在Java源码的类、方法、字段、参数前的一种特殊“注释”：@Resource("hello")
    注释会被编译器直接忽略，注解则可以被编译器打包进入class文件，因此，注解是一种用作标注的“元数据”。
    注解（Annotation）是Java语言用于工具处理的标注：
### 定义注解
    

