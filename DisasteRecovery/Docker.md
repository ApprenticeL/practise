1. Dockerfile构建镜像
    
    ```
    # 用 FROM 来表示构建所基于的镜像（镜像是可以复用的），本处使用golang官方镜像
    FROM golang:1.12-alpine
    # WORKDIR 行表示具体在哪个目录下进行构建
    WORKDIR /go/scr/app
    # COPY 将宿主机上的文件拷贝到容器镜像内
    COPY ...
    # RUN 行表示在具体的文件系统内执行相应的动作，例如：
    RUN go get -d -v ./...
    RUN go install -v ./...
    # CMD 行表示使用镜像时的默认程序名字
    CMD ["app"]
    
    # 有了 Dockerfile 之后，就可以通过 docker build 命令构建出所需要的应用
    ```
    
    
2. 运行容器
    
    ```
    # 从docker registry下载镜像
    docker pull busybox:1.25
    # 查看本地镜像
    docker images
    # 选择相应的镜像并运行
    docker run [-d] -name demo busybox:1.25 top
    ```
