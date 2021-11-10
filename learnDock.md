# Docker 学习

## 为什么会出现Docker？
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