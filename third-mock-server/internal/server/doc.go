// Package server 实现服务端 HTTP 层：mock 转发路由、管理端 API、静态资源服务。
// Package server implements the HTTP layer: mock routes, admin API and static asset serving.
//
// 静态资源双模式：优先使用配置的外部前端产物目录（web_dist，开发/热更新场景），
// 否则回退 go:embed 嵌入产物（单二进制部署场景），见 ResolveStatic。
// Static assets load in dual mode: the configured external dist directory
// (web_dist, for development) takes precedence, otherwise fall back to the
// go:embed assets for single-binary deployment. See ResolveStatic.
//
// 状态：骨架，路由与 API 待实现。
// Status: skeleton; routes and admin API are to be implemented.
package server
