// mock-server 模拟报文平台服务端入口。
// mock-server is the entry point of the mock platform server.
//
// 职责（待实现）：启动 mock 服务（HTTP）+ 管理 API + 内嵌 H5 管理页面。
// Responsibilities (TBD): start the mock service, admin API and embedded H5 console.
//
// 部署方式：上传至服务器后手动启动进程。
// Deployment: manually started on the target server.
package main

import "fmt"

// 版本号统一维护，构建时可通过 -ldflags 注入。
// Version is maintained here and can be injected via -ldflags at build time.
var version = "dev"

func main() {
	fmt.Printf("mock-server %s (not implemented yet, see AGENTS.md)\n", version)
}
