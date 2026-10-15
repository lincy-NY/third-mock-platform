import { createRouter, createWebHistory } from 'vue-router'

// 路由表：页面骨架已就位，对应视图随功能实现逐步充实。
// Route table: view skeletons are in place and will be fleshed out feature by feature.
const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', redirect: '/rules' },
    { path: '/rules', component: () => import('../views/RulesView.vue') },
    { path: '/records', component: () => import('../views/RecordsView.vue') },
    { path: '/settings', component: () => import('../views/SettingsView.vue') },
  ],
})

export default router
