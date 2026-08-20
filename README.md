# 09 企业邮件投递与入站路由平台

纯Go实现的多租户事务邮件平台，提供REST、轻量gRPC探针和SMTP提交端口。默认使用内存存储与模拟SMTP传输，开发验证不需要外部邮件服务；PostgreSQL、对象存储、DNS、签名器、扫描器和真实SMTP均通过接口替换。

## 启动

要求Go 1.22+。配置默认读取`configs/config.yaml`，环境变量`MAIL_HTTP_ADDR`、`MAIL_SMTP_ADDR`、`MAIL_GRPC_ADDR`、`MAIL_MAX_BODY`和`MAIL_QUEUE_WORKERS`可覆盖配置。

```sh
go test ./...
go vet ./...
go run ./cmd/server
```

HTTP默认监听`:18089`，gRPC探针`:19089`，SMTP提交`:2525`。运行`./scripts/smoke.sh`可验证健康、就绪、指标和提交接口。

## REST主流程

```sh
curl http://localhost:18089/healthz
curl -X POST http://localhost:18089/v1/messages \
  -H 'Content-Type: application/json' \
  -d '{"tenant_id":"demo","from":"sender@example.test","to":"receiver@example.test","subject":"hello","text":"body","idempotency_key":"demo-1"}'
curl http://localhost:18089/v1/messages?tenant_id=demo
curl http://localhost:18089/v1/jobs
```

提交请求进入持久化抽象的队列，内存适配器由worker领取并通过模拟SMTP传输标记为`delivered`。相同租户和幂等键重复提交会返回同一消息。真实部署应接入PostgreSQL、可靠队列和SMTP适配器。

## SMTP入站

连接`:2525`后依次发送EHLO、MAIL FROM、RCPT TO、DATA和QUIT；示例见`examples/smtp-session.txt`。入站邮件以`tenant_id=smtp`写入同一消息服务，大小和头部注入受到限制。

## 架构

`internal`按tenant、domain、message、template、smtpout、smtpin、queue、bounce、policy、webhook、tracking拆分，每个领域含`domain`、`application`、`adapter`、`infrastructure`层。所有关键接口采用构造函数注入，跨模块调用传递`context.Context`，错误使用`fmt.Errorf`包装，日志使用`slog`。HTTP和SMTP具备请求限制、超时、恢复、请求ID、优雅停机、`/healthz`、`/readyz`和`/metrics`。

## 数据库迁移与源码规模

`migrations/0001_init.sql`和`0002_webhooks.sql`描述生产PostgreSQL表结构。非测试Go源码不少于2000行；计数排除测试文件、生成代码、SQL迁移、配置样例、依赖和构建产物。文件按领域拆分，禁止用重复代码凑行数。
