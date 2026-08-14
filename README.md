# 多层配置合并

文件默认值 ← 环境变量覆盖；CLI 打印最终配置。

## 测试

```bash
set GOTOOLCHAIN=local
go test ./... -count=1
```
