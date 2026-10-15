import { defineStore } from 'pinia'

// 服务端连接信息：开发模式走 Vite 代理（baseUrl 为空即同源），
// 预留baseUrl 便于直连远端服务端调试。
// Server connection info: dev mode goes through the Vite proxy (empty baseUrl
// means same-origin); baseUrl is reserved for debugging against a remote server.
export const useServerStore = defineStore('server', {
  state: () => ({
    baseUrl: '',
    version: '',
  }),
})
