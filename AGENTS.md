# AGENTS.md

本目录为**模拟报文平台**（third-mock-platform）**仓库根目录**（单体 git 仓库，远端 https://github.com/lincy-NY/third-mock-platform.git），内含服务端与前端两个模块目录。解决开发过程中**对方系统未开发好、只提供报文样例或文档**时，己方系统无法联调的问题：由本平台模拟对方接口，让己方系统正常发起接口请求并获得模拟响应。

## 模块划分

| 模块目录 | 技术栈 | 职责 |
|---|---|---|
| `third-mock-server/` | Go ≥ 1.22，零第三方依赖 | mock 服务 + 管理 API + 内嵌前端产物分发；含 `cmd/mockctl` CLI 管理客户端 |
| `third-mock-web/` | Vue3 + Vite5 + Element Plus + Pinia + Vue Router | 浏览器访问的 mock 规则管理界面 |

- **服务端**部署在服务器上**手动启动**；前端构建产物可嵌入服务端二进制（go:embed），单文件部署
- **CLI（mockctl）**：在**不方便通过浏览器访问的环境**（如仅 SSH 终端）对远端 mock 服务进行管理维护，留在服务端仓库以共享 model
- **前端**开发时经 Vite proxy 访问本地服务端管理 API；生产环境由服务端同源分发
- **目标平台**：Windows / macOS / Linux（amd64 + arm64），统一走工作区根 `make release` 或服务端 `make build-all` 交叉编译

## 目录结构

```
third-mock-platform/          # 仓库根目录（单体 git 仓库）
├── .agents/                  # AI相关数据
│   ├── skills/               # 项目特定skill
│   └── work/                 # 工作任务目录
│       └── {branch-name}/    # 按分支名分目录
│           └── work.md       # 本次任务工作记录
├── .docs/                    # 文档目录
│   ├── env.md                # 环境配置（服务器地址、端口等）
│   └── {文档目录名}/         # 按任务分目录
├── Makefile                  # 聚合构建入口（release / build-web / embed-web / test）
├── AGENTS.md                 # 本文件
├── third-mock-server/        # 模块1：Go 服务端（含 CLI）
│   ├── cmd/server/           # 服务端入口
│   ├── cmd/mockctl/          # CLI 管理客户端入口
│   ├── internal/             # 服务端内部实现（不对外暴露）
│   │   ├── model/            # 核心数据模型（规则/场景/请求记录）
│   │   ├── store/            # JSON 文件持久化
│   │   ├── engine/           # mock 核心引擎（匹配/场景/模板渲染）
│   │   ├── server/           # HTTP 层（mock 路由 + 管理 API + 静态资源解析）
│   │   └── static/           # go:embed 目录（dist/ 为前端构建产物，含占位文件）
│   ├── configs/              # 配置文件样例
│   ├── examples/             # 对方提供的报文样例（素材，只读参考）
│   ├── docs/                 # 服务端设计文档
│   ├── go.mod                # module third-mock-server，零第三方依赖
│   └── Makefile              # build / build-all / embed-web / test / fmt
└── third-mock-web/           # 模块2：Vue3 前端
    ├── index.html
    ├── package.json          # vue3 / vite5 / element-plus / pinia / vue-router
    ├── vite.config.js        # dev 代理 /api → 本地 mock 服务
    └── src/
        ├── api/              # 管理 API 封装（契约随服务端实现同步调整）
        ├── router/           # 路由（规则管理/请求记录/设置）
        ├── stores/           # Pinia（服务端连接信息等）
        └── views/            # 页面
```

## 技术选型约定

| 项 | 约定 | 说明 |
|---|---|---|
| 服务端语言 | Go ≥ 1.22，单 module | `cmd/` 多入口（server + mockctl） |
| 服务端依赖 | **默认零依赖**，纯标准库 | 新增依赖必须先征得用户同意并说明理由 |
| 存储 | JSON 文件（`internal/store`） | 无数据库；原子写入（临时文件+rename） |
| 前端 | Vue3 + Vite5 + Element Plus + Pinia + Vue Router | 需 node ≥ 18（Vite 锁 v5）；产物经 go:embed 嵌入或外部目录加载 |
| 静态资源 | 双模式（`internal/server.ResolveStatic`） | 配置 `web_dist` 外部目录优先（开发调试），否则回退 embed 产物（单二进制部署） |
| 跨平台 | `CGO_ENABLED=0` 交叉编译 | 禁止引入 cgo；路径用 `filepath`，禁止硬编码 `/` 或 `\` |

## 开发流程

进行任何代码修改前，必须先加载 `project-dev-flow` skill 并按流程执行：

1. **检查分支状态**：确认当前分支是否为任务分支
2. **分支不正确时**：先还原修改，切换到正确分支后再重新应用
3. **更新 work.md**：每次修改后在 `.agents/work/{branch-name}/work.md` 记录变更
4. **改完必验证**：服务端 `make fmt && make vet && make build`（涉及测试跑 `make test`）；前端 `npm run build`；涉及嵌入产物时走根目录 `make release` 全链路验证

## 红线规则（必须遵守）

> **以下两条为最高优先级规则，任何情况下都必须遵守，违反将导致严重后果。**

### 1. 禁止直接提交代码
- **未经用户明确同意，禁止执行 `git add`、`git commit`、`git push` 等任何代码提交操作**
- 修改完成后，只能生成提交建议供用户参考
- 用户确认后才可执行提交操作

### 2. 禁止破坏运行数据
- **未经用户明确同意，禁止删除/覆盖 `data/` 下的规则与请求记录文件**
- 涉及数据结构变更（rules.json 格式调整）时，必须提供兼容旧数据的加载逻辑或迁移说明
- 用户确认后才可执行数据操作

---

## 重要规则

### 0. 修改代码前必须先确认分支
- **每次修改代码前，必须先检查当前分支是否正确**
- 确认命令：`git status` + `git branch`
- 多模块任务涉及不同目录时逐一确认目录归属，分支统一在本仓库管理

### 1. 触发条件
只有当用户下达明确的开发任务指令时，才遵循完整的开发流程 skill。日常的简单修改、问答等不需要触发完整流程。

### 2. 提交代码
- **永远不要主动提交代码**
- 修改完成后，生成提交建议供用户参考
- 用户确认后才执行提交操作

### 3. 处理冲突
- **永远不要主动解决冲突**
- 发现冲突时，告知用户并列出冲突文件（注明所属仓库）
- 等待用户手动解决冲突后再继续

### 4. 代码规范（Go）
- 格式以 `gofmt` 为准，禁止手工对齐导致的混合缩进
- 注释以 `//` 标准 godoc 风格编写：**包注释、导出标识符必须有英文+中文双语注释**（参考 `internal/model/doc.go` 现有风格）
- 错误必须显式处理，禁止 `_` 吞错；错误信息用 `%w` 包装保留链路
- 并发访问的共享状态必须加锁；文件写入必须原子（临时文件+rename）

### 5. 跨平台规范（红线级约束）
- 禁止 cgo、禁止平台专属 syscall（如需区分平台，放 `*_windows.go` / `*_unix.go`）
- 禁止在代码中硬编码可执行文件后缀（Windows 下为 `.exe`），构建后缀由 Makefile 处理
- 每次服务端修改后必须确认 `make build-all` 可通过（交叉编译不带 cgo 才有意义）

### 6. 前端规范（third-mock-web）
- Vue3 + Element Plus 技术栈；新增前端依赖必须先征得用户同意
- 前端通过管理 API（统一前缀 `/api/`）与服务端交互；API 变更需同步更新 `src/api/`
- 契约未冻结前，`src/api/` 中的接口函数需注明"待定"状态

### 7. CLI 规范（mockctl）
- 子命令风格：`mockctl <command> [flags]`，帮助信息完整
- CLI 通过管理 API 操作**远端** mock 服务，不直接读写服务端数据文件
- 连接目标通过 `--server` 参数或环境变量指定

### 8. 版本与构建
- 版本号在 `third-mock-server/cmd/server/main.go` 的 `version` 变量统一维护，构建时 `-ldflags -X` 注入
- 各模块构建用各自目录的 Makefile；**发布统一走仓库根 `make release`**（web 构建 → embed → 全平台编译）
- 服务端与 CLI 版本需保持同步演进，接口不兼容时必须升版本
- 前端版本随 package.json 演进

### 9. 待定项（需求未冻结，实现前必须与用户确认）
- **协议范围**：当前仅明确 HTTP/HTTPS；是否扩展 TCP 报文（8583/定长）、WebService/SOAP 待定
- mock 报文模板语法细节、鉴权方式、规则导入格式待定
- 管理 API 契约（路径/字段）未冻结，`src/api/` 与 CLI 请求结构以服务端实现为准
- 涉及待定项的开发，先提问确认再动手，禁止自行拍板

### 10. 文档生成
- PPT 生成使用 `edraw-ai-mcp-server_generate_ppt`
- 思维导图使用 `edraw-ai-mcp-server_generate_mindmap`
- 流程图使用 `edraw-ai-mcp-server_generate_flowchart`

### 11. 环境配置
- 服务器地址、端口等环境信息统一维护在 `.docs/env.md`
- 需要查看环境信息时，直接读取 `.docs/env.md`
