# Bug

分页构建重复追加记录，并返回与存储共享的切片，页面本地修改会污染服务状态。

## 触发方式

```bash
go test ./internal/platform/store -run '^TestListPageReturnsIndependentRecords$' -count=1
```

## 错误信息

```text
--- FAIL: list page independence check
    pagination.go:20: got 8 records
```
