import { request } from './client'

/**
 * 系统信息 API。
 * System info APIs.
 */

/** 服务端版本与健康状态。Server version and health. */
export const getSystemInfo = () => request('/api/system/info')
