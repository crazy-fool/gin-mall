# 构建golang运行环境 使用别名：builder
FROM golang:1.18-alpine as builder

# 设置环境变量
ENV HOME /app
ENV CGO_ENABLED 0
ENV GOOS linux

# 设置工作目录 - 我们所有的文件都存放在工作目录中
WORKDIR /app
COPY . .
# 下载依赖
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download

# 编译app
RUN go build -v -a -installsuffix cgo -o server cmd/server/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /bin/

# 将上个容器编译的二进制文件复制到 工作目录
# 也就是：copy golang环境/工作目录/demo alpine环境/工作目录
COPY --from=builder /app/server .
COPY --from=builder /app/config/* ./config/

# 所以这里执行的命令是：/bin/server
ENTRYPOINT ["/bin/server"]