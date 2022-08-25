1. Dockerfile构建镜像
    
    ``` bash
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
    
    ```bash
    # 从docker registry下载镜像
    docker pull busybox:1.25
    # 查看本地镜像
    docker images
    # 选择相应的镜像并运行
    docker run [-d] -name demo busybox:1.25 top
    ```

## Docker Hello World
    ```
    # 使用 docker run 命令来在容器内运行一个应用程序
    # docker: Docker 的二进制执行文件。
    # run:与前面的 docker 组合来运行一个容器。
    # ubuntu:15.10指定要运行的镜像，Docker首先从本地主机上查找镜像是否存在，如果不存在，Docker 就会从镜像仓库 Docker Hub 下载公共镜像。
    # /bin/echo "Hello world": 在启动的容器里执行的命令
    
    docker run ubuntu:15.10 /bin/echo "Hello world"
    ```
    
## 运行交互式的容器
    ```
    # 通过docker的两个参数 -i -t，让docker运行的容器实现"对话"的能力
    # -t:在新容器内指定一个伪终端或终端。
    # -i:允许你对容器内的标准输入 (STDIN) 进行交互。
    
    docker run -i -t ubuntu:15.10 /bin/bash
        ```
    
## 容器使用
    
    ```
    # 直接输入 docker 命令来查看 Docker 客户端的所有命令选项
    docker
    
    # 通过命令 docker command --help 更深入的了解指定的 Docker 命令使用方法
    docker stats --help
     
    # 在docker容器中运行一个 Python Flask 应用来运行一个web应用
    docker run -d -P training/webapp python app.py
     
    # 使用 docker ps 来查看正在运行的容器
    docker ps
     
    # 查看WEB应用程序日志
    docker logs -f 7a38a1ad55c6
     
    # 查看WEB应用程序容器的进程
    docker top determined_swanson
    
    # 停止WEB应用容器
    docker stop determined_swanson   
     
    # 重启WEB应用容器
    docker start determined_swanson
    ```
    
