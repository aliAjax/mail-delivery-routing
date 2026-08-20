# Bug

配置热加载包装错误时丢失原始错误类型，调用方无法识别无效配置。

## 触发方式

```bash
go test ./internal/platform/config -run '^TestReloadPreservesInvalidConfigError$' -count=1
```

## 错误信息

```text
--- FAIL: TestReloadPreservesInvalidConfigError
    runtime_test.go:16: error chain lost InvalidConfigError
```
