# 使用官方 golang 镜像作为基础镜像
FROM golang:1.17 AS builder

# 设置工作目录
WORKDIR /app

# 将 go.mod 和 go.sum 文件复制到容器中
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 将源代码复制到容器中
COPY . .

# 编译项目
RUN CGO_ENABLED=0 GOOS=linux go build -o VecTextSearch ./cmd/main.go

# 使用官方 alpine 镜像作为基础镜像以减小镜像大小
FROM alpine:latest

# 添加 ca 证书，以支持 HTTPS 请求
RUN apk --no-cache add ca-certificates

# 从 builder 阶段复制编译好的二进制文件
COPY --from=builder /app/VecTextSearch /app/VecTextSearch

# 设置工作目录
WORKDIR /app

# 设置环境变量
ENV VECTEXTSEARCH_OPENAI_KEY=your_openai_api_key_here
ENV VECTEXTSEARCH_API_PORT=8000
ENV VECTEXTSEARCH_WEAVIATE_URL=localhost:8888

# 暴露端口
EXPOSE 8000

# 运行程序
CMD ["/app/VecTextSearch"]
