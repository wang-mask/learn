# Docker 学习

## 为什么会出现Docker？
    镜像 = 环境 + 应用
    一款产品会有开发-上线两套环境，部署相同环境非常麻烦，能不能发布一个带环境的项目？
        项目打包带上镜像（环境），下载发布的景象即可运行，并且实现了隔离机制（可以将服务器分隔成几个互不相干的环境）
        Docker 每一个容器就是一台轻量化的虚拟机（具有最基本的环境，再加上你需要的jdk等环境）
    传统的虚拟机运行一个完整的OS，在这个系统上安装运行软件。
    而容器内的应用直接运行在宿主机，没有自己的内核和虚拟硬件更加轻便。每个容器具有自己独立的环境，自己的文件系统，互不影响

## 结构
    1. 镜像，类比如java的模版类（环境，如tomcat环境），通过这个类可以new多个实例（容器）
    2. 容器就是简易的linux，能够独立运行一个或一组应用
    3. 仓库 存放镜像的地方（公有、私有）

## 怎么下载镜像？
    先在本地查找镜像，如果没有则去远程DockerHub下载镜像

## 工作原理
    Client-Server结构系统，Docker守护进程运行在主机上（linux等），客户端通过socket与交流
    同一主机上多个容器共享虚拟机的内核

## 常用命令
    docker command --help 帮助文档
    docker info 显示系统信息
### 镜像命令
    docker images   显示下载的镜像
        选项 -a，显示所有
            -q，只显示id
    docker search imageName 搜索image
        选项 -f 过滤（-f=STARS=3000，星大于3000的）
    docker pull imageName[:tag] 下载镜像
    docker rmi -f 镜像id 删除镜像

## 容器命令
    docker run [可选参数] image 新建容器并命名
        参数：
        --name="name"
        -d 后台运行，后台运行时必须运行着应用不然会自动停止，除非还有前台应用（交互命令行）
        -it 以交互方式运行，进入容器查看内容
            docker run -it image bin/bash
        -p 指定容器端口
            -p ip:主机端口:容器端口
            -p 容器端口
        -v 本地目录:容器目录挂载卷
        -e key:value 配置环境
        --link 另一个容器名 将另一个容器名写到host中
        --net 网络名  所属网络
        run后不加-的命令就是追加到新建的容器内运行，如bin/bash命令
    exit 退出容器，并停止该容器
    ctrl + p + q 退出容器，但容器继续运行
    docker exec -it 容器id /bin/bash  进入容器，开启一个新的终端控制
    或者 docker attach 容器id，但是会进入正在输出的命令行（执行代码中）
    docker ps [参数]    查看正在运行的容器
        -a 查看正在和历史容器
        -q 只显示容器编号
    docker rm 容器id 删除容器
    docker start 容器id 启动容器
    docker restart 容器id   重启容器
    docker stop 容器id  停止容器
    docker kill 容器id  强制停止容器

    dock logs -tf 容器id 显示该容器所有日志
    dock logs -tf --tail n 容器id 显示n条日志

    docker top 容器id 查看容器内的进程
    docker inspect 容器id 查看容器元数据

    docker cp 容器id:容器内文件路径 主机路径 复制容器内文件到本机

## Docker 部署镜像流程
    1. docker search 或者官网搜索
    2. docker pull下载镜像到本地
    3. docker run 启动容器

## 可视化 portainer
    docker run -d -p 9000:9000 --restart=always -v /var/run/docker.sock:/var/run/docker.sock --name prtainer-test portainer/portainer
    然后进去localhost:9000即可管理docker

## 镜像分层
    镜像类似于linux的联合文件系统，每添加一个环境就相当于加了一层，如果与已有的环境有共用的部分，则公共用不会重新下载。
    所有的镜像都是只读的，你的操作只是在镜像层上的容器层上的操作

## 提交镜像
    docker commit -m "描述信息" -a "作者" 容器id 目标镜像名:[TAG] 会将本地修改过后的容器保存到本地镜像，以后就可以直接使用这个修改过的镜像

## 容器数据卷
    将容器数据持久化，防止容器消失，数据消失
    即将容器内的数据同步到主机，双向绑定，将容器内的目录挂载到主机上
    docker -it -v 主机目录:容器内目录

## 具名和匿名挂载
    匿名挂载：run时使用-v 只写容器内目录等参数但是不指定映射关系
        docker volume ls 查看挂载的券
    具名挂载：run时使用-v 挂载名:容器内目录
    所有docker容器内的卷，如果没有指定主机目录的情况下都是在主机的 /var/lib/docker/volumes/xxxxx/_data下

## 数据卷容器
    同步不同容器间内容，一个容器的数据卷映射同步到另一个容器的数据卷
    启动需要同步的镜像时 run 加参数 --volume-form 上一个容器名
    这两个容器就有同名数据卷，各个容器内的卷互为备份，但是内容实时同步，并不会应为一个删除了，另一个就不存在了
## Dockerfile
    用来构建镜像的文件，构建命令脚本
    官方命名Dockerfile，build时就不用指定文件了 
    然后使用build构建镜像，run --> push 
    Dockerfile中指令都是大写，每个指令就是镜像的一层
    命令：
    FROM 镜像名      #基础镜像
    MAINTAINER 姓名+邮箱    # 镜像是谁写的
    RUN     #镜像构造时需运行的命令，添加一些命令yum等
    ADD  文件名（主机上的） 镜像内的路径  # 向基础镜像添加内容（层）,如jdk什么的，压缩包的形式，会自动解压
    WORKDIR     # 镜像的工作目录，即启动镜像运行时进入的目录，不配置的话默认是根目录
    VOLUME     # 挂载的目录
    ESXPOSE     # 暴露端口
    CMD      # 指定容器启动时要运行的命令，可以用&&添加多条命令一起执行
    ENV     # 设置环境变量
    COPY 本机文件 镜像内的目录

    然后docker build -f dockerfilename -t 镜像名:tag .
    docker history 镜像名 查看镜像的构造过程，即dockerfile过程

    Dockerfile示例：
    FROM centos
    MAINTAINER wangwenxiao<724802019@qq.com>
    COPY readme.txt /usr/local/readme.txt

    ADD jdk-8u144-linux-x64.tar.gz /usr/local/

    RUN yum -y install vim
    ENV MYPATH /usr/local
    WORKDIR $MYPATH

    ENV JAVA_HOME $MYPATH/jdk1.8.0_144
    ENV CLASSPATH $JAVA_HOME/lib;$JAVA_HOME/jre/lib
    ENV PATH $PATH:$JAVA_HOME/bin

# 发布自己的镜像
    docker login -u 用户名 -p 密码 
    docker push 用户名/镜像名:tag

# Docker 网络
    ip addr 查看本机ip
    安装docker，就会为docker分配一个网卡docker0，相当于docker的路由器，将网络转发到各个dicker内部
    每启动一个docker就会分配一个ip
    各个docker之间是可以互相ping通的

# --link
    run 的时候加上 --link 要连接的另一个容器名，但这样只能是单向的，把另一个的容器名写到host中
    这样就可以这些link后的 就可以通过容器名连通了

# 自定义网络
    docker network ls   查看所有的网络
    docker network create --drive 网络模式 --subnet 网络/网络段（子网写法，表明网络号） --gateway 网关 网络名
    docker run 的时候指定网络即可 --net 网络名
    自定义网络下各个容器可以使用docker名互相连接
    docker network connect network名称 容器名 用于将一个容器再挂载到另一个网络下，使其可以通过容器名连通（即一个容器有多个ip），一个容器需要跨网络操作时才使用
    
