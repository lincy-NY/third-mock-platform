import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// 开发模式下将 /api 代理到本地 mock 服务管理端口，生产环境同源直接访问。
// In dev mode /api is proxied to the local mock server admin port; in production
// the frontend is served by the server itself (same origin).
const SERVER_TARGET = process.env.MOCK_SERVER || 'http://127.0.0.1:8080'

export default defineConfig({
  plugins: [vue()],
  server: {
    port: 5173,
    proxy: {
      '/api': {
        target: SERVER_TARGET,
        changeOrigin: true,
      },
    },
  },
  build: {
    outDir: 'dist',
    emptyOutDir: true,
  },
})
