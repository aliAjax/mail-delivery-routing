# Bug

SMTP 外发重试没有保留调用方 context，且包装错误时截断了可判定的底层错误。

## 触发方式

```bash
go test ./internal/smtpout/application -run '^TestSendWithRetryHonorsContext$' -count=1
```

## 错误信息

```text
--- FAIL: TestSendWithRetryHonorsContext
    service_test.go:28: retry ignored context cancellation
```
