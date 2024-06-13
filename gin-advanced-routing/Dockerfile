# 使用官方的golang镜像作为构建阶段的基础镜像
FROM golang:1.22.4-alpine AS build-stage

# 设置环境变量
ENV GO111MODULE=on \
CGO_ENABLED=0 \
GOOS=linux \
GOARCH=amd64

# 为了确保可重复的构建，创建一个工作目录
WORKDIR /app

# 将go.mod和go.sum文件复制到工作目录
COPY go.mod go.sum ./

# 下载所有依赖，缓存这些依赖，如果依赖没有变化，这一步将被缓存
RUN go mod download

# 将项目的源文件复制到工作目录
COPY . .

# 构建Go应用程序
RUN go build -o main .

# 使用一个更小的基础镜像作为最终运行阶段的基础镜像
FROM alpine:latest

# 设置时区信息（可选）
RUN apk --no-cache add ca-certificates tzdata

# 创建工作目录
WORKDIR /root/

# 从构建阶段复制可执行文件到最终镜像
COPY --from=build-stage /app/main .

# 暴露应用程序运行的端口（根据你的应用需要修改）
EXPOSE 8080

# 运行可执行文件
CMD ["./main"]
