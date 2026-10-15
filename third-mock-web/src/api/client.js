import { useServerStore } from '../stores/server'

/**
 * 管理API 请求封装：统一 JSON 编解码与错误抛出。
 * Admin API request wrapper: unified JSON codec and error throwing.
 *
 * 说明：管理 API 契约（路径/字段）尚未冻结，具体接口随服务端实现同步调整。
 * Note: the admin API contract (paths/fields) is not finalized; endpoints will
 * be adjusted in sync with the server implementation.
 */

/**
 * 发起管理 API 请求。
 * Perform an admin API request.
 *
 * @param {string} path 以 / 开头的 API 路径（如 /api/rules）
 * @param {object} [options] fetch 配置，body 为对象时自动 JSON 序列化
 * @returns {Promise<any>} 服务端返回的 JSON 数据
 */
export async function request(path, options = {}) {
  const { baseUrl } = useServerStore()
  const { body, headers, ...rest } = options
  const resp = await fetch(baseUrl + path, {
    headers: {
      'Content-Type': 'application/json',
      ...headers,
    },
    body: body !== undefined ? JSON.stringify(body) : undefined,
    ...rest,
  })
  if (!resp.ok) {
    throw new Error(`请求失败 ${resp.status}: ${await resp.text()}`)
  }
  return resp.json()
}
