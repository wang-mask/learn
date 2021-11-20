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
    kubectl get 资源名称 查看资源pods、replicationcontrollers（rc）、services，rs,ds，jobs,svc
    kubectl get pods -o wide 显示详细信息 
    
    kubectl describe 资源类型 资源名称 显示细节，比get更多的细节
    kubectl expose re kubia --type=LoadBalancer --name kubia-http 创建服务
    kubectl scale re kubia --replicas=3   增加期望的副本书，即一个镜像运行出多个pod，也可以将rc更改为job等
    kubectl create -f kubia-manual.yaml 从yaml文件创建pod
    kubectl get po kubia-manual -o yaml 查看pod的yaml文件
    kubectl logs podName 查看pod日志,单个容器的pod
    kubectl logs kubia-manual -c 容器名称 查看pod内某个容器的日志
    kubectl port-forward podName 本机端口:pod端口 将本地端口映射到pod端口
    
    kubectl exec podname -- command # 在该pod内执行命令
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
### 标签
    kubectl label pod podname key=value     # 给pod添加标签
        --overwrite 覆盖原有的标签
    kubectl label po podname key=value   添加/删除标签
    kubectl get po --show-labels 显示标签

### rc/rs相关
    kubectl delete rc rcname    # 删除rc，其管理的pod也将删除
        --cascade=false # 删除rc，但pod保持运行
    kubectl edit rc rcname    # 进入该rc的yaml编辑页面，修改即可。但是修改后pod并不会得到自动更新，只有把之前的pod删除后才能让更新
    kubectl scale rc rcname --replicas=n  # 更改rc所控制的副本数量
### 服务相关 svc
    kubectl expose 
    
# 资源
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

## 副本机制和其他控制器
    你希望你的部署能自动 保待运行， 并且保持健康， 无须任何手动干预。 要做到这 一点， 你几乎不会直接创
    建pod , 而是创建ReplicationController或Deployment这样的资源， 接着由它们来创 建并管理实际的pod。他们的作用就是当pod挂掉时能自动的重启服务

    存活探针:检查容器是否还在运行。 可以 为 pod 中的每个容器单独指定存活探针。 如果探测失败， Kubemetes 将定期执行探 针并重新启动容器。
    在yaml文件下的spec的containers字段的livenessProbe的httpGet字段设置，有两个属性path（http请求的路径）和port（探针连接的网络接口）
        path:/
        prot:8080
        initialDelaySeconds: 15     # 第一次探测延迟15s，防止刚启动就探测
    探针由承载pod的节点上的Kubelet 执行，如果节点挂了就无能无力了，此时要使用ReplicationController或类似机制管理pod
    
### ReplicationController
    ReplicationController是一种Kubemetes资源，可确保它的pod始终保持运行状态。如果pod因任何原因 消失(例如节点从集群中消失或由于该pod已从节点中逐出)， 则ReplicationController 会注意到缺少了pod并创建替代pod。
    一个ReplicationController有三个主要部分
        • label selector ( 标签选择器)， 用于确定ReplicationController作用域中有哪些pod
        • replica count (副本个数)， 指定应运行的pod 数量
        • podtemplate (pod模板)， 用于创建新的pod副本
    修改rc模版后，该模板仅影响由此 ReplicationController 创建的新 pod
    创建rc：
    也可以不指定标签，rc为自动给用模版创建的pod添加标签，并管理他们
        apiVersion: v1
        kind: ReplicationController     # 创建rc
        metadata:
            name: kubia     # rc的名称
        spec:
            replicas: 3     # 实例目标数
            selector:
                app: kubia      # 选择器，决定了管理哪些pod
            template:       # 创建pod所用的模版
                metadata:
                    labels:
                        app: kubia
                spec:
                    containers:
                    - name: kubia
                      image: wangmask/kubia
                      ports:
                      - containerPort: 8080
    更改pod的标签，可以将该pod移除该rc，也可以移入另一个rc

### ReplicaSet
    ReplicaSet是新一 代的 ReplicationController, 并且将 其完全替换掉 (ReplicationController 最终将被弃用)
    ReplicaSet 的行为与ReplicationController 完全相同， 但pod 选择器的表达能力 更强。
    创建：
        apiVersion: apps/v1beta2  # rs属于v1beta2版本的api
        kind: ReplicaSet     # 创建rs
        metadata:
            name: kubia     # rc的名称
        spec:
            replicas: 3     # 实例目标数
            selector:
                metchLabels:  # metchLabels 选择器
                    app: kubia
            template:       # 创建pod所用的模版
                metadata:
                    labels:
                        app: kubia
                spec:
                    containers:
                    - name: kubia
                      image: wangmask/kubia
                      ports:
                      - containerPort: 8080
    更强大的选择器：
    selector:
        matchExpressions:
            - key: app      ## pod的key
              operator: In  # 操作符
              values:   # 值
                - kubia
    每个表达式都必须 包含一个key、 一个operator (运算符)，并且可能还有一个values的列表(取决于 运算符)。
    你会看到四个有效的运算符:
        • In : Label的值 必须与其中 一个指定的values 匹配。
        • NotIn : Label的值与任何指定的values 不匹配。
        • Exists : pod 必须包含一个指定名称的标签(值不重要)。使用此运算符时，不应指定 values字段。
        • DoesNotExist : pod不得包含有指定名称的标签。values属性不得指定 
### DaemonSet
    DaemonSet让分别在每一个节点上运行一个该pod
    apiVersion: apps/v1beta2 
    kind: DaemonSet 
    metadata:
        name: ssd-rnonitor
    spec:
        selector: 
            matchLabels:
                app: ssd-monitor
        template:
            metadata: 
                labels:
                    app: ssd-monitor
            spec:
                nodeSelector: 
                    disk: ssd 
                containers:
                - name: main
                  image: luksa/ssd-monitior

### job
    job用于完成即停止的任务，任务完成后就停止，不像其他一样会重启该任务
    创建：
    apiVersion: batch/v1 
    kind: Job
    metadata:
        name: batch-job
    spec:
        completions: 5  # 顺序运行5次该镜像，上一个结束后才会启动新的pod
        parallelism: 2 # 指定可以多少个pod并行，即总目标事5个，但是可以两个并行（即又两个容器）  
        template:
            metadata: 
                labels:
                    app: batch-job
            spec:
                restartPolicy: OnFailure  # 重启规则
                containers:
                - name: main
                  image: luksa/batch-job
    若没有在yaml中定义标签选择器，而在template中定义了label，则默认管理template中的label

### CronJob 安排Job定期运行或在将来运行一次
    在计划的时间内，CronJob资源 会创建 Job资源，然后Job创建pod。
    apiVersion: batch/v1beta1 
    kind: CronJob
    metadata:
        name: xxxx
    spec:
        schedule: "0,15,30,45 * * * *"  # 该工作每天每小时0、15、30、45分钟运行
        jobTemplate:
            spec:
                template:
                    metadata: 
                        labels:
                            app: periodic-batch-job
                    spec:
                        restartPolicy: OnFailure 
                        containers:
                        - name: main
                          image: luksa/batch-job
    时间表从左到右包含以下五个条目:
    • 分钟
    • 小时
    • 每月中的第几天
    • 月
    • 星期几
    在该示例中，你希望每 15 分钟运行一 次任务因此 schedule 字段的值应该是 "0, 15,30, 45****"这意味着每小时的0、 15、 30和45 分钟(第一个星号)，每 月的每一天(第二个星号)， 每月(第三个星号)和每周的每一天(第四个星号)。
    
# 服务 让客户端发现pod并与之通讯
    Kubernetes在pod启动前会给已经调度到节点上的pod分配IP地址一—因此 客户端不能提前知道提供服务的 pod 的 IP 地址。
    服务是一种为一组功能相同的 pod 提供单一不变的接入点的资源
    当服务存在时，它的 IP 地址和端口不会改变。 客户端通过 IP 地址和端口号建立连接，这些连接会被路由到提供该服务的任意一个 pod 上。 通过这种方式， 客户端不需要 知道每个单独的提供服务的pod的地址， 这样这些pod就可以在集群中随时被创建 或移除。
    创建： 
    apiVersion: v1 
    kind: Service 
    metadata:
        name: kubia 
    spec:
        sessionAffinity: ClientIP       # 将同一ip访问转发到同一pod，而默认是服务将其随机转发到控制的pod
        ports:
        - port: 80      # 该服务可用的端口
        targetPort: 8080    # 该服务转发到的容器端口
        selector:
            app: kubia
    
    服务中定义多个端口
    spec:
        ports:
        - name: http
          port: 80
          targetPort: 8080
        - name: https
          port: 443
          targetPort: 8443
    在 pod开始运行的时候， Kubernetes会初始化一系列的环境变量指向现在存在 的服务。 如果你创建的服务早于客户端pod的创建， pod上的进程可以根据环境变 量获得服务 的 IP 地址和端口号 。
    
    当前端 pod 需要 后端数据库服 务 pod 时， 可以通过名为 backend-database 的服务将后端 pod 暴露出来，然后前端 pod 通过环境变量BACKEND_DATABASE_SERVICE_HOST和BACKEND_DATABASE_SERVICE_PORT去获得IP地址和端口信息。
    
    在同一个命名空间下，可以使用 http ://服务名:端口 访问其他任何服务
### 连接集群的外部服务
#### endpoint 就是暴露一个服务的 IP 地址和端口的列表
    将服务连接重定向到外部ip和端口，而不是内部pod
    服务并不是和 pod 直接相连的。 相反，有一种资源介于两者之间-—-它就是 Endpoint 资源。
    创建：（前提同名的服务没有设置选择器）
    apiVersion: v1 
    kind: Endpoints 
    metadata:
        name: external-service      # 必须与相应的服务同名 
    subsets:            
      - addresses:
        - ip: 11.11. 11. 11         # 服务将连接重定向到endpoint的ip地址
        - ip: 22.22.22.22 
        ports:
        - port: 80                  # endpoint的目标端口
    在服务创建后创建的容器将包含 服务的环境变量，并且与其 IP : port对的所有连接都将 在服务端点之间进行负载均衡。

### 将服务暴露给外部客户端，外网可以公开访问
    · NodePort
    将一组pod公开给外部客户端的第一种方法是创建一个服务并将其类型设置为 NodePort。通过创建NodePort服务， 可以让Kubemetes在其所有节点上保留一 个端口(所有节点上都使用相同的端口号)， 并将传入的连接转发给作为服务部分的pod。任何节点的IP和预留节点端口访 问NodePort服务。
    apiVersion: v1 
    kind: Service 
    metadata:
        name: kubia-nodeport
    spec:
        type: NodePort
        ports:
        - port: 80     # 服务集群ip的端口号 内部访问的端口号
          targetPort: 8080  # 背后pod的目标端口号
          nodePort: 30123   # 通过集群节点的30123端口可以访问该服务 外部访问的端口号，若不设置，则会默认指定一个
        selector:
            app: kubia
    
    · LoadBalancer
    设置服务的类型为Load Badancer而不是NodePort。 负 载均衡器拥有自己独 一 无二的可公开访问的 IP 地址， 并将所有连接重定向到服务。 可以通过负载均衡器的 IP 地址访问服务。
    apiVersion: v1 
    kind: Service 
    metadata:
        name: kubia-loadbalancer 
    spec:
        type: LoadBalancer 
        ports:
        - port: 80
          targetPort: 8080 
        selector:
            app: kubia
    这时可以通过EXTERNAL-IP来从外部访问服务

    · 通过Ingress暴露服务
    前提要开启ingress控制器


### eadless服务
    解决如何连接到服务背后的所有pod，而不是随机一个
    apiVersion: v1 
    kind: Service 
    metadata:
        name: kubia-headless
    spec:
        clusterIP: None # 这使得服务成为headless的
        ports:
        - port: 80
          targetPort: 8080 
        selector:
            app: kubia
    准备好pod后，现在可以尝试执行DNS查找以查看是否获得了实际的podIP
    查询该服务的dns即可返回该服务所拥有的pod的ip

