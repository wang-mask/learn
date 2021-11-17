# Kubernetes in Action

## Kubernetes 介绍
    --image-mirror-country='cn'
    运行在容器中的进程是运行在主机操作系统上的。
    Kubemetes 使开发者可以自主部署应用，并且控制部署的频率，完全脱离运维团队的帮助。
    
    Kubemetes 抽象了数据中心的硬件基础设施，使得对外暴露的只是 一个巨大的 资源地 。 它让我们在部署和运行组件时，不用关注底层的服务器。使用 Kubemetes 部署多组件应用时，它会为每个组件都选择 一个合适的服务器，部署之后它能够保 证每个组件可以轻易地发现其他组件，并彼此之间实现通信。

    Kubemetes 帮助人们自动化的部署应用，即将任务分配到节点并进行管理，类似集群的操作系统

    Kubernetes 集群架构
        - 主节点 ，它承载着 Kubernetes控制和管理整个集群系统的控制面板
        - 工作节点，它们运行用户实际部署的应用
    控制面板：包含 多个组件，组件可以运行在单个主 节点上或者通过副本分别部署在多个主节点 以确保高可用性
        • Kubernetes API 服务器，你和其他控制面板组件都要和它通信
        • Scheculer，它 调度你的应用(为应用的每个可部署组件分配一 个工作节 点〕
        • Controller Manager，它执行集群级别的功能，如复制组件、持续跟踪工作节点 、处理节点失败等
        • etcd，一个可靠的分布式数据存储，它能持久化存储集群配置
    工作节点：
        • Docker、时rtk 或其他的容器类型
        • Kubelet，它与 API 服务器通信，并管理它所在节点的容器
        • Kubernetes Service Proxy (kube-proxy)，它负责组件之间的负载均衡网络流量

## 命令
    po 是pods简称
    minikube 运行本地镜像较为麻烦，要先push到dockerhub再pull回来运行
    上传镜像：
        docker tag imageName id/imageName
        docker login -u id -p password
        docker push id/imageName
    minikube start --vm-driver=none  --image-mirror-country="cn" 启动
    kubectl run kubia --image=luksa/kubia --port=8080 运行
    kubectl get 资源名称 查看资源pods、replicationcontrollers、services
    kubectl get pods -o wide 显示详细信息 
    kubectl get po --show-labels 显示标签
    kubectl describe pod podName 显示pod细节
    kubectl expose re kubia --type=LoadBalancer --name kubia-http 创建服务
    kubectl scale re kubia --replicas=3   增加期望的副本书，即一个镜像运行出多个pod
    kubectl create -f kubia-manual.yaml 从yaml文件创建pod
    kubectl get po kubia-manual -o yaml 查看pod的yaml文件
    kubectl logs podName 查看pod日志,单个容器的pod
    kubectl logs kubia-manual -c 容器名称 查看pod内某个容器的日志
    kubectl port-forward podName 本机端口:pod端口 将本地端口映射到pod端口
    kubectl label po podname key=value   添加/删除标签
### 运行停止pod
    kubectl delete pods podname 停止pod
    kubectl delete po -1 creation method=manual
    同理也可以使用ns参数停止某一命名空间下的pod
### 命名空间
    kubectl get ns  获取所有命名空间
    kubectl get pods -n kube-system    获取kube-system下的pod
    kubectl create namespace custom-namespac 创建命名空间
    将一个pod放到某个命名空间下，在metadata字段下，使用namespace即可
    kubectl delete all --all 删除所有资源（服务，pod和rc，第一个all指明删除的资源类型）

## POD
    一个 pod是一组紧密相关的容器，它们总是一起运行在同一个工作节点上，以 及同一个 Linux 命名空间中。每 个 pod 就像 一 个独立的逻辑机器，拥有自己的 IP、 主机名、进程等，运行 一个独立 的应用程序 。应用程序可以是单个进程，运行在单 个容器中，也可以是 一个主应用进程或者其他支持进程，每个进程都在自己的容器 中运行 。一 个 pod 的所有容器都运行在同 一 个逻辑机器上，而其他 pod 中的容器， 即使运行在同 一个工作节点上，也会出现在不同的节点上 。
    每个 pod都有自己的凹，并包含一个或多个容器， 每个容器都运行一个应用进程。 pod 分布在不同的工作节点上 。

    ReplicationController:它确保始终存在一个运行中的 pod 实例。 通常， ReplicationController 用于复制 pod C即创建 pod 的多个副本)并让它保持运行

    服务：服务表示一组或多组提供相同服务的pod的静态地址。 到达服务IP和端口的请求将被转发到属于该服务的一个容器的 IP 和端口。

    现在有了一个正在运行的应用， 由 ReplicationController 监控并保持运行， 并通 过服务暴露访问。 

    容器被设计为每个容器只运行一个进程！！
    一个应用就是一个pod，一个应用可能有多个进程，每个进程运行在一个容器，这些容器都在一个pod中

    容器可以通过 localhost 与同一 pod 中的其他容器进行通信

    当决定是将两个容器放入一个 pod还是 两个单独的 pod 时，我们 需要问自己以下问题:
        · 它们需要 一起运行还是可以在不同的主机上运行?
        · 它们代表的是一个整体还是相互独立的组件?
        · 它们必须一起进行扩缩容还是可 以分别进行?
    基本上，我们总是应该倾向于在单独的 pod 中运行容器，除非有特定的原因要求它们是同-pod的一部分
## pod 定义
    KubemetesAPI版本 
    YAML描述的资源类型
    metadata 包括名称、命名空间、标签和关于该容器 的其他信息 。
    spec 包含 pod 内容的实际说明 ， 例如 pod 的容器、卷和其他数据 。
    status包含运行中的pod的当前信息，例如pod所处的条件、 每个容器的描述和状态，以及内部 IP 和其他基本信息 。创建pod时不需要提供该字段

    yaml 文件
    apiVersion: v1  # 文件遵循v1版本的Kubernetes API
    kind: Pod   # 在描述pod
    metadata:   
        name: kubia-manual  # pod名称
        labels: # 标签，给pod分组方便管理
            creation_method: manual     
            env:prod 
    spec:
        containers:
        - image: wangmask/kubia #所用镜像
          name: kubia   #容器名称
          ports:
          - containerPort: 8080 # 监听的端口
            protocol: TCP


## 标签选择器
    kubectl get po -l creation_method=manual #列出creation_method=manual的pod
    kubectl get po -l '!env' 列出没有env标签的pod
    env in (prod, devel)
    kubectl label node gke-kubia-85f6-node-Orrx gpu=true 给node添加标签
    要使pod运行着特定满足需求的节点，可以在yaml文件中：
        spec:
            nodeSelector:
                gpu: "true"
    这样调度该pod时只会放到node的gpu标签为true的节点
