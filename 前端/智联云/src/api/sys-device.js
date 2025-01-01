import request from '@/utils/request'

// 查询Device列表
export function listDevice(params) {
  return request({
    url: '/api/v1/device',
    method: 'get',
    params: params
  })
}

export function listDevices(query) {
  return request({
    url: '/api/v1/device',
    method: 'post',
    params: query
  })
}

export function addDevice(query) {
  return request({
    url: '/api/v1/device',
    method: 'post',
    data: query
  })
}

export function getDeviceIdByLocId(params) {
  return request({
    url: '/api/v1/device/getDeviceIdByLocId',
    method: 'get',
    params: params
  })
}

export function getlistByDeviceIds(params) {
  return request({
    url: '/api/v1/device/getlistByDeviceIds',
    method: 'post',
    data: params
  })
}

