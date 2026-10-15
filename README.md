# third-mock-platform 模拟报文平台

解决开发联调时**对方系统未就绪、只提供报文样例或文档**的问题：本平台模拟对方接口，让己方系统正常发起请求并获得模拟响应。

## 组成

| 组件 | 位置 | 说明 |
|---|---|---|
| 服务端 | `cmd/server` | mock 服务 + 管理 API + 内嵌 H5 管理页，部署在服务器手动启动 |
| H5 前端 | `web/` | 浏览器访问的规则管理界面，go:embed 嵌入服务端二进制 |
| CLI 客户端 | `cmd/mockctl` | 无浏览器环境（仅 SSH）下管理维护远端 mock 服务 |

## 快速开始（骨架阶段）

```bash
make build        # 构建当前平台
make build-all    # 交叉编译 windows/linux/macos (amd64+arm64)
```

> 当前为项目骨架阶段，功能开发计划与协作规范见 [AGENTS.md](AGENTS.md)。

## 目录说明

见 [AGENTS.md](AGENTS.md) 的「目录结构」章节。
