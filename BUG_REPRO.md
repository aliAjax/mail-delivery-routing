# Bug

并发预留配额时读写租户 map 和 Used 不是原子操作，可能多放行并触发数据竞争。

## 触发方式

```bash
go test -race ./internal/tenant/application -run '^TestReserveQuotaConcurrent$' -count=1
```

## 错误信息

```text
WARNING: DATA RACE
--- FAIL: TestReserveQuotaConcurrent
    service_test.go:32: successes=33, want 32
```
