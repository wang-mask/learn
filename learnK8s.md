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
    kubectl get pods name -o yaml       # 以yaml文件的形式显示
    
    kubectl describe 资源类型 资源名称 显示细节，比get更多的细节
    kubectl expose re kubia --type=LoadBalancer --name kubia-http 创建服务
    kubectl scale re kubia --replicas=3   增加期望的副本书，即一个镜像运行出多个pod，也可以将rc更改为job等
    kubectl create -f kubia-manual.yaml 从yaml文件创建pod
    kubectl get po kubia-manual -o yaml 查看pod的yaml文件
    kubectl logs podName 查看pod日志,单个容器的pod
    kubectl logs kubia-manual -c 容器名称 查看pod内某个容器的日志
    kubectl port-forward podName 本机端口:pod端口 将本地端口映射到pod端口
    
    kubectl exec podname -- command # 在该pod内执行命令
    kubect1 exec podname env  查看pod的环境变量 

    kubectl exec -it podname -c 容器名 bash # 进入pod中的某个容器

    kubectl describe nodes

    kubectl cp foo-pod:/var/log/foo.log foo.log # 将容器文件叫作/var/log/foo.log复制到本机

    kubectl cp localfile foo-pod:/etc/remotefile # 将本地文件复制到容器 
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

### 升级
    kubectl rolling-update kubia-v1 kubia-v2 --image=luksa/kubia:v2  # 将原来名为kubia-v1的rc替换为kubia-v2，并使用luksa/kubia:v2作为新的镜像

### Deployment
    kubectl create -f kubia-deploymen七vl.yaml --record  # 使用了 --record 选项。 这个选项会记录历史版本号
    kubectl get deployment name
    kubectl describe de­ployment name
    kubectl rollout status deployment name # 专门用于查看部署状态
    kubectl set image deployment dep名称 容器名称=新镜像名 # 更新deployment的镜像
    kubectl rollout undo deployment name  # 回滚到上一个版本
    kubectl rollout history deployment name # 显示升级回滚历史
    kubectl rollout pause deployment name  # 暂停更新
    kubectl rollout resume deployment name # 恢复更新
## 修改资源
    kubectl edit deployment name   # 直接进入文本编辑模式
    kubectl patch deployment kubia -p'{"spec": {"template": {"spec": {"containers": [ {"name": "nodejs", "image": "luksa/kubia:v2"}]}}}}'  # 修改耽搁资源属性
    kubectl apply -f kubia-deployment-v2.yaml # 通过一 个完整的YAML或JSON文件，应用其中新的值来修改对象。如果YAML/JSON中指定的对象不存在，则会被创建。
    kubectl replace -f kubia-deployment一v2.yaml  # 将原有对象替换为YAML/JSON文件中定义的新对象。与apply命令相反， 运行这个命令前要求对象必须存在，否则会打印错误
    kubectl set image deployment kubia nodejs=luksa/kubia:v2 # 修改Pod、ReplicationController、Deployment、DernonSet、Job或ReplicaSet内的镜像
    
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

## 卷：将主机磁盘挂载到容器
    即容器共享主机某一部分文件系统，从而实现各个容器之间的文件共享
### emptyDir 卷
    卷从一个 空 目录开始，运行在 pod 内的应用程序可以写入它 需要 的任何文件 。因为卷的生存周期与 pod 的生存周期相 关联 ，所以 当删除 pod 时， 卷的内容就会丢失 。
    apiVersion: v1 
    kind: Pod 
    metadata:
        name: fortune
    spec:
        containers:
        - image: luksa/fortune
          name: html-generator 
          volumeMounts:  # 挂载卷
          - name: html  # 卷名称
            mountPath: /var/htdocs  # 挂载到该容器的哪个文件夹
        - image: nginx:alpine
          name: web-server
          volumeMounts:
          - name: html
            mountPath: /usr/share/nginx/html
            readOnly: true
        volumes:   # 卷声明
        - name: html    # 卷名称
          emptyDir: {}  # 空卷挂在到两个容器上
    为卷来使用的 emptyDir，是在 承载 pod 的 工作节点的实际磁盘上创建的， 因 此其性能取 决于节点的磁盘类型。

### gitRepo 卷
    gitRepo 卷基本上也是 一 个 emptyDir 卷，它通过克隆 Git 仓库并在 pod 启 动时(但在创建容器之前 )检出特定版本来填充数据，即在pod启动时将git仓库内容填充到空卷里面
    但是启动后并不会同步更新卷中内容与远程仓库一致
    volumes: 
    - name: html
      gitPepo:          # 设置远程仓库
        repository: https: //github.com/luksa/kubia-website-example.git # 仓库地址
        revision: master    # 分支名称
        directory: .        # 挂载到卷的根目录

### hostPath卷
    hostPath 卷指向节点文件系统上的特定文件或目录(请参见图 6.4)。 在同一 个节点上运行并在其 hostPath 卷中使用相同路径的 pod 可以看到相同的文件。即将主机某个文件夹挂在容器当中
    实现持久性存储
    hostPath 卷通常用于尝试单节点集群中的持久化存储，譬如 Minikube 创建的集群。

### 跨节点的持久化存储 pv
    研发人员无须向他们的 pod 中添加特定技术的卷， 而是由集群管理员设置底层 存储， 然后通过 Kubernetes API 服务器创建持久卷并注册。 在创建持久卷时， 管理 员可以指定其大小和所支持的访问模式。
    当集群用户需要在其 pod 中使用持久化存储时， 他们首先创建持久卷声明 (PersistentVolumeClaim, 简称 PVC) 清单， 指定所需要的最低容量要求和访问模式， 然后用户将待久卷声明清单提交给 Kubernetes API 服务器， Kubernetes 将找到可匹 配的待久卷并将其绑定到持久卷声明。
    创建持久卷：
    apiVersion: v1
    kind: PersistentVolume
    metadata:
        name: mongodb-pv 
    spec:
        capacity:
            storage: 1Gi    # 持久存储的大小
        accessModes:        
        - ReadWriteOnce     # 可以被单个客户端挂载为读写模式
        - ReadOnlyMany      # 可以被多个客户端挂载为只读模式
        persistentVolumeReclaimPolicy: Retain   # 当声明被释放后，持久存储将会保留
        gcePersistentDisk:   # 指定持久卷支持的实际存储类型
            pdName: mongodb
            fsType: ext4
    
    创建持久卷声明来获取持久卷 pvc
    创建持久卷后不能在pod中直接使用，而先要声明它
    创建持久卷声明
    apiVersion: v1
    kind: PersistentVolumeClaim 
    metadata:
        name: mongodb-pvc       # 挂载到pod上使用的名称
     spec:
        resources: 
            requests:
                storage: 1Gi    # 申请1GiB的存储空间
        accessModes:
        - ReadWriteOnce         # 允许单个客户端读写访问
        storageClassName: ""    # 
    当 创 建好声明， Kubernetes 就会找到适当的持久卷并将其绑定到声明 ，持 久卷 的容量必须足够大以满足声明的需求，并且卷 的访 问模式必须包含声明中指定的访 问模式 。

    在pod中使用持久卷
    apiVersion: v1
    kind: Pod 
    metadata:
        name: mongodb 
    spec:
        containers:
        - image: mongo
        name: mongodb 
        volumeMounts:
        - name : mongodb-data 
            mountPath: / data/ db
        ports:
        - containerPort: 27017
            protocol: TCP
        volumes:
        - name: mongodb-data 
        persistentVolumeClaim:
            claimName: mongodb-pvc  # 引用持久券声明
    
    于动回收持久卷并 使其恢复可用的唯一方法是删除和重新创建持久卷资源。 

# ConfigMap 和 Secret：配置应用程序
    尽管可以直接使用CMD指令指定镜像运行时想要执行的命令， 正确的做法依旧 是借助ENTRYPOINT指令， 仅仅用CMD指定所需的默认参数。
## 为容器定义命令和参数
    dockerfile中：
    ENTRYPOINT ["/bin/for七uneloop.sh"] # 命令
    CMD ["10"]  # 参数
    在 Kubemetes 中定义容器时， 镜像的 ENTRYPOINT 和 CMD 均可以被覆盖， 仅需在容器定义中设置属性 command 和 args 的值， 
    containers:
    - image: some/image
      command: ["/bin/command"] 
      args: ["arg1", "arg2", "arg3"]
      字符串值无须用引号标记，数值需要。

## 为容器设置环境变量
    kind: Pod
    spec:
        containers:
        - image: luksa/fortune:env
          env:
          - name: INTERVAL 
            value: "30"
          name: html-generator
    采用$(VAR)来引用其他环境变量
    env·
    - name : FIRST_VAR
      value:"foo" 
    - name : SECOND_VAR
        value:”$(F工RST VAR)bar”
## ConfigMap
    用以存储配置 数据的Kubernetes资源称为ConfigMap
    Kubemetes 允 许将配置选项分离到单独的资源对象 ConfigMap 中， 本质上就是 一个键 /值对映射，值可以是短字面量，也可以是完整的配置文件。
    将配置文件跟从pod定义中牵出，解耦
    pod 是通过名称引用 ConfigMap 的，因此可以在多环境下使用相同的 pod 定义描述，同 时保持不同的配 置值 以适应不同环境
    创建：
    kubectl create configmap fortune-config --from-literal=foo=bar --from-literal=bar=baz --from-literal=one=two
    通过这条命令创建了 一 个叫作 fortune一config 的 ConfigMap，包含单映射 条目 foo=bar、bar=baz、one=two
    
    kubectl create configmap my-config --from-file=config-file.conf
    kubectl 会在当前目录下查找 config-file . conf 文件，并将 文件内容存储在 ConfigMap 中以 config-file.conf 为键名的条目下 。

    给容器传递ConfigMap 条目作为环境变量
    apiVersion: v1
    kind: Pod 
    metadata:
        name: fortune-env-from-configmap 
    spec:
        containers:
        - image : luksa/fortune:env 
          env :             # 设置环境变量
          - name : INTERVAL # 变量名
            valueFrom: 
                configMapKeyRef:    # 用configMap初始化
                    name: fortune-config # 引用哪个configMap
                    key: sleep-interval # 对应configMap中的哪个key的value
    
    将configMap中所有条目映射到容器的环境变量中
    spec:
        containers:
        - image: some-image
          envFrom:          # 使用envFrom字段
          - prefix: CONFIG_     # 所有环境变量均含前缀CONFIG_，即原本的kay再加上这个前缀
            configMapRef:       # 引用哪个configMap
                name: my-config-map
    Linux 系统挂载文件系统至非空文件夹时通常表现如此。文件夹中只会包含被 挂载文件系统中的文件 ，即便文件夹中原本的文件是不可访问的也是同样如此。

    如果想挂载卷，而又不是将主机整个文件夹挂载过去（会隐藏容器内的文件夹）可以使用subPath属性
    spec:
        containers:
        - image: some/image
          volumeMounts:
          - name: myvolume
            mountPath: /etc/someconfig.conf 
            subPath: myconfig.conf      # 只挂载主机卷下的myconfig.conf 
    使用环境变量或者命令行参数作为配置源的弊端在于无法在 进程运行时更新配置。 将ConfigMap暴露为卷可以达到配置热更新的效果， 无须重新创建pdo 或者重启容器。
    kubectl edit configmap fortune-config 修改configmap会自动同步到configmap对应的卷

# 从应用访问pod元数据以及其他资源
    在之前的章节中， 我们已经了解到如何通过环境变量或者configMap和 secret卷向应用传递配置数据。 
    但是对于那些不能预先知道的数据， 比如pod的IP、 主机名或者是pod自身的名称 (当名称被生成， 比如当pod通过ReplicaSet或类似的控制器生成时)呢?
    Downward API允许我们通过环境变量或者文件(在downwardA釭卷中)的传递pod的元数据。
    
    可以给容器传递以下数据:
    • pod的名称
    • pod的IP
    • pod所在的命名空间
    • pod运行节点的名称
    • pod运行所归属的服务账户的名称
    • 每个容器请求的CPU和内存的使用量
    • 每个容器可以使用的CPU和内存的限制
    • pod的标签
    • pod的注解

## 通过环境变量暴露元数据
    env:
    - name: POD_NAME 
      valueFrom:
            fieldRef:
                fieldPath: metadata.name # 引用metadata.nam
    - name: CONTAINER_CPU_REQUEST_MILLICORES
      valueFrom:
        resourceFieldRef:       # 容器的cpu的内存使用量的引用是resourceFieldRef字段
            resource: requests.cpu
            divisor: 1m # 基数单位

## 通过downwardAPI卷来传递元数据
    apiVersion: v1 
    kind: Pod 
    metadata:
        name: downward 
        labels:
            foo: bar
        annotations: 
            key1: value1 
            key2:
                multi
                line
                value
    volumes:
    - name: downward 
      downwardAPI:
      items:
      - path: "podName"         # 放在容器内挂载卷的padName文件中
        fieldRef:           # pod的名称被写入padName文件中
            fieldPath: metadata.name
      - path: "containerMemoryLimitBytes" # 引用容器中的数据
        resourceFieldRef: 
            containerName: main 
            resource: limits.memory 
            divisor: 1
            
    定义了 一 个叫作 downward 的卷，并且通过/etc/downward目录挂载到我们的容器中，卷所包含的文件会通过卷定义中的downwardAPI.items属性来定义
    要在文件中保存的每一 个 pod 级的字段或者容器资源字段， 都分别在downwardAPI. 江ems 中说明了元数据被保存和引用的 path( 文件名)，
    文件的内容就是元数据字 段和值

## 与 Kubernetes API 服务器交互
    解决要知道其他 pod 的 信息， 甚至是集群中其他资源的信息问题。
    要实现pod与Kubernetes API 服务器交互
    • 确定API服务器的位置
    • 确保是与API服务器进行交互，而不是一个冒名者
    • 通过服务器的认证，否则将不能查看 任何内容以及进行任何操作

    获取api服务器位置：
    在容器内运行：env ｜ grep KUBERNETES_SERVICE，会显示api服务器的ip地址和端口，因为默认会运行一个服务，pod创建是该服务将会写入pod的环境变量
    也可以简单的指向 https://kubernetes 就是api服务器的地址
    配置访问证书，将证书配置为环境变量
    export CURL_CA_BUNDLE=/var/run/secrets/kubernetes.io/serviceaccount/ca.crt
    配置证书后仍然没办法授权允许访问

    获得api服务器授权：
    # 指向内部 API 服务器的主机名
    APISERVER=https://kubernetes.default.svc
    # 服务账号令牌的路径
    SERVICEACCOUNT=/var/run/secrets/kubernetes.io/serviceaccount
    # 读取 Pod 的名字空间
    NAMESPACE=$(cat ${SERVICEACCOUNT}/namespace)
    # 读取服务账号的持有者令牌
    TOKEN=$(cat ${SERVICEACCOUNT}/token)
    # 引用内部证书机构（CA）
    CACERT=${SERVICEACCOUNT}/ca.crt
    # 使用令牌访问 API
    curl --cacert ${CACERT} --header "Authorization: Bearer ${TOKEN}" -X GET ${APISERVER}/api
    访问api加上各个参数就能对部分对象进行修改

## ambassador 容器简化与 API 服务器的交互
    如果一个应用需要查询API服务器(此外还有其他原因)。 除了 像之前章节讲到的直接与API服务器交互， 可以在主容器运行的同时， 启动 一个 ambassador容器，并在其中运行kubecctl proxy命令， 通过它来实现与API服务器的交互。
    在这种模式下， 运行在主容器中的应用不是直接与API服务器进行交互， 而是 通过HTTP协议(不是HTTPS协议)与ambassador连接， 并且由ambassador通过 HTTPS协议来连接API服务器， 对应用透明地来处理安全问题。式同样使用了默认凭证Seeret卷中的文件。
    spec:
        containers:
        - name: main
          image: tu七um/curl
          command: ["sleep", "9999999"] 
        - name: ambassador      # 基于kubectl-proxy的 ambassador
          image: luksa/kubectl-proxy:l.6.2
    这样在mian容器中就可以直接curl localhost:8001 而无需配置
## 访问api服务器的编程库
    • Galang client—https://github.com/kubemetes/client-go
    • Python—https://github.com/kubemetes-incubator/client-python

# Deployment 声明式地升级应用
    Kubemetes 提供 了另一种基于 ReplicaSet 的 资源 Deployment， 并支持声明式地更新应用程序 
    使用 ReplicationController实现自动的滚动升级
    一个yaml可以创建多个资源 使用---分隔不同资源
    apiVersion: v1
    kind: Replicationcontroller 
    metadata:
        name: kubia-v1
    spec:
        replicas: 3
        template:
            metadata: 
                name: kubia 
                labels:
                    app: kubia
            spec:
                containers:
                - image: kuksa/kubia:v1
                  name: nodejs
    ---
    apiVersion: v1
    kind: Service
    metadata:
        name: kubia-v1
    spec:
        type: LoadBalancer 
        selector:
            app: kubia
        ports:
        - port: 80
          targetPort: 8080
    Deployment 是一 种更高阶资源， 用千部署应用程序并以声明的方式升级应用， 而不是通过 ReplicationController 或 ReplicaSet 进行部署， 它们都被认为是更底层的 概念。
    在使用Deployment时，实际的pod 是由 Deployment 的 Replicaset 创建和管理的， 而不是由 Deployment 直接创建和管 理的

    创建：
    apiVersion: apps/v1beta1
    kind: Deployment 
    metadata:
        name: kubia 
    spec:
        replicas: 3 
        template:
            metadata: 
                name: kubia 
                labels:
                    app: kubia 
            spec:
                containers:
                - image: luksa/kubia:v1
                  name: nodejs
    升级应用只需修改deployment的yaml文件即可，
    kubectl set image deployment kubia nodejs=luksa/kubia:v2
    每次升级就会创建新的ReplicaSet，并留下旧的ReplicaSet，将pod逐步从旧的ReplicaSet转移到新的ReplicaSet
    还记得第一次修改 Deployment时留下的 ReplicaSet吗?这个 ReplicaSet便表示 Deployment 的第一次修改版本 。 由 Deployment创建的所有 ReplicaSet表示完整的修 改版本历史，如图 9.1l 所示 。每个 ReplicaSet 都用特定的版本号来保存 Deployment 的完整信息，所以不应该手动删除 ReplicaSet。如果这么做便会丢失 Deployment 的 历史版本记录而导致无法回滚 。
    默认保留两个ReplicaSet，其他更久的会删除

# StatefulSet:部署有状态的多副本应用
    解决pod中的副本如何有独立的存储卷而不和多个副本共享，即在贡献一个存储卷的情况下如何让每个pod实例都保持自己的持久化状态
    RelicaSet或ReplicationController管理的pod副本比较像牛， 这 是因为它们都是 无状态的， 任何时候它们都可以被一 个全新的pod替换。 然而有状态的pod需要不 同的方法， 当 一 个有状态的pod挂掉后(或者它所在的节点故障)， 这 个pod实例需要在别的节点上重建， 但是新的实例必须与被替换的 实例拥有相同的名称、 网络 标识和状态。 这 就是Statefu!Set如何管理pod的。
    Statefulset 保证了pod在重新调度后保留它们的标识和状态。
    与ReplicaSet 不同的是， Statefulset创建的pod副本并不是完全一样的。 每个pod都可以拥有一 组独立的数据卷(持久化状态)而有所区别。 

    一个Statefulset创建的每个pod都有一个从零开始的顺序索引， 这个会体现在 pod的名称和主机名上，同样还会体现在pod对应的固定存储上。 这些pod的名称 则是可预知的， 因为它是由Statefulset 的名称加该实例的顺序索引值组成的。 
    当 一个Statefulset管理的一个pod实例消失后(pod所在节点发生故障， 或有 人手动删除pod,) Statefulset会保证重启一个新的pod实例替换它， 但与ReplicaSet 不同的是， 新的pod会拥有与之前pod完全 一 致的名称和主 机名
    扩容 一个Statefulset会使用下一个还没用到的顺序索引值创建一个新的pod实 例。
    当缩容一个 Statefulset时\ 比较好的是很明确哪个 pod将要被删除。缩容一个 Statefulset将会最先删除最高索引值 的实例

    因为缩容 Statefulset时会保留持久卷声明， 所以在随后的扩容操作中， 新的 pod 实例会使用绑定在持久卷上的相同声明和其上 的数据(如图 10.9 所示)。当你因为 误操作而缩容一个 Statefulset后，可 以做一次扩容来弥补 自己的过失， 新的 pod实 例会运行到与之前完全一致的状态
    
    创建Statefulset实例P296页
    1，创建持久卷
    2，创建控制 Service。在部署一个Statefulset之前，需要创建一 个用于在有状态的 pod之间提供网络标识的headlessService。
    apiVersion: v1
    kind: Service
    metadata:
        name: kubia
    spec:
        clusterIP: None 
        selector:
            app: kubia 
        ports:
        - name: http 
          port: 80
    指定了 clusterIP 为 None, 这就标记了它是 一 个 headless Service 。 它 使得你的 pod 之间可以彼此发现(后续会用到这个功能)。 
    3，创建Statefulset
    apiVersion: apps/v1beta1 
    kind: StatefulSet 
    metadata:
        name: kubia
    spec:
        serviceName: kubia 
        replicas: 2 
        template:
            metadata: 
                labels:
                    app: kubia 
            spec:
                containers:
                - name: kubia
                  image: luksa/kubia-pet 
                  ports:
                  - name: http
                    containerPort: 8080 
                  volumeMounts:
                  - name: data
                    mountPath: /var/data 
        volumeClaimTemplates:   # 依照这个模版为每一个pod创建一个持久卷声明
         - metadata:
          name: data
          spec:
            resources: 
                requests:
                    storage: 1Mi 
            accessModes:
            - ReadWriteOnce
    访问同一个StatefulSet中的pod实例（访问某个指定的pod），创建一个handless service然后通过api服务器访问，要先创建代理(kubectl proxy)，访问localhost:8001/api/v1/namespaces/default/pods/<podname>/proxy/<path>
    普通的service只会吧请求随机分配到pod

    通过DNS各个pod实例伙伴间彼此发现通信，使用dns查询srv记录P305页

# 了解Kubernetes机理
    集群组成：
    • Kubemetes控制平面
        - etcd分布式持久化存储
        - API服务器
        - 调度器
        - 控制器管理器
    • (工作)节点
        - Kubelet
        - Kubelet服务代理( kube-proxy)
        - 容器运行时(Docker、rkt或者其他)

    Kubemetes系统组件间只能通过API服务器通信， 它们之间不会直接通信。
    尽管工作节点上的组件都需要运行在同一 个节点上， 控制平面的组件可以被简 单地分割在多台服务器上。为了保证高可用性， 控制平面的每个组件可以有多个实 例。
    KubemeteAs PI服务器作为中心组件， 其 他组件或者客户端 (如kubectl)都 会去调用它。以RESTfulAPI的形式提供了可以查询、修改集群状态的CRUD(Create、
    Read、 Update、 Delete)接口。 它将状态存储到etcd中。

    调度器 做的就是通过 API 服务器更新 pod 的定义。 然后 API 服务器再去通知 Kubelet(同样， 通过之前描述的监听机制)该 pod 已经被调度过。 当目标节点上的 Kubelet 发现该 pod 被调度到本节点， 它就会创建并且运行 pod 的容器。
    选择节点操作可以分解为两部分所示:
    • 过滤所有节点， 找出能分配给 pod 的可用节点列表。
    • 对可用节点按优先级排序， 找出最优节点。 如果多个节点都有最高的优先级分数， 那么则循环分配，确保平均分配给 pod。

    控制器：API 服务器只做了存储资源到 etcd 和通知客户端有变更的工作。 调度器则只是给 pod 分配节点， 所以需要有活跃的组件确保系统真实状态朝 API 服 务器定义的期望的状态收敛。 这个工作由控制器管理器里的控制器来实现。控制器就是活跃的 Kubernetes组件， 去做具体工作部署资源。总的来说， 控制器执行 一 个 “ 调和 “ 循环， 将实际状态调整为期望状态(在资 源 spec 部分定义)， 然后将新的实际状态写入资源的 s七atus 部分。 
    控制器更新 API服务器的 一个资源后， Kubelet和 Kubemetes Service Proxy(也 不知道控制器的存在)会做它们的工作 ， 例如启动 pod 容器、加载网络存储，或者 就服务而言 ，创建跨 pod 的负载均衡 。

    Kubelet：
    Kubelet就是负责所有运行在工作节点上内容的组件。
    需要持续监控API服 务器是否把该节点分配给 pod， 然后启动 pod 容器 。 具体实现方式是告知配置好的容器运行时’ (Docker、 CoreOS 的 Rkt，或者其他 一 些东西)来从特定容器镜像运行 容器 。 Kubelet 随后持续监控运行的容器，向 API 服务器报告它们的状态、事件和 资源消耗 。

# 保障集群内节点和网络安全
## 在pod中使用宿主节点
     spec:
        containers
        - image: luksa/ kubia
          name: kubia
          ports:
          - containerPort: 8080
            hostPort : 9000 
            protocol : TCP
    创建这个 pod 之后，可以通过它所在节点的 9000 端口访问这个 pod。 主节点时，并不能通过其他宿主节点的同一端口访问该 podo有多个宿

## 在pod中使用宿主节点的 PID 与 IPC 命名空间
    spec:
        hostPID: true
        hostIPC: true 
        containers :
        - name : main
          image: alpine
          command: [”/bin/sleep”,” 999999”]
    pod spec 中的 hostPID 和 hostI PC 选项与 hostNetwork 相似。当它们被设 置为true时， pod中的容器会使用宿主节点的PID和IPC命名空间，分别允许它 们看到宿主机上的全部进程，或通过 IPC机制与它们通信。

# 计算资源管理
    我们创建 一 个 pod 时， 可以指定容器对 CPU 和内存的资源请求量(即 requests), 以及资源限制量(即Lim心)。 它们并不在 pod 里定义， 而是针对每个容器单独指定。pod 对资源的请求量和限制量是它所包含的所有容器的请求量和限制量之和。

    containers:
    - image: busybox
      command: ["dd", "if=/dev/zero", "of=/dev/null"]
      name: main 
      resources:        # 为容器定义资源请求量
        requests:
            cpu: 200m   # 申请200毫核（1/5核）
            memory: 10Mi    申请10MB内存

    调度器在调度时只考虑那些未分配资源量满足pod 需求量 的节点。如果节点的 未分配资源量小千pod 需求量，这 时节点没有能力提供pod对资源需求的最小 量，因此Kubemetes不会将该pod调度到这个节点。

    调度器在调度时并不关注各类资源在当前时刻的 实际使用噩，而只关心节点上部署的所有pod 的资源申请量之和。

    调度器首先会 对节点列表进行过滤， 排除那些不满足需求的节点， 然后根据预先配置的优先级函数对其余节点进行排序 。 其
    中有两个基于资源请求量的优先级排序函数: LeastRequestedPriority和 MostReques七edPriority。 前者优先将pod调度到请求量少的节点上(也就是 拥有更多未分配资源的节点)， 而后者相反， 优先调度到请求量多的节点(拥有更少 未分配资源的节点)。 但是， 正如我们刚刚解释的， 它们都只考虑资源请求量， 而不关注实际使用资源量。

    kubectl describe pods podname #可以查看pod未被调度时的原因

    CPUrequests不仅仅在调度时起作用，它还决定着剩余(未使用) 的CPU时间 如何在pod之间分配。正如图14.2描绘的那样， 因为第一 个pod 请求了200毫核， 另 一 个请求了1000毫核，所以未使用的CPU将按照1:5的比例来划分给这两个 pod。如果两个pod 都全力使用CPU, 第一 个pod 将获得16.7%的CPU时间，另一个将获得83.3%的cpu空间
    另一方面，如果一个容器能够跑满CPU, 而另一个容器在该时段处于空闲状态， 那么前者将可以使用整个CPU时间(当然会减掉第二个容器消耗的少量时间)

    自定义资源
    首先要让Kubemetes知道有哪些自定义资源，通过执行HTTP的PATCH请求来完成。然后，创建 pod 时\只要简单地在容器 spec 的 resources . requests 宇段下， 或者像之前例子那样使用带 一 requests 参数的 kubectl run 命令来指定自定 义资源名称和申请 量 ，调度器就可以确保这个 pod 只能部署到满足自定义资源申 请 量 的节点，同时每个 己部署的 pod 会减少节点的这类可分配资源数量 。

## 限制容器可用资源
    CPU 是一种可压缩资源，意味着我们可 以在不对容器内运行的进程产生不利影 响的 同时，对其使用量进行限制 。 而内存明显不同一一是 一种不可压缩资源。一旦 系统为进程分配了 一 块内存，这块内存在进程主动 释放之前将无法被回收。这就是 我们为什么需要限制容器的最大内存分配量的根本原因。
    如果不对内存进行限制， 工作节点上的容器(或者pod)可能会吃掉所有可用 内存，会对该节点上所有其他pod和任何新调度上来的pod (记住新调度的pod是 基于内存的申请量而不是实际使用 量的)造成影响。
    创建：
    spec:
        containers:
        - image: busybox
          command: ["dd”,”if=/dev/ zero”,”of=/dev/null"] 
          name: main
          resources:
            limits:             # 限制容器的资源使用量
              cpu: 1
              memory: 2OMi
    容器内的进程不允许消耗 超过 1 核 CPU 和 20MB 内存 。
    与资源 requests 不同的是，资源 limits 不受节点可分配资源量的约束 。 所有 limits 的总和允许超过节点资源总量的 100%
    如果节点资源使用 量超过 100% ， 一些容器将被杀掉 ， 这是一个很重要的 结果 。
    对一个进程的 CPU 使用率可以进行限制， 因此当为 一 个容器设置 CPU 限额时，该进程只会分不到比限额更多的 CPU 而己 。
    而内存却有所不同 。 当进程尝试申请分配比限额更多的内存时会被杀掉(我们 会说这个容器被 OOMKilled 了， OOM 是 Out Of Memory 的缩写〉。 如果 pod 的重 启策略为 Always 或 OnFailure，进程将会立即重启，因此用户可能根本察觉不 到它被杀掉。但是如果它继续超限并被杀死， Kubernetes 会再次尝试重启，并开始 增加下次重启的间隔时间 。 

    在容器内看到的始终是节点的 内存， 而不是容器本身的内存。即使你为容器设置了最 大可用内存的限额， top 命令显示 的是运行该容器的节 点的内存数量 ，而容器无法感知到此限制 。
    与内存完全 一 样，无论有没有配置 CPU limits， 容器内也会看到节点所有的 CPU。将 CPU 限额配置为 l ，并不会神奇地只为容器 暴露一 个核。 CPU limits 做的 只是限制容器使用的 CPU 时间 。

## pod QoS等级
    当某个pod突然需要更多资源（和其他pod使用资源之和超过了节点资源总和），那么怎么决定哪些pod能够继续运行呢。
    Kubernetes 将pod 划分为 3 种 QoS 等级 :
    • BestEffort (优先级最低) 
    • Burstable
    • Guaranteed (优先级最高)
    QoS 等级来源 于 pod所包含的容器的资源 requests和 limits 的配置 
    为pod分配等级：
    • 最低优先级的 QoS 等级是 BestEffort。 会分配给那些没有(为任何容器) 设置任何 requests 和 limits 的 pod。
    • Guaranteed 等级，会分配给那些所有资源 request 和 limits相等的 pod。 因为如果容器的 资 源 requests 没有显式设置，默认与 limits 相同，所以 只设 置所有资源(pod 内每个容器的每种资源)的限制量就可以使 pod 的 QoS 等级为 Guaranteed。 
    • Burstable QoS 等 级，其他所有 的 pod都属于这个等级
    pod的QoS等级同样适用于容器QoS等级
    对千多容器pod, 如果所有的容器的QoS等级相同， 那么这个等级就是pod的QoS等级。 如果至少有一个容器的QoS等级与其他不同，无论这个容器是什么等级， 这个pod的QoS等级都是Burstable等级。 

    BestEffort等级的pod首先被杀掉， 其次是 Burstable pod, 最后是Guaranteed pod。 Guaranteedpod只有在系统进程需要内存时才会被杀掉。

    对于同优先级的pod，系统通过比 较所有运行进程的OOM分数来选择要杀掉的进程。 当需要释放内存时， 分数最高的进程将被杀死。OOM分数由两个参数计算得出:进程已消耗内存占可用内存的百分比， 与 一 个基千pod QoS等级和容器 内存申请量固定的OOM分数调节因子。

    设置某个命名空间默认的pod或者容器的资源requests和limits可以创建LimitRange资源
    限制某个命名空间中可用资源的总量，可以创建ResourceQuota对象

    LimitRange应用于单独的pod ; ResourceQuota应用千命名空间中所有的pod

    创建了ResourceQuota后，必须在创建pod是指定相关资源的request和limit量，或者创建LimitRange对象

## 收集、获取实际资源使用情况
    Kubelet 自身就 包含了一个名为 cAdvisor 的 agent，它会收集整个节点和节点上运行的所有单独容 器的资源消耗情况 。 集中统计整个集群的监控信息需要运行 一个叫作 Heapster 的附 加组件。
    Heapster 以 pod 的方式运行在某个节点上，它通过普通的 Kubernetes Service 暴 露服务，使外部可以通过一个稳定的 IP 地址访问 。 它从集群中所有的 cAdvisor 收 集数据，然后通过一个单独的地址暴露 。
    启用Heapster：
    minikube addons enable heapster
    kubectl top node   # 显示该节点的cpu和内存使用情况
    kubectl top pod     # 显示该节点下每个pod的cpu和资源的使用情况

    cAdvisor和Heapster 都只保存一个很短时间窗的资源使用量数据。 如果需要分析一段时间的pod的资 源使用情况， 必须使用额外的工具。如果是本地 Kubemetes 集群，往往使用Influx.DB来存储统计数据， 然后使用 Grafana 对数据进行可视化和分析。（启用Heapster 插件会自动部署上面两个东西）

    kubectl cluster-info # 查找Grafanaweb控制台的URL。
    使用minikube时， Grafana的web控制台通过NodePort Service暴露， 因此 我们使用以下命令在浏览器中将其打开:
    minikube service monitoring-grafana -n kube-system

# 自动横向伸缩pod与集群节点 hpa
    Kubemetes可以监控你的pod, 并在检测到CPU使用率或其他度量 增长时自动对它们扩容。
## 横向自动伸缩
    横向pod自动伸缩是指由控制器管理的pod副本数量的自动伸缩。 
    它由 Horizontal控制器执行， 我 们通过创建 一 个HorizontalpodAutoscaler C HPA)资源来启用和配置Horizontal控制器。 该控制器周期性检查pod度量， 计算满足HPA 资源所配置的目标 数 值所 需 的 副本数量， 进而调整目标资源(如Deployment、 ReplicaSet、 ReplicationController、 StatefulSet等)的replicas字段。
    自动伸缩 的过程可以分为三个步骤:
    • 获取被伸缩资源对象所管理的所有pod度量。
    • 计算使度量数值到达(或接近)所指定目标数值所需的pod数量。
    • 更新被伸缩资源的rep巨cas字段。

    获取度量总量通过前一章的Heapster，基于多个pod度量的自动伸缩(例如: CPU使用率和每秒查询率[QPS])的计 算也并不复杂。 Autoscaler单独计算每个度量的副本数， 然后取最大值(例如:如 果需要4个pod达到目标CPU使用率， 以及需要3个pod来达到目标QPS, 那么Autoscaler 将扩展到4个pod)。当Autoscaler配置为只考虑单个度量时， 计算所需 副本数很简单。 只要将所有pod的度量求和后除以HPA资源上配置的目标值， 再向上取整即可。 
    计算度量cpu用量/cpu请求量
    就 Autoscaler而言， 只有pod的保证CPU用量(CPU请求)才与确认pod的 CPU使用有关。 Autoscaler对比pod的实际CPU使用与它的请求， 这意味着你需要 给被伸缩的pod设置CPU请求，

    创建HPA（基于cpu自动调整）：
    kubectl autoscale deployment(资源类型) kubia(资源实例名称) --cpu-percent=30 --min=1 --max=5
    设置了pod的目标 CPU使用率为30%, 指定了副本的最小和最大数量。
    一定要确保自动伸缩的目标是 Deployinent 而不是底层的 ReplicaSet。

    kubectl get hpa
    Autoscaler只会在Deployment上调节预 期的 副 本 数 量。接下来由 Deployment控制器负责更新ReplicaSet对象上的副本数量，从而使ReplicaSet控制 器删除多余的两个pod而留下 一 个。

    如果当 前副本数大于 2, Autoscaler 单次操作至 多 使副本数翻倍;如果副本数只有 1 或 2, Autoscaler 最多扩 容到 4个副本。
    另外 ， Autoscaler 两次扩容操作之间的时间间隔 也 有限制 。 目前，只有当 3 分 钟内没有任何伸缩操作时才会触发扩容，缩容操作频率更低一-5分钟。

    创建HPA（基于内存自动调整）：
    基于内存的自动伸缩比基于CPU的困难很多。 主要原因在于，扩容之后原有的 pod 需要有办法释放内存 。 这只能由应用完成，系统无法代芳 。 系统所能做的只有 杀死并重启应用，希望它能比之前少占用 一些内存;但如果应用使用了跟之前一样 多的内存 ， Autoscaler 就会扩容、扩容 ， 再扩容 ， 直 到达到 HPA 资源上配置的最大 pod 数量 。 显然没有 人想 要这种行为 。 基于内存使用的自动伸缩在 Kubernetes 1.8 中 得到支持，配置方法与基于 CPU 的自动伸缩完全相同 。具体使用方式留作读者练习 。

    基于其他自定义度量进行自动伸缩：
    P454页，总共有Resource、pods、Object度量类型

## 纵向扩容
    增加分配给pod的cpu、内存等资源

# 高级调度
    Kubernetes允许你去影响pod被调度到哪个节点。 

## 使用污点和容忍度阻止节点调度到特定节点
    这些特性被用于限制哪些pod可以被调度到某一个节点。 只有当 一个pod容忍某个 节点的污点， 这个pod才能被调度到该节点。
    节点 选择器和节点亲缘性规则，是通过明确的 在pod中添加 的 信 息，来决定 一 个pod可 以或者不可以被调度到哪些节点上。而污点则是在不修改巳有pod信息的前提下， 通过在节点上添加污点信息，来拒绝pod在某些节点上的部署。

    通过kubectl describe node查看节点的污点信息
    污点包含了一个key、value, 以及一个effect, 表现为<key>=<value>:<effect>
    Taints:node-role.kubernetes.io/master:NoSchedule, 包含一 个为node­-role.kubernetes.io/master的key, 一个空的value, 以及值为NoSchedule的effect，这个污点将阻止pod调度到这个节点上面，除非有pod能容忍这个污点

    通过kubectl describe pops podname 显示pod的污点容忍度
    Tolerations: node-role.kubernetes.io/master=:NoSchedule

    • NoSchedule 表示如果 pod 没有容忍这些污点， pod 则不能被调度到包含 这些污点的节点上。
    • PreferNoSchedule 是 NoSchedule 的一个宽松的版本， 表示尽量阻止 pod 被调度到这个节点上， 但是如果没有其他节点可以调度， pod 依然会被调度到这个节点上。
    • NoExecute 不同于 NoSchedule 以及 PreferNoSchedule, 后两者只在 调度期间起作用， 而 NoExecute 也会影响正在节点上运行着的 pod。 如果 在一个节点上添加了 NoExecute 污点， 那些在该节点上运行着的 pod, 如 果没有容忍这个 NoExecute 污点， 将会从这个节点去除。

    添加污点：
    kubectl taint node nodel.k8s node-type=production:NoSchedule

    在pod上添加污点容忍度：
    spec:
        tolerations:
        - key: node-type
          operator: Equal 
          value: production 
          effect: NoSchedule

    节点可以拥有多个污点信息，而pod也可以有多个污点容忍度。正如你所见，污点可以只有一个key和 一 个效果，而不必设置value。污点容忍度可以通过设置Equal操作符Equal操作符来指定匹配的value (默认情况下的操作符)，或者也可以通过设置Exists操作符来匹配污点的key。
## 使用节点亲缘性将pod调度到特定节点上
    topologyKey字段决定了pod不能被调度的范围。
    pod亲缘性的topologyKey表示了被调度的pod和另一个pod的距离(在 同一个节点、 同一个 机柜、 同 一个 可用性局域或者可用性地域)。

    污点可以用来让 pod 远离特定的几点。 节点亲缘性 (node affinity), 这种机制允许你通知 Kubemetes 将 pod 只调度到某个几点子集上面。
    它更倾向于调度到某些节点上， 之后 Kubemetes 将尽量把这个 pod 调度到这些节点上面。 如果没法实现的话， pod 将被调度到其他某个节点上。
    节点亲缘性根据节点的标签来进行选择， 这点跟节点选择器是一 致的。
    spec:
      affinity: 
        nodeAffinity:
          requiredDuringSchedulingignoredDuringExecution:    # 只影响正在调度的pod
            nodeSelectorTerms :
            - matchExpressions: 
              - key: gpu
                operator: In 
                values:
                - "true"
    节点亲缘性的最大好处就是，当调度某 一个 pod 时，指定调度器可 以优先考虑哪些节点，这个功能是通过 preferredDuringSchedulingignored DuringExecution 宇段来实现的 。

# 使用pod亲缘性与非亲缘性对pod进行协同部署
    指定pod自身之间的亲缘性（来使多个pod部署到相近位置等）。
    template:
      spec:
        affinity: 
        podAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:    # 只影响正在调度的pod
          - topologyKey: kubernetes.io/hostname
            labelSelector:
              matchLabels:
                app: backend
    要求pod将被调度到和其他包含app=backend标 签 的pod所在的相同节点上(通过topologyKey字段指定)
    调度器首先找出所有匹配前端 pod 的 podAffinity 配置中 labelSelector 的 pod, 之后将前端 pod 调度到相同的节点上。

    利用pod非亲缘性分开调度pod
    将pod彼此远离
    它和pod亲缘 性的表示方式 一样， 只不过是将podAffin江y字段换成podAntiAffin江y, 这将 导致调度器永远不会选择那些有包含podAn巨Affinity匹配标签的pod所在的节 点
    template:
      spec:
        affinity: 
        podAntiAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:    # 只影响正在调度的pod
          - topologyKey: kubernetes.io/hostname
            labelSelector:
              matchLabels:
                app: frontend
    该pod自身标签为app: frontend，设置了不要调度到运行有pod标签为app: frontend的节点，所以每个pod会被调度到不同节点

# 开发应用的最佳实践
     但是你可以阻止一 个主容器的启动，直到它的预置条件被满足。这个是通过在 pod中包含 一 个叫作init的容器来实现的
     一个pod可以拥有任意数量的init容器。init容器是顺序执行的， 并且仅当最后 一 个init容器执行完毕才会去启动主容器。
     创建：
    spec:
      initContainers:
      - name: init 
        image: busybox 
        command:
        - sh
        - -c
        - 'while true; do echo "Waiting for fortune service to come up..."; wget http://fortune -q -T 1 -O /dev/null >/dev/null 2>/dev/null && break; sleep 1; done; echo "Service is up! Starting main container."'
    创建init容器，循环执行脚本直到执行完毕

    pod还有启动后和停止前两个生命周期钩子函数
    生命周期的钩子是基于每个容器来指定的， 和init容器 不同的是 ，init容器 是应用到整个pod。 
    启动后钩子是在容器的主进程启动之后立即执行的（不是主进程初始化完毕才执行）。可以用它在应用启动时做 一些额外的工作。
    在钩子执行执行完毕之前，容器会一直停留在Waiting状态，如果钩子执行失败或者返回了非0的状态码，主容器会被杀死。
    spec:
      containers:
      - image: luksa/kubia
        name:kubia
        lifecycle:
          postStart:        # 启动后钩子
            exec:
             command:
             - sh
             - -c
             - "echo 'hook will fail with exit code 15'; sleep 5; exit 15"
    他会在容器启动时执行/bin目录下的postStart.sh脚本

    停止前钩子是在容器被终止之前立即执行的。 当一个容器需要终止运行的时候， Kubelet在配置了停止前钩子的时候就会执行这个停止前钩子， 并且仅在执行完钩子程序后才会向容器进程发送SIGTERM信号
    lifecycle:
      preStop:          # 停止前钩子
        httpGet:        # 执行http get请求
          port: 8080 
          path: shutdown
    这个代码清单中定义的停止前钩子在Kuble et开始终止容器的时候就立即执行 到 http://podIP:8080/ shut down的HTTPGET请求。
    默认情况下，host的值是pod的IP地址。
    和启动后钩子不同的是，无论钩子执行是否成功容器都会被终止。

## 不用将镜像推送到docker hub而构建pod
    在Minikube VM中使用DockerDaemon来构建镜像
    如果你正在使用Minikube开发应用，并且计划在每个更改之后都构建 一个镜像， 可以在MinikubeVM中使用DockerDaemon来进行镜像构建， 而不是通过本地的 DockerDaemon构建然后再推送到镜像中心， 最后拉取到MinikubeVM中。 为了使 用Minikube的DockerDaemon, 只需要将你的DOCKER_HOST环境变量指向它。 幸运的是， 这个做起来实际上比听上去容易多了， 只需要在本地机器上运行下面的命 令:
    eval ${minikube docker-env)
    这个命令会帮你设置所有需要的环境变量， 然后你就可以像DockerDaemon运 行在你本地的时候那样构建镜像了。 构建完镜像之后， 不需要再去推送镜像， 因为 它已经存储在MinikubeVM中了， 这样新的pod就可以立即使用这个镜像了。 如果你的pod已经在运行了， 那么可以删除它们或者杀死容器让它们重启。

    在本地构建镜像然后直接复制到Minikube VM中
    如果你无法使用MinikubeVM内部的DockerDaemon来构建镜像， 这里仍然有 方法来避免将镜像推送到镜像中心， 然后使用运行在MinikubeVM内部的Kubelet拉取镜像这样的流程。 如果你在本地机器构建好了镜像， 可以使用下面的命令将镜 像直接复制到MinikubeVM中:
    docker save <image> ｜ (eval $(minikube docker-env) && docker load)
    和之前一 样， 这个 镜像也可以在 pod中立即使用了。 这里注意确保podspec 中 的imagePullPo巨cy不要设置为Always, 因为这会导致 从外部镜像中心拉取镜 像， 从而导致你复制过去的镜像的更改丢失。

# Kubernetes 扩展

## 自定义资源对象 CRD
    创建资源对象：
    apiVersion: apiextensions.k8s.io/v1beta1 
    kind : CustomResourceDefinition
      metadata :websites. extensions. example. com  # 自定义对象的全名
    spec:
      scope: Namespaced     # 命名空间作用域
    group: extensions.example.com
    version: v1     # 版本
    names:          # 指定自定义对象名称的各种形式
      kind: Website 
      singular: website 
      plural: websites

    创建实例：
    kind: Website 
    metadata:
     name: kubia 
    spec:
      gitRepo: https://github com/ luksa/kubia-website-example.gi

## 注册自定义 API 服务器
    要将自定义API服务器 添加到集群中，可以将其部署为一 个pod并通过 Service暴露。

# 使用 Kubernetes服务目录扩展Kubernetes
    服务目录就是列出所有服务的目录。 用户可以浏览目录并自行设置 目录中列出的服务实例， 却无须处理服务运行所需的Pod、 Service、 ConfigMap和 其他资源。 
    • 一个ClusterServiceBroker, 描述一个可以提供服务的(外部)系统
    • 一个ClusterServiceClass, 描述 一个可供应的服务类型
    • 一个Servicelnstance, 已配置服务的一个实例
    • 一个ServiceBinding, 表示 一 组客户端(pod)和Servicelnstance之间的绑定

    集 群 管 理 员 为 会 每 个 服 务 代 理 创 建 一 个 C l u s t e r S e r v i c e Br o k e r 资 源 ， 而这些服务代理需要在集群中提供它们的服务。 接着， Kubemetes从服务代理获 取它可以提供的服务列表， 并为它们中的每个服务创建一个ClusterServiceClass资源。 当用户调配服务时， 首先需要创建 一 个Servicelnstance资源， 然后创建 一 个 ServiceBinding以将 该Servicelnstance 绑定到它们的pod。 下一 步， 这些pod 会被注 入一 个Secret, 该Secret 包含连接到配置的Servicelnstance 所需的凭证和其他 数据。
    
    与核心Kubemetes类似的是， 服务目录也是由三个组件组成的分布式系统:
    • 服务目录API 服务器
    • 作为存储的etcd
    • 运行所有控制器的控制器管理器

# Pod 优先级抢占
    要使用优先级和抢占：
    1.新增一个或多个 PriorityClass。
    2.创建 Pod，并将其 priorityClassName 设置为新增的 PriorityClass。 当然你不需要直接创建 Pod；通常，你将会添加 priorityClassName 到集合对象（如 Deployment） 的 Pod 模板中。

    PriorityClass 是一个无名称空间对象，它定义了从优先级类名称到优先级整数值的映射。 名称在 PriorityClass 对象元数据的 name 字段中指定。 值在必填的 value 字段中指定。值越大，优先级越高。 PriorityClass 对象的名称必须是有效的 DNS 子域名， 并且它不能以 system- 为前缀。

    PriorityClass 还有两个可选字段：globalDefault 和 description。 globalDefault 字段表示这个 PriorityClass 的值应该用于没有 priorityClassName 的 Pod。 系统中只能存在一个 globalDefault 设置为 true 的 PriorityClass。 如果不存在设置了 globalDefault 的 PriorityClass， 则没有 priorityClassName 的 Pod 的优先级为零。

    创建：
    apiVersion: scheduling.k8s.io/v1
    kind: PriorityClass
    metadata:
        name: high-priority
    value: 1000000
    globalDefault: false
    description: "此优先级类应仅用于 XYZ 服务 Pod。"

    创建使用该优先级的pod：
    apiVersion: v1
    kind: Pod
    metadata:
    name: nginx
    labels:
        env: test
    spec:
    containers:
    - name: nginx
        image: nginx
        imagePullPolicy: IfNotPresent
    priorityClassName: high-priority

    当启用 Pod 优先级时，调度程序会按优先级对悬决 Pod 进行排序， 并且每个悬决的 Pod 会被放置在调度队列中其他优先级较低的悬决 Pod 之前。 因此，如果满足调度要求，较高优先级的 Pod 可能会比具有较低优先级的 Pod 更早调度。 如果无法调度此类 Pod，调度程序将继续并尝试调度其他较低优先级的 Pod。

    当 Pod P 抢占节点 N 上的一个或多个 Pod 时， Pod P 状态的 nominatedNodeName 字段被设置为节点 N 的名称。 该字段帮助调度程序跟踪为 Pod P 保留的资源，并为用户提供有关其集群中抢占的信息。

    请注意，Pod P 不一定会调度到“被提名的节点（Nominated Node）”。 在 Pod 因抢占而牺牲时，它们将获得体面终止期。 如果调度程序正在等待牺牲者 Pod 终止时另一个节点变得可用， 则调度程序将使用另一个节点来调度 Pod P。 因此，Pod 规约中的 nominatedNodeName 和 nodeName 并不总是相同。 此外，如果调度程序抢占节点 N 上的 Pod，但随后比 Pod P 更高优先级的 Pod 到达， 则调度程序可能会将节点 N 分配给新的更高优先级的 Pod。 在这种情况下，调度程序会清除 Pod P 的 nominatedNodeName。 通过这样做，调度程序使 Pod P 有资格抢占另一个节点上的 Pod。

    当有多个节点可供执行抢占操作时，调度器会尝试选择具有一组优先级最低的 Pod 的节点。 




## 非抢占式 PriorityClass
    配置了 PreemptionPolicy: Never 的 Pod 将被放置在调度队列中较低优先级 Pod 之前， 但它们不能抢占其他 Pod。等待调度的非抢占式 Pod 将留在调度队列中，直到有足够的可用资源， 它才可以被调度。非抢占式 Pod，像其他 Pod 一样，受调度程序回退的影响。 这意味着如果调度程序尝试这些 Pod 并且无法调度它们，它们将以更低的频率被重试， 从而允许其他优先级较低的 Pod 排在它们之前。
    非抢占式 Pod 仍可能被其他高优先级 Pod 抢占。
    PreemptionPolicy 默认为 PreemptLowerPriority， 这将允许该 PriorityClass 的 Pod 抢占较低优先级的 Pod（现有默认行为也是如此）。 如果 PreemptionPolicy 设置为 Never，则该 PriorityClass 中的 Pod 将是非抢占式的。
    创建：
    apiVersion: scheduling.k8s.io/v1
    kind: PriorityClass
    metadata:
        name: high-priority-nonpreempting
    value: 1000000
    preemptionPolicy: Never
    globalDefault: false
    description: "This priority class will not cause other pods to be preempted."

## 被抢占牺牲者的体面终止
    当 Pod 被抢占时，牺牲者会得到他们的 体面终止期。 它们可以在体面终止期内完成工作并退出。如果它们不这样做就会被杀死。 这个体面终止期在调度程序抢占 Pod 的时间点和待处理的 Pod (P) 可以在节点 (N) 上调度的时间点之间划分出了一个时间跨度。 同时，调度器会继续调度其他待处理的 Pod。当牺牲者退出或被终止时， 调度程序会尝试在待处理队列中调度 Pod。 因此，调度器抢占牺牲者的时间点与 Pod P 被调度的时间点之间通常存在时间间隔。 为了最小化这个差距，可以将低优先级 Pod 的体面终止时间设置为零或一个小数字。

### 与低优先级 Pod 之间的 Pod 间亲和性
     “如果从此节点上删除优先级低于悬决 Pod 的所有 Pod，悬决 Pod 是否可以在该节点上调度？”
     如果悬决 Pod 与节点上的一个或多个较低优先级 Pod 具有 Pod 间亲和性， 则在没有这些较低优先级 Pod 的情况下，无法满足 Pod 间亲和性规则。 在这种情况下，调度程序不会抢占节点上的任何 Pod。 相反，它寻找另一个节点。调度程序可能会找到合适的节点， 也可能不会。无法保证悬决 Pod 可以被调度。

## 节点压力驱逐
    节点压力驱逐是 kubelet 主动终止 Pod 以回收节点上资源的过程。
    kubelet 监控集群节点的 CPU、内存、磁盘空间和文件系统的 inode 等资源。 当这些资源中的一个或者多个达到特定的消耗水平， kubelet 可以主动地使节点上一个或者多个 Pod 失效，以回收资源防止饥饿。

    在节点压力驱逐期间，kubelet 将所选 Pod 的 PodPhase 设置为 Failed。这将终止 Pod。

    如果 Pod 是由替换失败 Pod 的工作负载资源 （例如 StatefulSet 或者 Deployment）管理， 则控制平面或 kube-controller-manager 会创建新的 Pod 来代替被驱逐的 Pod。

    kubelet 使用各种参数来做出驱逐决定，如下所示：
    ·驱逐信号
    ·驱逐条件
    ·监控间隔

    驱逐信号是特定资源在特定时间点的当前状态。 kubelet 使用驱逐信号，通过将信号与驱逐条件进行比较来做出驱逐决定， 驱逐条件是节点上应该可用资源的最小量。

# 调度框架
    调度框架定义了一些扩展点。调度器插件注册后在一个或多个扩展点处被调用。 这些插件中的一些可以改变调度决策，而另一些仅用于提供信息。

    每次调度一个 Pod 的尝试都分为两个阶段，即 调度周期 和 绑定周期。

1. 队列排序
    队列排序插件用于对调度队列中的 Pod 进行排序。 队列排序插件本质上提供 less(Pod1, Pod2) 函数。 一次只能启动一个队列插件。

2. 前置过滤
    前置过滤插件用于预处理 Pod 的相关信息，或者检查集群或 Pod 必须满足的某些条件。 如果 PreFilter 插件返回错误，则调度周期将终止。
3. 过滤
过滤插件用于过滤出不能运行该 Pod 的节点。对于每个节点， 调度器将按照其配置顺序调用这些过滤插件。如果任何过滤插件将节点标记为不可行， 则不会为该节点调用剩下的过滤插件。节点可以被同时进行评估。
4. 后置过滤
这些插件在筛选阶段后调用，但仅在该 Pod 没有可行的节点时调用。 插件按其配置的顺序调用。如果任何后过滤器插件标记节点为“可调度”， 则其余的插件不会调用。典型的后筛选实现是抢占，试图通过抢占其他 Pod 的资源使该 Pod 可以调度。
5. 前置评分
前置评分插件用于执行 “前置评分” 工作，即生成一个可共享状态供评分插件使用。 如果 PreScore 插件返回错误，则调度周期将终止。
6. 评分
评分插件用于对通过过滤阶段的节点进行排名。调度器将为每个节点调用每个评分插件。 将有一个定义明确的整数范围，代表最小和最大分数。 在标准化评分阶段之后，调度器将根据配置的插件权重 合并所有插件的节点分数。
7. 标准化评分
标准化评分插件用于在调度器计算节点的排名之前修改分数。 在此扩展点注册的插件将使用同一插件的评分 结果被调用。 每个插件在每个调度周期调用一次。
8. Reserve
Reserve 是一个信息性的扩展点。 管理运行时状态的插件（也成为“有状态插件”）应该使用此扩展点，以便 调度器在节点给指定 Pod 预留了资源时能够通知该插件。 这是在调度器真正将 Pod 绑定到节点之前发生的，并且它存在是为了防止 在调度器等待绑定成功时发生竞争情况。
9. Permit
Permit 插件在每个 Pod 调度周期的最后调用，用于防止或延迟 Pod 的绑定。 一个允许插件可以做以下三件事之一：
- 批准：一旦所有 Permit 插件批准 Pod 后，该 Pod 将被发送以进行绑定。
- 拒绝：如果任何 Permit 插件拒绝 Pod，则该 Pod 将被返回到调度队列。 这将触发Unreserve 插件。
- 等待（带有超时）：如果一个 Permit 插件返回 “等待” 结果，则 Pod 将保持在一个内部的 “等待中” 的 Pod 列表，同时该 Pod 的绑定周期启动时即直接阻塞直到得到 批准。如果超时发生，等待 变成 拒绝，并且 Pod 将返回调度队列，从而触发 Unreserve 插件。
10. 预绑定
预绑定插件用于执行 Pod 绑定前所需的任何工作。 例如，一个预绑定插件可能需要提供网络卷并且在允许 Pod 运行在该节点之前 将其挂载到目标节点上。
如果任何 PreBind 插件返回错误，则 Pod 将被拒绝 并且 退回到调度队列中。
11. Bind
Bind 插件用于将 Pod 绑定到节点上。直到所有的 PreBind 插件都完成，Bind 插件才会被调用。 各绑定插件按照配置顺序被调用。绑定插件可以选择是否处理指定的 Pod。 如果绑定插件选择处理 Pod，剩余的绑定插件将被跳过。
12. 绑定后
这是个信息性的扩展点。 绑定后插件在 Pod 成功绑定后被调用。这是绑定周期的结尾，可用于清理相关的资源。
13. Unreserve
这是个信息性的扩展点。 如果 Pod 被保留，然后在后面的阶段中被拒绝，则 Unreserve 插件将被通知。 Unreserve 插件应该清楚保留 Pod 的相关状态。
使用此扩展点的插件通常也使用Reserve。



    加上上面的 sched.scheduleOne 函数，3个子队列整体的工作流程就是：

    每隔1秒，检测 backoffQ 里是否有 Pod 可以被放进 activeQ 里

    每隔30秒，检测 unscheduleodQ 里是否有 Pod 可以被放进 activeQ 里(默认条件是等待时间超过60秒)

    不停的调用 scheduleOne 方法，从 activeQ 里弹出 Pod 进行调度

    如果一个 Pod 调度失败了，正常就是不可调度的，应该放入 unscheduleableQ 队列。如果集群内的资源状态一直不发生变化，这种情况，每隔60s这些 Pod 还是会被重新尝试调度一次。

    但是一旦资源的状态发生了变化，这些不可调度的 Pod 就很可能可以被调度了，也就是 unscheduleableQ 中的 Pod 应该放进 backoffQ 里面去了。等待安排重新调度，backoffQ 里的 Pod 会根据重试的次数设定等待重试的时间，重试的次数越少，等待重新调度的时间也就越少。backOffQ 里的 Pod 调度的速度会比 unscheduleableQ 里的 Pod 快得多。










