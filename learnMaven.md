# Maven学习

    根目录：工程名
    |---src：源码
    |---|---main:存放主程序
    |---|---|---java：java源码文件
    |---|---|---resource：存放框架的配置文件
    |---|---test：存放测试程序
    |---pop.xml：maven的核心配置文件

## 1. 什么是Maven
    Maven 是专门用于构建和管理Java相关项目的工具。使用Maven管理项目的好处主要有两点，
      1. 其一是使用Maven管理的 Java 项目都有着相同的项目结构。 有一个pom.xml 用于维护当前项目都用了哪些jar包；所有的java代码都放在 src/main/java 下面； 所有的测试代码都放在src/test/java 下面 
      2. 其二是便于统一维护jar包。maven风格的项目，把所有的jar包都放在了本地"仓库“ 里，然后哪个项目需要用到这个jar包，只需要给出jar包的名称和版本号就行了，这样就实现了jar包共享，避免每一个项目都有自己的jar包带来的麻烦。
|目录|目的|
|----|----|
|${basedir}|存放pom.xml和所有的子目录|
|${basedir}/src/main/java|项目的java源代码|
|${basedir}/src/main/resources|项目的资源，比如说property文件，springmvc.xml|
|${basedir}/src/test/java|项目的测试类，比如说Junit代码|
|${basedir}/src/test/resources|测试用的资源|
|${basedir}/src/main/webapp/WEB-INF|web应用文件目录，web项目的信息，比如存放web.xml、本地图片、jsp视图页面|
|${basedir}/target|打包输出目录|
|${basedir}/target/classes|编译输出目录|
|${basedir}/target/test-classes|测试编译输出目录|
|Test.java|Maven只会自动运行符合该命名规则的测试类|
|~/.m2/repository|Maven默认的本地仓库目录位置|

## 2. pom.xml 配置文件
    POM 中可以指定以下配置：
      项目依赖
      插件
      执行目标
      项目构建 profile
      项目版本
配置文件类似html的标签语言
|节点|描述|
|----|----|
|project|工程的根标签。|
|modelVersion|模型版本需要设置为 4.0|
|groupId|这是工程组的标识。它在一个组织或者项目中通常是唯一的。例如，一个银行组织 com.companyname.project-group 拥有所有的和银行相关的项目。|
|artifactId|这是工程的标识。它通常是工程的名称。例如，消费者银行。groupId 和 artifactId 一起定义了 artifact 在仓库中的位置。|
|version|这是工程的版本号。在 artifact 的仓库中，它用来区分不同的版本。例如：com.company.bank:consumer-banking:1.0com.company.bank:consumer-banking:1.1

## 2. Maven 构建生命周期
    Maven 有以下三个标准的生命周期：
      clean：项目清理的处理
      default(或 build)：项目部署的处理
      site：项目站点文档创建的处理
    生命周期的各个阶段都是由插件完成的
    在一个生命周期中，运行某个阶段的时候，它之前的所有阶段都会被运行，也就是说，如果执行 mvn clean 将运行以下两个生命周期阶段：
      pre-clean, clean

    Clean 生命周期如下：
      pre-clean：执行一些需要在clean之前完成的工作
      clean：移除所有上一次构建生成的文件
      post-clean：执行一些需要在clean之后立刻完成的工作

    buid生命周期如下：
|阶段|处理|描述|
|----|----|----|
验证 validate|验证项目|验证项目是否正确且所有必须信息是可用的
编译 compile|执行编译|源代码编译在此阶段完成
测试 Test|测试|使用适当的单元测试框架（例如JUnit）运行测试。
包装 package|打包|创建JAR/WAR包如在 pom.xml 中定义提及的包
检查 verify|检查|对集成测试的结果进行检查，以保证质量达标
安装 install|安装|安装打包的项目到本地仓库，以供其他项目使用
部署 deploy|部署|拷贝最终的工程包到远程仓库中，以共享给其他开发人员和工程   
    Site 生命周期：
      Maven Site 插件一般用来创建新的报告文档、部署站点等。
      pre-site：执行一些需要在生成站点文档之前完成的工作
      site：生成项目的站点文档
      post-site： 执行一些需要在生成站点文档之后完成的工作，并且为部署做准备
      site-deploy：将生成的站点文档部署到特定的服务器上

## 3. Maven 构建配置文件
    构建配置文件是一系列的配置项的值，可以用来设置或者覆盖 Maven 构建默认值。
构建配置文件的类型
|类型|在哪定义|
|----|----|
|项目级（Per Project）|定义在项目的POM文件pom.xml中|
|用户级 （Per User）|定义在Maven的设置xml文件中 (%USER_HOME%/.m2/settings.xml)|
全局（Global）|定义在 Maven 全局的设置 xml 文件中 (%M2_HOME%/conf/settings.xml)|

## 4. 包管理
    公共的依赖可以在 pom.xml 里不用一个个加到<dependences></dependences>标签中，而是写一个父 pom.xml 将公用的依赖放在父 pom.xml 文件的<dependences></dependences>标签中
