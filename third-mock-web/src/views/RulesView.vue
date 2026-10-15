<script setup>
// 规则管理页骨架：规则 CRUD 界面随管理 API 实现后充实。
// Rules page skeleton: the CRUD UI will be fleshed out once the admin API lands.
</script>

<template>
  <el-card>
    <template #header>
      <div class="card-header">
        <span>Mock 规则</span>
        <el-button type="primary" disabled>新建规则</el-button>
      </div>
    </template>
    <el-empty description="管理 API 待实现后开放规则管理" />
  </el-card>
</template>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
