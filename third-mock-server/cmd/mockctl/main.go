// mockctl 模拟报文平台 CLI 管理客户端入口。
// mockctl is the CLI admin client entry of the mock platform.
//
// 定位：服务端部署在服务器上手动启动；CLI 用于不方便通过浏览器访问的环境
// （如仅 SSH 终端）对远端 mock 服务进行管理维护。
// The server is started manually on the host; mockctl administers it from
// environments where a browser is not available (e.g. SSH-only terminals).
//
// 状态：骨架占位，待实现（子命令规划见 AGENTS.md）。
// Status: skeleton placeholder; planned subcommands are listed in AGENTS.md.
package main

import "fmt"

func main() {
	fmt.Println("mockctl (not implemented yet, see AGENTS.md)")
}
