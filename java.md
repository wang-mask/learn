
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
    void check(Person person) throws IllegalArgumentException, ReflectiveOperationException {
        // 遍历所有Field:
        for (Field field : person.getClass().getFields()) {
            // 获取Field定义的@Range:
            Range range = field.getAnnotation(Range.class);
            // 如果@Range存在:
            if (range != null) {
                // 获取Field的值:
                Object value = field.get(person);
                // 如果值是String:
                if (value instanceof String) {
                    String s = (String) value;
                    // 判断值是否满足@Range的min/max:
                    if (s.length() < range.min() || s.length() > range.max()) {
                        throw new IllegalArgumentException("Invalid field: " + field.getName());
                    }
                }
            }
        }
    }

## 泛型
    泛型就是定义一种模板，例如ArrayList<T>，然后在代码中为用到的类创建对应的ArrayList<类型>
    编译器如果能自动推断出泛型类型，就可以省略后面的泛型类型。例如，对于下面的代码：List<Number> list = new ArrayList<Number>();
    使用泛型时，把泛型参数<T>替换为需要的class类型，例如：ArrayList<String>，ArrayList<Number>等；
    不指定泛型参数类型时，编译器会给出警告，且只能将<T>视为Object类型；
    可以在接口中定义泛型类型，实现此接口的类必须实现正确的泛型类型。
    自定义类型只要实现Comparable<T>接口的int compareTo(T o)方法即可使用Arrays.sort()方法排序 
    泛型实现机制：编译器把类型<T>视为Object，编译器根据<T>实现安全的强制转型。
    Java的泛型是由编译器在编译时实行的，编译器内部永远把所有类型T视为Object处理，但是，在需要转型的时候，编译器会根据T的类型自动为我们实行安全地强制转型。
    泛型的局限：
    <T>不能是基本类型，例如int，因为实际类型是Object，Object类型无法持有基本类型：
    无法取得带泛型的Class，编译后它们全部都是Pair<Object>
### extends通配符
    使用类似<? extends Number>通配符作为方法参数时表示：
        方法内部可以调用获取Number引用的方法，例如：Number n = obj.getFirst();；
        方法内部无法调用传入Number引用的方法（null除外），例如：obj.setFirst(Number n);。
    即一句话总结：使用extends通配符表示可以读，不能写。
    使用类似<T extends Number>定义泛型类时表示：
        泛型类型限定为Number以及Number的子类。

### super通配符
    使用<? super Integer>通配符表示：
        允许调用set(? super Integer)方法传入Integer的引用；
        不允许调用get()方法获得Integer的引用。
    唯一例外是可以获取Object的引用：Object o = p.getFirst()。
    换句话说，使用<? super Integer>通配符作为方法参数，表示方法内部代码对于参数只能写，不能读。

## 集合
    集合就是“由若干个确定的元素所构成的整体”
    由于数组初始化后大小不可变，且只能只能按索引顺序存取。所以产生了可变大小的集合
    Java的java.util包主要提供了以下三种类型的集合：
    - List：一种有序列表的集合，例如，按索引排列的Student的List；
    - Set：一种保证没有重复元素的集合，例如，所有无重复名称的Student的Set；
    - Map：一种通过键值（key-value）查找的映射表集合，例如，根据Student的name查找对应Student的Map。
    Java集合的设计有几个特点：一是实现了接口和实现类相分离，例如，有序表的接口是List，具体的实现类有ArrayList，LinkedList等，二是支持泛型，我们可以限制在一个集合中只能放入同一种数据类型的元素

### List
    List的行为和数组几乎完全相同：List内部按照放入元素的先后顺序存放，每个元素都可以通过索引确定自己的位置，
    ArrayList 最常用的 List 的实现类
    实际上，ArrayList在内部使用了数组来存储所有元素。例如，一个ArrayList拥有5个元素，实际数组大小为6（即有一个空位），即将当初原始数组的增添移位自动化，没有空闲位置的时候，ArrayList先创建一个更大的新数组，然后把旧数组的所有元素复制到新数组，紧接着用新数组取代旧数组
    ArrayList把添加和删除的操作封装起来，让我们操作List类似于操作数组，却不用关心内部元素如何移动。
    常用操作：
    .add(int index, E e)默认末尾
    .remove(int index) .remove(Object e)
    .get(int index)
    .size()
    .indexOf() 内部使用equals()方法判断两个元素是否相等
    for (String s : list) {
        `````
    }
    List 转 Array
    Integer[] array = list.toArray(new Integer[3]);
    给toArray(T[])传入一个类型相同的Array，List内部自动把元素复制到传入的Array中
    Array 转 List
    List<String> list = Arrays.asList(array);找不到或无法加载主类 Main

    自定义数组元素要使用indexOf()，必须覆写equals()方法
    编写equals()方法：
    1.先确定实例“相等”的逻辑，即哪些字段相等，就认为实例相等；
    2.用instanceof判断传入的待比较的Object是不是当前类型，如果是，继续比较，否则，返回false；
    3.对引用类型用Objects.equals()比较，对基本类型直接用==比较。

### Map
    Map也是一个接口，最常用的实现类是HashMap
    .get(K key)
    .put(K key, V value)
    for (String key : map.keySet()) {
        
    }
    for (Map.Entry<String, Integer> entry : map.entrySet()) {
        String key = entry.getKey();
        Integer value = entry.getValue();
    }
    Map和List不同的是，Map存储的是key-value的映射关系，并且，它不保证顺序。
    在Map的内部，对key做比较是通过equals()实现的，这一点和List查找元素需要正确覆写equals()是一样的，即正确使用Map必须保证：作为key的对象必须正确覆写equals()方法。
    通过key计算索引的方式就是调用key对象的hashCode()方法，它返回一个int整数。HashMap正是通过这个方法直接定位key对应的value的索引，继而直接返回value。
    
    因此，正确使用Map必须保证：

    1.作为key的对象必须正确覆写equals()方法，相等的两个key实例调用equals()必须返回true；
    2.作为key的对象还必须正确覆写hashCode()方法，且hashCode()方法要严格遵循以下规范：
    3.如果两个对象相等，则两个对象的hashCode()必须相等；
    4.如果两个对象不相等，则两个对象的hashCode()尽量不要相等。

    编写equals()和hashCode()遵循的原则是：
    equals()用到的用于比较的每一个字段，都必须在hashCode()中用于计算；equals()中没有使用到的字段，绝不可放在hashCode()中计算。
    int hashCode() {
        return Objects.hash(firstName, lastName, age);
    }

    既然HashMap内部使用了数组，通过计算key的hashCode()直接定位value所在的索引，那么第一个问题来了：hashCode()返回的int范围高达±21亿，先不考虑负数，HashMap内部使用的数组得有多大？
    实际上HashMap初始化时默认的数组大小只有16，任何key，通过对15取余确定位置。
    在容量不够时，会自动吧长度扩为2倍，并重新计算索引
    如果索引存在冲突，则在冲突的位置使用索引

    也有按照指定顺序存取的map， treeMap

### 使用Properties
    读写配置文件，它的Key-Value一般都是String-String类型的
    Java集合库提供了一个Properties来表示一组“配置”。
    Java默认配置文件以.properties为扩展名，每行以key=value表示，以#课开头的是注释。
    String f = "setting.properties";
    Properties props = new Properties();
    props.load(new java.io.FileInputStream(f));

    String filepath = props.getProperty("last_open_file");
    String interval = props.getProperty("auto_save_interval", "120");
    props.setProperty("language", "Java");
    props.store(new FileOutputStream("C:\\conf\\setting.properties"), "这是写入的properties注释");

    用Properties读取配置文件，一共有三步：
    1.创建Properties实例；
    2.调用load()读取文件；
    3.调用getProperty()获取配置。
    如果有多个.properties文件，可以反复调用load()读取，后读取的key-value会覆盖已读取的key-value：

### Set
    Set用于存储不重复的元素集合
    .add()
    .remove()
    .contains()
    Set实际上相当于只存储key、不存储value的Map。我们经常用Set用于去除重复元素。
    放入Set的元素和Map的key类似，都要正确实现equals()和hashCode()方法，否则该元素无法正确地放入Set。
    最常用的Set实现类是HashSet，实际上，HashSet仅仅是对HashMap的一个简单封装

### Queue
    .add() .offer()
    .remove()  .poll()
    .peek() 获取队首元素
    如果当前Queue是一个空队列，调用remove()方法，它会抛出异常，如果我们调用poll()方法来取出队首元素，当获取失败时，它不会抛异常，而是返回null
    Queue<String> queue = new LinkedList<>();

### PriorityQueue
    Queue<String> q = new PriorityQueue<>();
    放入PriorityQueue的元素，必须实现Comparable接口，PriorityQueue会根据元素的排序顺序决定出队的优先级。若元素没有实现该接口，则在创建PriorityQueue将比较函数传入，提供一个Comparator对象来判断两个元素的顺序。
    Queue<User> q = new PriorityQueue<>(new UserComparator());

    class UserComparator implements Comparator<User> {
        public int compare(User u1, User u2) {
            if (u1.number.charAt(0) == u2.number.charAt(0)) {
                // 如果两人的号都是A开头或者都是V开头,比较号的大小:
                return u1.number.compareTo(u2.number);
            }
            if (u1.number.charAt(0) == 'V') {
                // u1的号码是V开头,优先级高:
                return -1;
            } else {
                return 1;
            }
        }
    }

### Stack
    .push()
    .pop()
    .peek()
    在Java中，我们用Deque可以实现Stack的功能
    因为方法调用栈有容量限制，嵌套调用过多会造成栈溢出，即引发StackOverflowError

### Iterator
    Java的集合类都可以使用for each循环，List、Set和Queue会迭代每个元素，Map会迭代每个key
    Java编译器并不知道如何遍历List，只是因为编译器把for each循环通过Iterator改写为了普通的for循环
    for (Iterator<String> it = list.iterator(); it.hasNext(); ) {
        String s = it.next();
        System.out.println(s);
    }
    使用迭代器的好处在于，调用方总是以统一的方式遍历各种集合类型，而不必关系它们内部的存储结构。
    如果我们自己编写了一个集合类，想要使用for each循环，只需满足以下条件：
    1.集合类实现Iterable接口，该接口要求返回一个Iterator对象；
    2.用Iterator对象迭代集合内部数据（实现hasNext()和 next()函数）。
    class ReverseList<T> implements Iterable<T> {

        private List<T> list = new ArrayList<>();

        public void add(T t) {
            list.add(t);
        }

        @Override
        public Iterator<T> iterator() {
            return new ReverseIterator(list.size());
        }

        class ReverseIterator implements Iterator<T> {
            int index;

            ReverseIterator(int index) {
                this.index = index;
            }

            @Override
            public boolean hasNext() {
                return index > 0;
            }

            @Override
            public T next() {
                index--;
                return ReverseList.this.list.get(index);
            }
        }
    }
    在编写Iterator的时候，我们通常可以用一个内部类来实现Iterator接口，这个内部类可以直接访问对应的外部类的所有字段和方法。
### Collections
    Collections是JDK提供的工具类，同样位于java.util包中。它提供了一系列静态方法，能更方便地操作各种集合。

## IO
    IO 的输入输出是相对于内存来说的
    字节流：InputStream、OuputStream
    字符流：Reader、Writer
    Reader和Writer本质上是一个能自动编解码的InputStream和OutputStream。
    InputStream、OutputStream、Reader和Writer都是同步IO的抽象类，对应的具体实现类，以文件为例，有FileInputStream、FileOutputStream、FileReader和FileWriter
### File对象
    File f = new File("/usr/bin/javac");
    构造一个File对象，即使传入的文件或目录不存在，代码也不会出错，因为构造一个File对象，并不会导致任何磁盘操作。
    · 可以获取路径／绝对路径／规范路径：getPath()/getAbsolutePath()/getCanonicalPath()；
    · 可以获取目录的文件和子目录：list()/listFiles()；
    · 可以创建或删除文件和目录。
### InputStream
    InputStream并不是一个接口，而是一个抽象类，它是所有输入流的超类。
    这个抽象类定义的一个最重要的方法就是int read()，这个方法会读取输入流的下一个字节，并返回字节表示的int值（0~255）。如果已读到末尾，返回-1表示不能继续读取了。
    FileInputStream是InputStream的一个子类
    通过close()方法来关闭流
    public void readFile() throws IOException {
        try (InputStream input = new FileInputStream("src/readme.txt")) {
            int n;
            while ((n = input.read()) != -1) {
                System.out.println(n);
            }
        } // 编译器在此自动为我们写入finally并调用close()
    }
    一次读取多个字节int read(byte[] b)：读取若干字节并填充到byte[]数组，返回读取的字节数
    面向抽象编程原则的应用：接受InputStream抽象类型，而不是具体的FileInputStream类型，从而使得代码可以处理InputStream的任意实现类。
### OutputStream
    和InputStream类似，OutputStream也是抽象类，它是所有输出流的超类。这个抽象类定义的一个最重要的方法就是void write(int b)
    这个方法会写入一个字节到输出流。要注意的是，虽然传入的是int参数，但只会写入一个字节，即只写入int最低8位表示字节的部分（相当于b & 0xff）
    output.write("Hello".getBytes("UTF-8")); // Hello
### Filter模式
    FileInputStream：从文件读取数据，是最终数据源；
    ServletInputStream：从HTTP请求读取数据，是最终数据源；
    Socket.getInputStream()：从TCP连接读取数据，是最终数据源；
    当我们需要给一个“基础”InputStream附加各种功能时，我们先确定这个能提供数据源的InputStream，因为我们需要的数据总得来自某个地方，例如，FileInputStream，数据来源自文件：

    InputStream file = new FileInputStream("test.gz");
    紧接着，我们希望FileInputStream能提供缓冲的功能来提高读取的效率，因此我们用BufferedInputStream包装这个InputStream，得到的包装类型是BufferedInputStream，但它仍然被视为一个InputStream：

    InputStream buffered = new BufferedInputStream(file);
    最后，假设该文件已经用gzip压缩了，我们希望直接读取解压缩的内容，就可以再包装一个GZIPInputStream：

    InputStream gzip = new GZIPInputStream(buffered);
    无论我们包装多少次，得到的对象始终是InputStream，我们直接用InputStream来引用它，就可以正常读取：
    上述这种通过一个“基础”组件再叠加各种“附加”功能组件的模式，称之为Filter模式

    可以把一个InputStream和任意个FilterInputStream组合；
    可以把一个OutputStream和任意个FilterOutputStream组合。
### 序列化
    序列化是指把一个Java对象变成二进制内容，本质上就是一个byte[]数组。
    把Java对象存储到文件或者通过网络传输出去了。
    一个Java对象要能序列化，必须实现一个特殊的java.io.Serializable接口
    把一个Java对象变为byte[]数组，需要使用ObjectOutputStream。它负责把一个Java对象写入一个字节流
    try (ObjectOutputStream output = new ObjectOutputStream(buffer)) {
        // 写入int:
        output.writeInt(12345);
        // 写入String:
        output.writeUTF("Hello");
        // 写入Object:
        output.writeObject(Double.valueOf(123.456));
    }
    反序列化：
    try (ObjectInputStream input = new ObjectInputStream(...)) {
        int n = input.readInt();
        String s = input.readUTF();
        Double d = (Double) input.readObject();
    }
    调用readObject()可以直接返回一个Object对象。要把它变成一个特定类型，必须强制转型。
    反序列化时不调用构造方法，可设置serialVersionUID作为版本号（非必需）
### Reader
    Reader是Java的IO库提供的另一个输入流接口。和InputStream的区别是，InputStream是一个字节流，即以byte为单位读取，而Reader是一个字符流，即以char为单位读取：
    java.io.Reader是所有字符输入流的超类，它最主要的方法是：
    public int read() throws IOException;
    FileReader是Reader的一个子类，它可以打开文件并获取Reader
    我们需要用try (resource)来保证Reader在无论有没有IO错误的时候都能够正确地关闭
    Reader本质上是一个基于InputStream的byte到char的转换器
### Writer
    Writer就是带编码转换器的OutputStream，它把char转换为byte并输出。
    Writer是所有字符输出流的超类，它提供的方法主要有：
    · 写入一个字符（0~65535）：void write(int c)； 
    · 写入字符数组的所有字符：void write(char[] c)；
    · 写入String表示的所有字符：void write(String s)。
    FileWriter就是向文件中写入字符流的Writer
### PrintStream和PrintWriter
    PrintStream是一种FilterOutputStream，它在OutputStream的接口上，额外提供了一些写入各种数据类型的方法：
    写入int：print(int)
    写入boolean：print(boolean)
    写入String：print(String)
    写入Object：print(Object)，实际上相当于print(object.toString())
    PrintStream最终输出的总是byte数据，而PrintWriter则是扩展了Writer接口，它的print()/println()方法最终输出的是char数据。

## 正则
### 规则
    s.matches(regx)
    如果正则表达式有特殊字符，那就需要用\转义。例如，正则表达式a\&c，其中\&是用来匹配特殊字符&的，它能精确匹配字符串"a&c"
    正则表达式在Java代码中也是一个字符串，所以，对于正则表达式a\&c来说，对应的Java字符串是"a\\&c"，因为\也是Java字符串的转义字符，两个\\实际上表示的是一个\
    .匹配一个任意字符
    \d仅限一个数字字符   而\D则匹配一个非数字
    \w可以匹配一个字母、数字或下划线

    修饰符*可以匹配任意个字符，包括0个字符
    修饰符+可以匹配至少一个字符
    修饰符?可以匹配0个或一个字符   
    修饰符{n}指定匹配重复次数
    修饰符{n,m}指定匹配n~m个字符

    [...]可以匹配范围内的字符，[123456789]可以匹配1~9，[^1-9]表示匹配非1～9

    |连接的两个正则规则是或规则，例如，AB|CD表示可以匹配AB或CD
    把公共部分提出来，然后用(...)把子规则括起来表示
### 提取匹配到的部分
    引入java.util.regex包，用Pattern对象匹配，匹配后获得一个Matcher对象，如果匹配成功，就可以直接从Matcher.group(index)返回子串
    Matcher.group(index)方法的参数用1表示第一个子串，2表示第二个子串，0为原字符串
    Pattern p = Pattern.compile("(\\d{3,4})\\-(\\d{7,8})");
    Matcher m = p.matcher("010-12345678");
    if (m.matches()) {
        String g1 = m.group(1);
        String g2 = m.group(2);
        System.out.println(g1);
        System.out.println(g2);
    } else {
        System.out.println("匹配失败!");
    }
    在规则\d+后面加个?即可表示非贪婪匹配(尽可能少的匹配)

### 搜索和替换
    搜索
    Pattern p = Pattern.compile("\\wo\\w");
    Matcher m = p.matcher(s);
    while (m.find()) {
        String sub = s.substring(m.start(), m.end());
        System.out.println(sub);
    }
    获取到Matcher对象后，不需要调用matches()方法（因为匹配整个串肯定返回false），而是反复调用find()方法，在整个串中搜索能匹配上\\wo\\w规则的子串，并打印出来。这种方式比String.indexOf()要灵活得多，因为我们搜索的规则是3个字符：中间必须是o，前后两个必须是字符[A-Za-z0-9_]。

    替换
    使用正则表达式替换字符串可以直接调用String.replaceAll()，它的第一个参数是正则表达式，第二个参数是待替换的字符串
    如果我们要把搜索到的指定字符串按规则替换，比如前后各加一个<b>xxxx</b>，这个时候，使用replaceAll()的时候，我们传入的第二个参数可以使用$1、$2来反向引用匹配到的子串

## 编码&加密
    URL编码的目的是把任意文本数据编码为%前缀表示的文本，便于浏览器和服务器处理；
    Base64编码的目的是把任意二进制数据编码为文本，但编码后数据量会增加1/3。

## 多线程
    当Java程序启动的时候，实际上是启动了一个JVM进程，然后，JVM启动主线程来执行main()方法。在main()方法中，我们又可以启动其他线程。
### 创建新线程
    方式一：从Thread派生一个自定义类，然后覆写run()方法：
    public class Main {
        public static void main(String[] args) {
            Thread t = new MyThread();
            t.start(); // 启动新线程
        }
    }
    class MyThread extends Thread {
        @Override
        public void run() {
            System.out.println("start new thread!");
        }
    }
    方式二：创建Thread实例时，传入一个Runnable实例
    public class Main {
        public static void main(String[] args) {
            Thread t = new Thread(new MyRunnable());
            t.start(); // 启动新线程
        }
    }
    class MyRunnable implements Runnable {
        @Override
        public void run() {
            System.out.println("start new thread!");
        }
    }
    当main线程对线程对象t调用join()方法时，主线程将等待变量t表示的线程运行结束，即join就是指等待该线程结束，然后才继续往下执行自身线程。
### 中断线程
    t.interrupt()
    在其他线程中对目标线程调用interrupt()方法，目标线程需要反复检测自身状态是否是interrupted状态，如果是，就立刻结束运行
    注意，interrupt()方法仅仅向t线程发出了“中断请求”，至于t线程是否能立刻响应，要看具体代码。
    
    另一个常用的中断线程的方法是设置标志位。我们通常会用一个running标志位来标识线程是否应该继续运行，在外部线程中，通过把HelloThread.running置为false，就可以让线程结束
    线程类中 public volatile boolean running = true;  # 线程共享变量
    volatile关键字解决的是可见性问题：当一个线程修改了某个共享变量的值，其他线程能够立刻看到修改后的值。
### 守护线程
    如果有一个线程没有退出，JVM进程就不会退出。所以，必须保证所有线程都能及时结束。 
    守护线程是指为其他线程服务的线程。在JVM中，所有非守护线程都执行完毕后，无论有没有守护线程，虚拟机都会自动退出。
    Thread t = new MyThread();
    t.setDaemon(true);
    t.start();
### 线程同步
    synchronized关键字对一个对象进行加锁
    synchronized保证了代码块在任意时刻最多只有一个线程能执行
    synchronized(共享变量) {
        // 操作
    }
    注意加锁对象必须是同一个实例；