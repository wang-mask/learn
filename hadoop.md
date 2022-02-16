# Hadoop
  组成部分 common、HDFS（分布式文件系统）、YARN（调度器）、MapReduce（负责计算）
  HDFS：NameNode（存储文件的元数据以及文件所在的DataNode），DataNode存储文件的节点，2nn每隔一段时间对NameNode元数据进行备份
  YARN：ResourceManager（管理整个集群的资源），NodeManager（管理单个节点资源），ApplicationMaster（管理单个任务），Container
  MapReduce：Map分配任务，Reduce汇总任务