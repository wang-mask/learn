# Hadoop
  组成部分 common、HDFS（分布式文件系统）、YARN（调度器）、MapReduce（负责计算）
  HDFS：NameNode（存储文件的元数据以及文件所在的DataNode），DataNode存储文件的节点，2nn每隔一段时间对NameNode元数据进行备份
  YARN：ResourceManager（管理整个集群的资源），NodeManager（管理单个节点资源），ApplicationMaster（管理单个任务），Container
  MapReduce：Map分配任务，Reduce汇总任务
  HDFS（分布式文件系统）:NameNode（管理者）、SecondaryNameNode（辅助管理者）、DataNode（工作者）
  MapReduce：海量数据的计算系统
  Yarn（集群资源管理、调度框架）：ResourceManage（管理者）、NodeManager（工作者）

## 快速恢复集群
  1.停止集群 stop_dfs.sh
  2.删除所有集群的hadoop文件夹下的data和logs文件下
  3.hdfs namenode -format
  4.启动集群
  每次启动的集群，namenode的id都有对应的datenode的id，如果namenode挂了之后，重置了一个namenode那他就找不到对应的datenode，所以需要吧data和logs文件夹删除掉

## 常用命令
  集群启停：start-dfs.sh/stop-dfs.sh
  启停Yarn：start-yarn.sh/stop-yarn.sh
  各个服务组件逐一启停：hdfs --daemon start/stop namenode/datanode/secondarynamenode （在哪个节点挂掉了，就可以把它单独再启动）
  启停Yarn：yarn --daemon start/stop  resourcemanager/nodemanager 

## 常用端口
  hadoop3.x：
  HDFS NameNode 内部通信端口：8020/9000/9820
  HDFS NameNode 对用户的查询端口：9870
  Yarn查看任务运行情况的：8088
  历史服务器：19888

## 常用配置
  hadoop3.x：
  core-site.xml
  hdfs-site.xml
  yarn-site.xml
  mapred-site.xml
  workers

## 优缺点
  优点：高容错性（多副本）、可构建在廉价机器上
  缺点：不适合低延迟数据访问、无法搞笑对大量小文件进行存储（nameNode容量有限）、不支持并发写入、仅支持数据的追加（不支持随机修改）
## HDFS块大小设置
  过小：浪费寻址时间，
  过大：计算时浪费传输时间
  HDFS块大小的设置依据：1s磁盘的传输速率
  HDFS中小于一个块大小的文件不会占据整个块的空间（当一个1MB的文件存储在一个128MB的块中时，文件只使用1MB的磁盘空间，而不是128MB）

## 命令
  hadoop fs -command
  hadoop fs -help command
  与linux相同的命令：mkdir、ls、cat、chmod、chown、cp、mv、rm -r、tail、
  上传：
    hadoop fs -put [sorce] [target]
    追加：hadoop fs -appendToFile [source] [target]
  
  下载：
    hadoop fs -get [source] [localTargrt]

  查看文件夹大小
    hadoop fs -du -s -h 文件夹
  查看文件夹下各个文件大小
    hadoop fs -du -h 文件夹
  
## API操作
    // 1 获取配置对象
    Configuration configuration = new Configuration();
    //2 获取文件系统
    // FileSystem fs = FileSystem.get(new URI("hdfs://hadoop102:8020"), configuration);
    FileSystem fs = FileSystem.get(new URI("hdfs://hadoop102:8020"), configuration,"atguigu");
    // 3 操作 创建目录
    fs.mkdirs(new Path("/xiyou/huaguoshan/"));

    // 4 关闭资源
    fs.close();

    创建文件夹
    fs.mkdirs(new Path("/xiyou/huaguoshan1"));
    上传文件：
    fs.copyFromLocalFile(false, true, new Path("D:\\sunwukong.txt"), new Path("hdfs://hadoop102/xiyou/huaguoshan"));
    下载文件
    fs.copyToLocalFile(false, new Path("hdfs://hadoop102/a.txt"), new Path("D:\\"), false);
    删除文件
    fs.delete(new Path("/jdk-8u212-linux-x64.tar.gz"),false);
    重命名
    fs.rename(new Path("/input"), new Path("/output"));
    查看文件夹下文件信息
    获取所有文件信息
      RemoteIterator<LocatedFileStatus> listFiles = fs.listFiles(new Path("/"), true);
      // 遍历文件
      while (listFiles.hasNext()) {
          LocatedFileStatus fileStatus = listFiles.next();
          System.out.println("==========" + fileStatus.getPath() + "=========");
          System.out.println(fileStatus.getPermission());
          System.out.println(fileStatus.getOwner());
          System.out.println(fileStatus.getGroup());
          System.out.println(fileStatus.getLen());
          System.out.println(fileStatus.getModificationTime());
          System.out.println(fileStatus.getReplication());
          System.out.println(fileStatus.getBlockSize());
          System.out.println(fileStatus.getPath().getName());
          // 获取块信息
          BlockLocation[] blockLocations = fileStatus.getBlockLocations();
          System.out.println(Arrays.toString(blockLocations));
          // 判断是否是文件
          if (status.isFile()) {
                System.out.println("文件：" + status.getPath().getName());
            } else {
                System.out.println("目录：" + status.getPath().getName());
            }
      }
