# Bug

SMTP 入站解析和扫描器把调用方 context 替换为 Background，取消后仍继续读取慢速流。

## 触发方式

```bash
go test ./internal/smtpin/adapter -run '^TestParseContextStopsAfterCancellation$' -count=1
```

## 错误信息

```text
--- FAIL: TestParseContextStopsAfterCancellation
    parser_test.go:31: parser continued reading after cancellation
```
