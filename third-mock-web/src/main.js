import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'
import router from './router'

// 应用入口：挂载 Element Plus（全量引入，保持零额外构建插件）、Pinia 与路由。
// App entry: mounts Element Plus (full import, no extra build plugins), Pinia and the router.
createApp(App).use(createPinia()).use(router).use(ElementPlus).mount('#app')
