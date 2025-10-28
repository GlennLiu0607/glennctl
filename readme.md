# glennctl

简体中文 | English

glennctl 是基于 go-zero 增强的代码生成与脚手架工具，统一以 `glennctl` 命令使用（不再支持 `goctl`）。

## 安装

- 安装最新版本：

```
go install github.com/GlennLiu0607/glennctl@latest
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

### DM8 模型生成（新增）

- 现已支持达梦 DM8 数据源模型生成，命令结构与 MySQL/PostgreSQL 一致：

```
glennctl model dm datasource \
  --url "dm://USER:PASSWORD@127.0.0.1:5236?autoCommit=true" \
  --schema SYSDBA \
  --table "*" \
  --dir ./model \
  --style go_zero
```

- 关键参数说明：
  - `--url` DM8 连接串，示例：`dm://SYSDBA:SYSDBA@127.0.0.1:5236?autoCommit=true`
  - `--schema` 指定所有者/模式（默认 `SYSDBA`），用于过滤 `ALL_*` 视图
  - `--table` 支持通配（如：`user*`, `ORDER_*`），多个用逗号分隔
  - 其余选项与 MySQL/PostgreSQL 一致：`--dir` 输出目录、`--style` 命名风格、`--cache` 是否生成缓存、`--strict` 严格模式
  - 可通过命令组持久化参数：
    - `--ignore-columns` 默认忽略时间字段：`create_at, created_at, create_time, update_at, updated_at, update_time`
    - `--prefix` 缓存前缀，默认 `cache`

- 注意事项：
  - DM8 连接示例亦可使用构造函数形式（等价）：`dm.BuildDsn(host, port, user, password, params)`，例如生成 `dm://user:password@127.0.0.1:5236?autoCommit=true`
  - `--schema` 与连接串无强耦合，主要用于元数据检索（`ALL_TABLES/ALL_TAB_COLUMNS/ALL_INDEXES` 等），请与数据库账号具备的可见对象一致
  - 表筛选在工具侧完成，连接只需指向可访问的实例地址

## 常用选项说明

- `--home` 指定模板根目录
- `--remote` 指定远程模板仓库
- `--branch` 指定模板分支
- `--style` 文件命名风格（默认 `go_zero`）
- `-v/--verbose` 输出更多日志

## 发布说明

1. 移除 `go.mod` 中的 `replace` 指令（发布模块不允许包含会改变解析的 replace）。
2. 模块路径必须与仓库地址一致：
   - 若发布到 `github.com/GlennLiu0607/glennctl`，请将 `go.mod` 第一行修改为 `module github.com/GlennLiu0607/glennctl`，并统一更新项目内的 import 前缀为该模块路径；
   - 若保持 `github.com/GlennLiu0607/glennctl`，则仓库地址需与之匹配，安装时使用该路径。
3. 推送代码到远端仓库并打版本标签：

```
git tag v1.0.0
git push origin v1.0.0
```

4. 用户安装方式：

```
go install github.com/GlennLiu0607/glennctl@v1.0.0
```

5. 验证：

```
glennctl version
glennctl -h
```

6. 注意：
    每次发布前必须修改 `internal/version/version.go` 中的 `const BuildVersion = "1.0.3"` 为当前要发布的tag版本，不然 glennctl -v/--version 显示有问题
    更新本地版本直接 `go install github.com/GlennLiu0607/glennctl@v1.0.4` 或 `go install github.com/GlennLiu0607/glennctl@latest` 即可
