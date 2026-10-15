# init-前后端分离拆分 工作记录

- **创建时间**：2026-10-15 09:30
- **任务描述**：项目前后端分离拆分（web 独立 Vue 仓库 + 服务端仓库含 CLI）
- **文档目录**：.docs/20261015项目拆分/
- **需求描述**：将原单仓库骨架拆分为独立 git 仓库：third-mock-server（Go 服务端 + mockctl CLI）、third-mock-web（Vue3 前端）；服务端支持 embed 与外部目录两种静态资源加载方式。**（后续调整：经用户确认改为单体 git 仓库，远端 github.com/lincy-NY/third-mock-platform）**
- **状态**：已完成

---

## 用户已确认决策

| 决策项 | 结论 |
|--------|------|
| 目录布局 | ~~两独立子仓库~~ → **调整：单体 git 仓库**（远端 https://github.com/lincy-NY/third-mock-platform.git），server/web 为仓库内两个模块目录 |
| 仓库划分 | third-mock-server（含 cmd/mockctl）、third-mock-web |
| 前端集成 | 两者兼容：优先 go:embed 产物，支持配置外部目录覆盖 |
| CLI 去留 | 留在服务端仓库，共享 model |
| Vue 技术栈 | Vue3 + Vite5 + Element Plus + Pinia + Vue Router（JS，非 TS） |

## 工作计划

- [x] 阶段1：需求分析（用户问答确认）
- [x] 阶段2：方案设计（见 .docs/20261015项目拆分/20261015_前后端分离_设计.md）
- [x] 阶段3：编码实现
- [x] 阶段4：自测验证
- [x] 阶段5：文档更新
- [x] 阶段6：交付提交（生成提交建议，不自动提交）

## 工作记录

### 2026-10-15 09:30

**操作内容**：
- 需求分析与方案设计：确认两仓库布局、静态资源双模式加载、CLI 留在服务端
- 环境确认：Node v18.20.8 / npm 10.8.2，Vite 锁定 v5

### 2026-10-15 10:00

**操作内容**：
- 编码实现：两仓库结构落地 + git init（main 分支，未提交）
- 占位方案调整：`internal/static/dist` 用 `.keep` 占位（go:embed 要求目录非空），发布后 git status 干净
- 新增 `ResolveStatic` 静态资源双模式解析（外部目录优先，回退 embed）+ 3 个单测

**涉及文件**：
| 仓库 | 文件 | 操作 |
|------|------|------|
| third-mock-server | `go.mod`、`cmd/{server,mockctl}/main.go`、`internal/**` | 迁移自原单仓库骨架，module 更名 third-mock-server |
| third-mock-server | `internal/static/static.go`、`internal/static/dist/.keep` | 新建：go:embed 封装与占位 |
| third-mock-server | `internal/server/static.go`、`static_test.go`、`doc.go` | 新建/修改：双模式静态资源解析 |
| third-mock-server | `Makefile`、`.gitignore` | 修改/新建：embed-web 目标、dist 排除规则 |
| third-mock-web | `package.json`、`vite.config.js`、`index.html`、`.gitignore` | 新建：Vite5 + Vue3 脚手架 |
| third-mock-web | `src/main.js`、`App.vue`、`router/`、`stores/`、`api/`、`views/` | 新建：布局骨架 + 管理 API 封装（契约待定） |
| 工作区 | `Makefile`、`.gitignore`、`AGENTS.md` | 重写：聚合构建入口与新结构规范 |

**自测结果**：
| 测试项 | 结果 |
|--------|------|
| server `make fmt && make vet && make test` | ✓（3 个静态解析用例通过）|
| server `make build` | ✓ |
| web `npm install && npm run build` | ✓ |
| 根目录 `make release`（web 构建 → embed → 12 平台产物） | ✓ |
| 发布后两仓库 git status | ✓（无入库噪音，.keep 保留）|

**遗留事项**：
- ~~两仓库均为全新 git init，首次 commit 由用户确认后执行~~ → 已按用户指令改为单体仓库，见 10:20 记录
- 管理服务端启动配置加载 `web_dist` 并接入 ResolveStatic 属后续功能开发任务

---

## 提交建议（用户确认后执行）

```
main（首次提交）
    1、单体仓库初始化：third-mock-server（Go 服务端 + mockctl CLI）与 third-mock-web（Vue3 前端）模块
    2、静态资源双模式解析 ResolveStatic 与 go:embed 封装及单测
    3、聚合构建链路（根 Makefile release）与 AGENTS.md 仓库规范
```

确认后执行：`git add . && git commit`，随后可 `git push -u origin main`
