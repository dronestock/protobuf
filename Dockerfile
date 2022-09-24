FROM golang:alpine AS gtag


ENV GOPROXY https://goproxy.cn,https://mirrors.aliyun.com/goproxy,https://goproxy.io,direct
# 标签修改程序版本
ENV TAG_VERSION 1.4.0

# 安装标签处理程序
RUN go install github.com/favadi/protoc-go-inject-tag@v${TAG_VERSION}





# 打包真正的镜像
FROM ccr.ccs.tencentyun.com/storezhang/protobuf:0.0.1


LABEL author="storezhang<华寅>" \
    email="storezhang@gmail.com" \
    qq="160290688" \
    wechat="storezhang" \
    architecture="AMD64/x86_64" version="latest" build="2022-01-20" \
    description="Drone持续集成Protobuf插件，集成所有常见的Protobuf语言工具以及常用的插件"


# 复制文件
COPY --from=gtag /go/bin/protoc-go-inject-tag /usr/bin/gtag
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
