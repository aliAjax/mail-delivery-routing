# Bug

Webhook 状态错误被格式化成普通字符串，重试器无法通过 errors.As 识别 HTTPStatusError。

## 触发方式

```bash
go test ./internal/webhook/application -run '^TestDeliverCheckedPreservesHTTPStatusError$' -count=1
```

## 错误信息

```text
--- FAIL: TestDeliverCheckedPreservesHTTPStatusError
    service_test.go:20: err=webhook rejected (server): webhook returned status 503 is not HTTPStatusError
```
