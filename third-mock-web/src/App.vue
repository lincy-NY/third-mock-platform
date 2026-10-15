<script setup>
import { useRoute } from 'vue-router'

const route = useRoute()

// 顶部菜单与路由的映射，新增页面时同步补充。
// Mapping between the top menu and routes; extend when adding new pages.
const menus = [
  { path: '/rules', title: '规则管理' },
  { path: '/records', title: '请求记录' },
  { path: '/settings', title: '设置' },
]
</script>

<template>
  <el-container class="layout">
    <el-header class="header">
      <span class="brand">模拟报文平台</span>
      <el-menu mode="horizontal" :default-active="route.path" router class="menu">
        <el-menu-item v-for="m in menus" :key="m.path" :index="m.path">
          {{ m.title }}
        </el-menu-item>
      </el-menu>
    </el-header>
    <el-main class="main">
      <router-view />
    </el-main>
  </el-container>
</template>

<style>
body {
  margin: 0;
  background: #f5f7fa;
}
.layout {
  min-height: 100vh;
}
.header {
  display: flex;
  align-items: center;
  gap: 24px;
  background: #fff;
  border-bottom: 1px solid #e4e7ed;
}
.brand {
  font-size: 18px;
  font-weight: 600;
}
.menu {
  flex: 1;
  border-bottom: none;
}
.main {
  max-width: 1200px;
  width: 100%;
  margin: 0 auto;
}
</style>
