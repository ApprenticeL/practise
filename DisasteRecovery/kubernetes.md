## k8s核心  
#### 核心功能
    1. 调度：Kubernetes 可以把用户提交的容器放到 Kubernetes 管理的集群的某一台节点上去。Kubernetes 会观察正在被调度的这个容器的大小、规格。
    
    1. 调度：Kubernetes 可以把用户提交的容器放到 Kubernetes 管理的集群的某一台节点上去。Kubernetes 会观察正在被调度的这个容器的大小、规格。
    2. 
    3. 自动修复：Kubernetes 对集群中所有的宿主机进行监测，并把运行在失败节点上的容器自动迁移到一个正在健康运行的宿主机上，来完成集群内容器的自动恢复。
    4. 水平伸缩：Kubernetes 有业务负载检查的能力，它会监测业务上所承担的负载，如果这个业务本身的 CPU 利用率过高，或者响应时间过长，它可以对这个业务进行一次扩容。
    
  2. 架构
    ![image](https://user-images.githubusercontent.com/111548435/185942709-e1e64cfc-2aea-44b2-8af1-71b89e6363f0.png)
    1. API Server：顾名思义是用来处理 API 操作的，Kubernetes 中所有的组件都会和 API Server 进行连接，组件与组件之间一般不进行独立的连接，都依赖于 API Server 进行消息的传送；
    2. Controller：是控制器，它用来完成对集群状态的一些管理。比如刚刚我们提到的两个例子之中，第一个自动对容器进行修复、第二个自动进行水平扩张，都是由 Kubernetes 中的 Controller 来进行完成的；
    3. Scheduler：是调度器，“调度器”顾名思义就是完成调度的操作，就是我们刚才介绍的第一个例子中，把一个用户提交的 Container，依据它对 CPU、对 memory 请求大小，找一台合适的节点，进行放置；
    4. etcd：是一个分布式的一个存储系统，API Server 中所需要的这些原信息都被放置在 etcd 中，etcd 本身是一个高可用系统，通过 etcd 保证整个 Kubernetes 的 Master 组件的高可用性。
  
