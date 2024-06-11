# Gin Demo

# Go 学习顺序

1. [gin-hello](./gin-hello): 最简单的 gin 项目
2. [gin-params](./gin-params): 处理请求参数
3. [gin-router](./gin-router): 路由配置
4. [gin-middleware](./gin-middleware): 中间件配置
5. [gin-template](./gin-template): 模板渲染
6. [gin-session](./gin-session): 会话管理
7. [gin-mysql](./gin-mysql): Mysql + ORM 集成
8. [gin-jwt](./gin-jwt): JWT 集成
9. [gin-swagger](./gin-swagger): Swagger 集成
10. [gin-logger](./gin-logger): 日志记录
11. [gin-rbac](./gin-rbac): 基于角色的访问控制
12. [gin-upload](./gin-upload): 文件上传
13. [gin-cache](./gin-cache): 缓存
14. [gin-websocket](./gin-websocket): WebSocket 集成
15. [gin-captcha](./gin-captcha): 验证码
16. [gin-cors](./gin-cors): 跨域请求
17. [gin-ratelimit](./gin-ratelimit): 限流
18. [gin-grpc](./gin-grpc): gRPC 集成
19. [gin-graphql](./gin-graphql): GraphQL 集成
20. [gin-oauth2](./gin-oauth2): OAuth2 集成
21. [gin-mq](./gin-mq): 消息队列集成
22. [gin-sse](./gin-sse): Server-Sent Events 集成
23. [gin-cron](./gin-cron): 定时任务集成

# 进阶

1. [gin-advanced-routing](./gin-advanced-routing): 高级路由配置
2. [gin-https](./gin-https): HTTPS 集成
3. [gin-multilingual](./gin-multilingual): 多语言支持
4. [gin-docker](./gin-docker): 使用 Docker 部署 Gin 应用
5. [gin-ci-cd](./gin-ci-cd): 使用 CI/CD 集成和部署 Gin 应用
6. [gin-config](./gin-config): 配置管理和加载
7. [gin-validation](./gin-validation): 数据验证
8. [gin-rate-limit-advanced](./gin-rate-limit-advanced): 高级限流策略
9. [gin-multi-db](./gin-multi-db): 多数据库支持
10. [gin-microservices](./gin-microservices): 构建微服务应用
11. [gin-metrics](./gin-metrics): 性能监控和指标收集

# 高级主题

24. **性能优化**

    - 学习如何分析和优化 Go 代码性能。
    - 使用 Go Profiler（pprof）进行性能分析。
    - 优化内存使用和垃圾回收。

25. **微服务架构**

    - 学习如何使用 Gin 和其他工具构建微服务架构。
    - 了解服务发现、负载均衡和服务间通信（如 gRPC 或 REST）。

26. **分布式系统**

    - 掌握分布式系统的基本原理，如一致性、可用性、分区容忍性（CAP 理论）。
    - 使用分布式追踪（如 Jaeger）监控系统。

27. **云原生开发**

    - 学习如何将 Gin 应用部署到云平台（如 AWS、GCP、Azure）。
    - 了解容器化（Docker）和编排工具（Kubernetes）。

28. **DevOps 和 CI/CD**
    - 学习如何使用 CI/CD 工具（如 Jenkins、GitHub Actions）进行自动化部署。
    - 了解基础设施即代码（Terraform、Ansible）。

# 相关技术

29. **Web 安全**

    - 学习常见的 Web 安全问题及其解决方案（如 SQL 注入、XSS、CSRF）。
    - 使用安全库和工具（如 go-secure、OWASP ZAP）进行安全检测。

30. **前端集成**

    - 学习如何将前端框架（如 React、Vue、Angular）与 Gin 后端集成。
    - 使用 Webpack 或 Vite 进行前后端代码的打包和构建。

31. **消息队列**

    - 学习如何使用其他消息队列系统（如 RabbitMQ、Apache Kafka）。
    - 了解消息队列的使用场景和最佳实践。

32. **数据存储**

    - 学习如何集成和使用其他数据库（如 PostgreSQL、MongoDB）。
    - 了解分布式数据库和缓存系统（如 Redis、Cassandra）。

33. **测试**
    - 学习如何编写单元测试、集成测试和端到端测试。
    - 使用 Go 的 testing 包和测试框架（如 Testify）。

# 错误解决

> go: module github.com/gin-gonic/gin: Get "https://proxy.golang.org/github.com/gin-gonic/gin/@v/list": dial tcp 142.251.42.241:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.

tips: 如果出现以上错误，可以配置代理

```shell
# 1、开启 go module
go env -w GO111MODULE=on

# 2、设置 GOPROXY
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/

# 如果是 macos 或 linux，可以直接设置环境变量：
# export GOPROXY=https://mirrors.aliyun.com/goproxy/

# 3、设置 GOSUMDB
go env -w GOSUMDB=off
#  export GOSUMDB=off // macOS 或 Linux
```

# 公认的工程目录（DDD 架构）

- cmd：存放项目的可执行文件,可能有多个 main 文件
- internal：存放项目的内部代码，不对外公开
- pkg：存放项目的公开代码，可以被其他项目依赖
- config/configs/etc：存放项目的配置文件
- scripts：存放项目的脚本文件
- docs：存放项目的文档
- third_party：存放项目的第三方依赖库
- bin: 编译的二进制文件
- build: 持续集成相关的配置文件
- deploy: 发布相关的配置文件
- test: 单元测试相关的配置文件
- api: 开发 API 接口
- init: 初始化项目的配置文件

# 创建项目

```shell
gin mod init [project name]
```

# 安装 Gin

```shell
go get -u github.com/gin-gonic/gin
```

# 运行项目

```shell
gin run main.go
```

# go.work

```shell
go 1.22.4

use (
    ./gin-hello
    ./gin-hello
    ./gin-middleware
    ./gin-params
    ./gin-template
    ./gin-session
    ./gin-orm
    ./gin-mysql
)

// 每次修改 go.mod 后，需要执行 go mod tidy 命令更新依赖
// go work sync
```
# Git 提交规范

-   fix: 修复 bug
-   feat: 新功能
-   docs: 文档
-   style: 格式
-   refactor: 重构
-   perf: 性能优化
-   test: 测试
-   build: 构建
-   ci: 持续集成
-   chore: 其他修改
