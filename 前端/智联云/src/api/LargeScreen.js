import request from "@/utils/request";

export function getTempData() {
  return request({
    url: '/api/v1/device/getTempData',
    method: 'get'
  })
}

export function getHumiData() {
  return request({
    url: '/api/v1/device/getHumiData',
    method: 'get'
  })
}
