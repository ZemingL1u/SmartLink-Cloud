import request from '@/utils/request'

// 查询Device列表
export function listCategory(params) {
  return request({
    url: '/api/v1/devicetype',
    method: 'get',
    params: params
  })
}

export function getCateList(query) {
  return request({
    url: '/api/v1/devicetype',
    method: 'get',
    params: query
  })
}

export function addCate(data) {
  return request({
    url: '/api/v1/devicetype',
    method: 'post',
    data: data
  })
}

export function updateCate(data,TypeId) {
  return request({
    url: '/api/v1/devicetype/' + TypeId,
    method: 'put',
    data: data
  })
}

export function delCate(data) {
  return request({
    url: '/api/v1/devicetype',
    method: 'delete',
    data: data
  })
}

