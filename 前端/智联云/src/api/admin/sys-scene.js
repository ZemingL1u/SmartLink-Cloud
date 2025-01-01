import request from '@/utils/request'

// 查询SysApi列表
export function lisSceneApi(query) {
  return request({
    url: '/api/v1/scene',
    method: 'get',
    params: query
  })
}
//新增
export function addSceneApi(query,token) {
  return request({
    url: `/api/v1/scene?token=${token}`,
    method: 'post',
    data: query
  })
}

// 修改
export function updateSceneApi(query,token) {
  return request({
    url: `/api/v1/scene?token=${token}`,
    method: 'put',
    data: query
  })
}

// 删除设备
export function delSence(data) {
  return request({
    url: '/api/v1/scene',
    method: 'delete',
    data: data
  })
}
