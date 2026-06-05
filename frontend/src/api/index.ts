/**
 * 百宝箱 — API 请求封装
 * 统一管理后端 API 调用
 */

// 后端基础地址
// H5 开发时通过 vite proxy 代理，小程序需要填写真实地址
const BASE_URL = ''

interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
}

interface RequestOptions {
  url: string
  method?: 'GET' | 'POST'
  data?: any
  header?: Record<string, string>
}

/**
 * 通用请求方法
 */
function request<T = any>(options: RequestOptions): Promise<ApiResponse<T>> {
  return new Promise((resolve, reject) => {
    uni.request({
      url: BASE_URL + options.url,
      method: options.method || 'GET',
      data: options.data,
      header: {
        'Content-Type': 'application/json',
        ...options.header
      },
      success: (res) => {
        const data = res.data as ApiResponse<T>
        if (data.code === 0) {
          resolve(data)
        } else {
          uni.showToast({
            title: data.message || '请求失败',
            icon: 'none'
          })
          reject(data)
        }
      },
      fail: (err) => {
        console.error('请求失败:', err)
        uni.showToast({
          title: '网络异常，请稍后重试',
          icon: 'none'
        })
        reject(err)
      }
    })
  })
}

/**
 * GET 请求
 */
export function get<T = any>(url: string, data?: any): Promise<ApiResponse<T>> {
  return request<T>({ url, method: 'GET', data })
}

/**
 * POST 请求
 */
export function post<T = any>(url: string, data?: any): Promise<ApiResponse<T>> {
  return request<T>({ url, method: 'POST', data })
}

// ========== 工具 API ==========

/** JSON 格式化 */
export function jsonFormat(action: 'pretty' | 'minify' | 'valid', input: string) {
  return post('/api/tools/json-format', { action, input })
}

/** Base64 编解码 */
export function base64Convert(action: 'encode' | 'decode', input: string) {
  return post('/api/tools/base64', { action, input })
}

/** 时间戳转换 */
export function timestampConvert(action: 'now' | 'toDate' | 'toStamp', input?: string, unit?: 's' | 'ms') {
  return post('/api/tools/timestamp', { action, input, unit: unit || 's' })
}

/** 二维码生成 */
export function qrcodeGenerate(input: string, size?: number) {
  return post('/api/tools/qrcode', { input, size: size || 256 })
}

/** 颜色转换 */
export function colorConvert(input: string) {
  return post('/api/tools/color', { input })
}

/** 文本处理 */
export function textProcess(action: string, input: string) {
  return post('/api/tools/text', { action, input })
}

/** UUID 生成 */
export function uuidGenerate(count?: number, format?: string) {
  return post('/api/tools/uuid', { count: count || 1, format: format || 'default' })
}

/** 健康检查 */
export function ping() {
  return get('/api/ping')
}

export default {
  get,
  post,
  jsonFormat,
  base64Convert,
  timestampConvert,
  qrcodeGenerate,
  colorConvert,
  textProcess,
  uuidGenerate,
  ping
}
