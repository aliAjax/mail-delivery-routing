# Bug

批量领取更新 job 状态后没有写回队列 map，两个 worker 可以领取同一个 job。

## 触发方式

```bash
go test -race ./internal/queue/application -run '^TestClaimBatchDoesNotDuplicateJobs$' -count=1
```

## 错误信息

```text
--- FAIL: TestClaimBatchDoesNotDuplicateJobs
    service_test.go:35: job j1 claimed twice
```
