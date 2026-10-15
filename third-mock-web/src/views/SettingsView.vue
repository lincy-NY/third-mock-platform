<script setup>
import { onMounted, ref } from 'vue'
import { getSystemInfo } from '../api/system'

// 设置页骨架：展示服务端连通性与版本信息。
// Settings page skeleton: shows server connectivity and version info.
const info = ref(null)
const error = ref('')

onMounted(async () => {
  try {
    info.value = await getSystemInfo()
  } catch (e) {
    error.value = e.message
  }
})
</script>

<template>
  <el-card>
    <template #header><span>设置</span></template>
    <el-alert v-if="error" type="warning" :title="`服务端未连接：${error}`" :closable="false" />
    <el-descriptions v-else-if="info" title="服务端信息" :column="1" border>
      <el-descriptions-item label="版本">{{ info.version }}</el-descriptions-item>
    </el-descriptions>
  </el-card>
</template>
