FROM golang:alpine AS gtag


ENV GOPROXY https://mirrors.aliyun.com/goproxy,https://goproxy.io,direct
# 标签修改程序版本
ENV TAG_VERSION 1.4.0

# 安装标签处理程序
RUN go install github.com/favadi/protoc-go-inject-tag@v${TAG_VERSION}





FROM golang:alpine AS protolint


ENV GOPROXY https://mirrors.aliyun.com/goproxy,https://goproxy.io,direct
# 静态检查版本
ENV LINT_VERSION 0.41.0

# 安装静态检查检查程序
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update
RUN apk add git
RUN go install github.com/yoheimuta/protolint/cmd/protolint@v${LINT_VERSION}





# 打包真正的镜像
FROM rvolosatovs/protoc:3.3.0


LABEL author="storezhang<华寅>" \
    email="storezhang@gmail.com" \
    qq="160290688" \
    wechat="storezhang" \
    description="Drone持续集成Protobuf插件，集成所有常见的Protobuf语言工具以及常用的插件"


# 复制文件
COPY --from=gtag /go/bin/protoc-go-inject-tag /usr/bin/gtag
COPY --from=protolint /go/bin/protolint /usr/bin/protolint
COPY protobuf /bin


RUN set -ex \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/protobuf \
    \
    \
    \
    && rm -rf /var/cache/apk/*


ENTRYPOINT /bin/protobuf
