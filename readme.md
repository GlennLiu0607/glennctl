# glennctl

简体中文 | English

glennctl 是基于 go-zero 增强的代码生成与脚手架工具，统一以 `glennctl` 命令使用（不再支持 `goctl`）。

## 安装

- 安装最新版本：

```
go install github.com/glenn/glennctl@latest
```

- 验证安装：

```
glennctl -h
glenn -h
```

## 模板管理

- 初始化所有内置模板：

```
glennctl template init
```

- 指定模板目录初始化：

```
glennctl template init --home ./_templates
```

- 更新某一分类模板（示例：TDengine）：

```
glennctl template update -c tdengine --home ./_templates
```

- 回滚某一分类的指定模板文件：

```
glennctl template revert -c tdengine -n model.tpl --home ./_templates
```

## RPC 代码生成

- 创建新的 zRPC 服务骨架：

```
glennctl rpc new user --module github.com/your/repo --style go_zero
```

- 从 .proto 生成 zRPC 目录结构（单服务）：

```
glennctl rpc protoc ./idl/user.proto --zrpc_out ./rpc/user --style go_zero
```

- 多服务模式：

```
glennctl rpc protoc ./idl/service.proto --zrpc_out ./rpc/service -m
```

## API 代码生成

- 新建 API 项目：

```
glennctl api new user-api --module github.com/your/repo --style go_zero
```

- 从 .api 文件生成 Go 服务骨架：

```
glennctl api go --api ./api/user.api --dir ./api/user --style go_zero
```

- 生成文档与校验：

```
glennctl api doc --dir ./api/user -o ./api/user/openapi.json
glennctl api validate --api ./api/user.api
```

## 模型生成（SQL / Mongo / TDengine）

- 对应模型生成命令依照 go-zero 保持一致；
- 使用模板管理命令可对 `sql`、`mongo`、`tdengine` 分类模板进行初始化、更新与回滚：

```
glennctl template update -c sql
glennctl template update -c mongo
glennctl template update -c tdengine
```

## 常用选项说明

- `--home` 指定模板根目录
- `--remote` 指定远程模板仓库
- `--branch` 指定模板分支
- `--style` 文件命名风格（默认 `go_zero`）
- `-v/--verbose` 输出更多日志

## 发布说明

1. 确保 `go.mod` 模块路径为 `github.com/glenn/glennctl`；
2. 推送代码到远端仓库并打版本标签：

```
git tag v1.0.0
git push origin v1.0.0
```

3. 用户安装方式：

```
go install github.com/glenn/glennctl@v1.0.0
```

4. 验证：

```
glennctl version
glennctl -h
```

如需自动发布二进制，可添加 GitHub Actions（例如使用 `goreleaser`），但通过 `go install` 即可满足大多数使用场景。
