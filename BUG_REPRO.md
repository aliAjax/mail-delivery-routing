# Bug

模板文件读取路径未关闭 reader，渲染失败时还丢失了底层错误链。

## 触发方式

```bash
go test ./internal/template/infrastructure -run '^TestLoadTemplateClosesReader$' -count=1
```

## 错误信息

```text
--- FAIL: TestLoadTemplateClosesReader
    filesystem_test.go:24: reader was not closed
```
