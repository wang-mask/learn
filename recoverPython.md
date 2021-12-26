# Python 复习

## 基础语法
    单行注释 #
    多行注释''' xxxxx '''
    以缩进来区别代码块
    多行语句 使用 \ 来换行
    Python 中单引号 ' 和双引号 " 使用完全相同
    使用三引号(''' 或 """)可以指定一个多行字符串
    反斜杠可以用来转义，使用 r 可以让反斜杠不发生转义。 如 r"this is a line with \n" 则 \n 会显示
    字符串可以用 + 运算符连接在一起，用 * 运算符重复。
    print 默认输出是换行的，如果要实现不换行需要在变量末尾加上 end=""：print( x, end=" " )
    将整个模块(somemodule)导入，格式为： import somemodule
    从某个模块中导入某个函数,格式为： from somemodule import somefunction

## 数据类型
    Python 中的变量不需要声明。每个变量在使用前都必须赋值，变量赋值以后该变量才会被创建。
    使用del语句删除一些对象引用。
    del var1[,var2[,var3[....,varN]]]
    类型转换 Type(varName)
### 数字类型
    Python3 支持 int、float、bool、complex（复数）。
    数值的除法包含两个运算符：/ 返回一个浮点数，// 返回一个整数。
### String字符串
    字符串用单引号 ' 或双引号 " 括起来，同时使用反斜杠 \ 转义特殊字符。
    截取：变量[头下标:尾下标] 索引值以 0 为开始值，-1 为从末尾的开始位置。左闭右开
    加号 + 是字符串的连接符， 星号 * 表示复制当前字符串，与之结合的数字为复制的次数。
    使字符串不发生转译 r'Ru\noob'
    python字符串不可改变，可使用[下标]访问某个字符，但不可以赋值
### List 列表  []
    列表中元素的类型可以不相同
    列表是写在方括号 [] 之间、用逗号分隔开的元素列表。
    截取 变量[头下标:尾下标:步长]
    加号 + 是列表连接运算符，星号 * 是重复操作。
    列表中的元素可以改变，使用下标访问赋值
### Tuple 元组 ()
    元组（tuple）与列表类似，不同之处在于元组的元素不能修改。
    元组写在小括号 () 里，元素之间用逗号隔开。
    元组的访问跟列表一致
    构造元组：
    tup1 = ()    # 空元组
    tup2 = (20,) # 一个元素，需要在元素后添加逗号
### Set 集合 {}
    集合（set）是由一个或数个形态各异的大小整体组成的，构成集合的事物或对象称作元素或是成员。
    基本功能是进行成员关系测试和删除重复元素。
    可以使用大括号 { } 或者 set() 函数创建集合，注意：创建一个空集合必须用 set() 而不是 { }，因为 { } 是用来创建一个空字典。
    两个集合操作： - 差集、 ｜ 并集、 & 交集、
### Dictionary 字典 {}
    列表是有序的对象集合，字典是无序的对象集合。两者之间的区别在于：字典当中的元素是通过键来存取的，而不是通过偏移存取。
    字典是一种映射类型，字典用 { } 标识，它是一个无序的 键(key) : 值(value) 的集合。
    键(key)必须使用不可变类型。
    在同一个字典中，键(key)必须是唯一的。
    dict = {}
    dict['one'] = "1 - 菜鸟教程"
    tinydict = {'name': 'runoob','code':1, 'site': 'www.runoob.com'}
    构造函数：dict()

## 运算符
    **	幂 - 返回x的y次幂
    //	取整除 - 向下取接近商的整数
    :=	海象运算符，可在表达式内部为变量赋值。
        if (n := len(a)) > 10:
    and 逻辑和
    or 逻辑或
    not 逻辑非
    in 
    not in
    is	is 是判断两个标识符是不是引用自一个对象，is 用于判断两个变量引用对象是否为同一个， == 用于判断引用变量的值是否相等。
    not is

## 数字类型
    数据类型是不允许改变的,这就意味着如果改变数字数据类型的值，将重新分配内存空间。
    常用函数：abs(), ceil(), floor(), max(), min(), round(x, n), sqrt()

## 字符串
    字符串就像数组可以使用下标访问单个字符
    变量[头下标:尾下标] 左闭右开
    索引值以 0 为开始值，-1 为从末尾的开始位置。
    + 拼接
    * 重复
    in 子串是否存在
### 格式化
    print ("我叫 %s 今年 %d 岁!" % ('小明', 10))
    %c 字符
    %s 字符串
    %d	 格式化整数
    %f	 格式化浮点数字，可指定小数点后的精度

    f-string
    name = 'Runoob'
    'Hello %s' % name
    f-string 格式化字符串以 f 开头，后面跟着字符串，字符串中的表达式用大括号 {} 包起来，它会将变量或表达式计算后的值替换进去
    name = 'Runoob'
    f'Hello {name}'  # 替换变量
    print(f'{x+1}')
### 函数
    str.count(str, beg= 0,end=len(string)) 返回子串出现次数
    str.find(str, beg=0, end=len(string))
    str.index(str, beg=0, end=len(string))
    join(seq)
    len(string)
    lower()， upper()大小写转换
    str.replace(old, new [, max])
    str.split(str="", num=string.count(str))

## 列表
    列表的数据项不需要具有相同的类型
    列表的数据项可以直接使用下标来修改
    删除 del list[n]
    len(), max(), min(), list(seq)
    list1 + list2 列表拼接
    list.append()
    list.count()
    list.index()
    list.insert(index, obj)
    list.pop() 返回list队尾元素
    list.popleft() 返回list队首元素
    list.remove(obj)
    list.reverse()
    list.sort()

## 元祖
    元组与列表类似，不同之处在于元组的元素不能修改。
    所谓元组的不可变指的是元组所指向的内存中的内容不可变。
    元组使用小括号 ( )，列表使用方括号 [ ]。
    元组中只包含一个元素时，需要在元素后面添加逗号 , tup1 = (50,)
    len(), max(), min(), tuple(iterable)
## 字典
    字典的每个键值 key=>value 对用冒号 : 分割，每个对之间用逗号(,)分割，整个字典包括在花括号 {} 中
    删除 del tinydict['Name']
    len(), str()
    radiansdict.get(key, default=None)
    key in dict
    radiansdict.items()
    radiansdict.keys()
    radiansdict.values()

## 集合
    集合（set）是一个无序的不重复元素序列。
    可以使用大括号 { } 或者 set() 函数创建集合，注意：创建一个空集合必须用 set() 而不是 { }，因为 { } 是用来创建一个空字典。
    s.add()
    s.remove()
    len()

## 循环
    for <variable> in <sequence>:
        <statements>
    for i in range(n):

## 迭代器与生成器
    迭代器是一个可以记住遍历的位置的对象。
    迭代器对象从集合的第一个元素开始访问，直到所有的元素被访问完结束。迭代器只能往前不会后退。
    迭代器有两个基本的方法：iter() 和 next()。
    list=[1,2,3,4]
    it = iter(list)    # 创建迭代器对象
    print (next(it))
    for x in it:

    在 Python 中，使用了 yield 的函数被称为生成器（generator）。
    在调用生成器运行的过程中，每次遇到 yield 时函数会暂停并保存当前所有的运行信息，返回 yield 的值, 并在下一次执行 next() 方法时从当前位置继续运行。

## 函数
    函数代码块以 def 关键词开头，后接函数标识符名称和圆括号 ()。
    return [表达式] 结束函数，选择性地返回一个值给调用方，不带表达式的 return 相当于返回 None。
    def 函数名（参数列表）:
        函数体

    不定长参数
    def functionname([formal_args,] *var_args_tuple ):
    加了星号 * 的参数会以元组(tuple)的形式导入，存放所有未命名的变量参数。
    还有一种就是参数带两个星号 **基本语法如下：
    def functionname([formal_args,] **var_args_dict ):
    加了两个星号 ** 的参数会以字典的形式导入。

    匿名函数
    lambda [arg1 [,arg2,.....argn]]:expression
    sum = lambda arg1, arg2: arg1 + arg2
    sum( 10, 20 )

## 模块
    一个模块被另一个程序第一次引入时，其主程序将运行。如果我们想在模块被引入时，模块中的某一程序块不执行，我们可以用__name__属性来使该程序块仅在该模块自身运行时执行。
    if __name__ == '__main__':
        print('程序自身在运行')
    else:
        print('我来自另一模块')

### 包
    一个文件下所有的模块称为包，每个包下必须有一个叫做 __init__.py 的文件才会被认作是一个包（可以包含初始代码）
    引入包 import sound.effects.echo
    如果包定义文件 __init__.py 存在一个叫做 __all__ 的列表变量，那么在使用 from package import * 的时候就把这个列表中的所有名字作为包内容导入。

## 输入、输出
    格式化输出
    str.format()
    print('{}网址： "{}!"'.format('菜鸟教程', 'www.runoob.com'))
    print('{name}网址： {site}'.format(name='菜鸟教程', site='www.runoob.com'))
    print('常量 PI 的值近似为 {0:.3f}。'.format(math.pi))
    print('常量 PI 的值近似为：%5.3f。' % math.pi)

    输入
    str = input("请输入：");

    读写文件
    open(filename, mode)
    f.read()、 f.readline()
    f.write(string) 
    f.close()

## 类对象
    类有一个名为 __init__() 的特殊方法（构造方法），该方法在类实例化时会自动调用
    self代表类的实例，而非类
    在类的内部，使用 def 关键字来定义一个方法，与一般函数定义不同，类方法必须包含参数 self, 且为第一个参数，self 代表的是类的实例。
    class Complex:
        def __init__(self, realpart, imagpart):
            self.r = realpart
            self.i = imagpart
        def speak(self):
            print("%s 说: 我 %d 岁。" %(self.name,self.age))
    类的私有属性
    __private_attrs：两个下划线开头，声明该属性为私有，不能在类的外部被使用或直接访问。在类内部的方法中使用时 self.__private_attrs。

    Python 中只有模块（module），类（class）以及函数（def、lambda）才会引入新的作用域

## 作用域
    当内部作用域想修改外部作用域的变量时，就要用到 global 和 nonlocal 关键字了。
    当函数内想要改变全局变量时，要先对全局变量进行 global varName操作
    如果要修改嵌套作用域（enclosing 作用域，外层非全局作用域）中的变量则需要 nonlocal 关键字

# 常用函数
    None 
    if elif
## 类型转换
    int()
    str()

## 列表：
    l1 = [0]* 9 初始化
    multilist2 = [[0]*5 for row2 in range(3)] 初始化二维 3 * 5
    list(val)
    增：list.append()、list.insert(inex, val)
    删：list.pop(index), 默认最后一个、list.remove(val)
    查：list.index()、list.count()
    list.reverse()
    list.sort()
    
## 字典
    d1 = {key:val} 
    迭代器：d1.itesm()、 d1.keys(), d1.values()
    增改：dict[key] =new_val 
    删：dict.pop(key)

## 集合：
    s1 = set()
    增：s1.add()
    删：s1.remove()
    查：x in s1
    去重：new_list = list ( set(list) )
    
## 字符串操作：
    字符串是不可更改的对象，因此无法直接修改字符串的某一位字符。 一种可行的方式是：将字符串转换为列表，修改列表的元素后，再重新连接为字符串。
    s = "thisisastring"
    l = list(s)
    l[0] = "T"
    
    ''.join(list)
    .split()
    len()
    查： str.finx(), str.index()
    
