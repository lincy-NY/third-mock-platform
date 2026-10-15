import { request } from './client'

/**
 * 规则相关 API：契约待冻结，先按 REST 风格占位。
 * Rule APIs: contract not finalized; REST-style placeholders for now.
 */

/** 规则列表。List rules. */
export const listRules = () => request('/api/rules')

/** 创建规则。Create a rule. */
export const createRule = (rule) => request('/api/rules', { method: 'POST', body: rule })

/** 更新规则。Update a rule. */
export const updateRule = (id, rule) => request(`/api/rules/${id}`, { method: 'PUT', body: rule })

/** 删除规则。Delete a rule. */
export const deleteRule = (id) => request(`/api/rules/${id}`, { method: 'DELETE' })
