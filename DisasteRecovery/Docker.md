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
    ```
