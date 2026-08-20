# Bug

消息状态迁移未拒绝终态回滚，已完成或已退信消息仍可被改回处理中。

## 触发方式

```bash
go test ./internal/message/domain -run '^TestMessageTransitionRejectsTerminalRollback$' -count=1
```

## 错误信息

```text
--- FAIL: TestMessageTransitionRejectsTerminalRollback
    lifecycle_test.go:22: terminal transition accepted
```
