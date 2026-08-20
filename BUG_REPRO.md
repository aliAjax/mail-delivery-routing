# Bug

模板缺少变量时默认值路径向 nil map 写入，预览请求触发 panic。

## 触发方式

```bash
go test ./internal/template/application -run '^TestRenderWithFallbackAcceptsNilVariables$' -count=1
```

## 错误信息

```text
panic: assignment to entry in nil map
```
