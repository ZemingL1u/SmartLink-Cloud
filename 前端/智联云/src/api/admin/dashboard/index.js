import request from '@/utils/request'

export function home() {
  return request({
    url: '/api/v1/home',
    method: 'get'
  })
}

// 折线图接口
export function line() {
    return request({
      url: '/api/v1/line',
      method: 'get'
    })
  }