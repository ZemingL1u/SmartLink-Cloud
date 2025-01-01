import request from '@/utils/request'

export function getDeptList(query) {
  return request({
    url: '/api/v1/loc',
    method: 'get',
    params: query
  })
}

// 查询部门详细
export function getDept(locId) {
  return request({
    url: '/api/v1/loc/' + locId,
    method: 'get'
  })
}

// 查询部门下拉树结构
export function treeselect() {
  return request({
    url: '/api/v1/locTree',
    method: 'get'
  })
}

export function typetreeselect() {
  return request({
    url: '/api/v1/typeTree',
    method: 'get'
  })
}

// 根据角色ID查询部门树结构
export function roleDeptTreeselect(roleId) {
  return request({
    url: '/api/v1/roleLocTreeselect/' + roleId,
    method: 'get'
  })
}

// 新增部门
export function addDept(data) {
  return request({
    url: '/api/v1/loc',
    method: 'post',
    data: data
  })
}

// 修改部门
export function updateDept(data, id) {
  return request({
    url: '/api/v1/loc/' + id,
    method: 'put',
    data: data
  })
}

// 删除部门
export function delDept(data) {
  return request({
    url: '/api/v1/loc',
    method: 'delete',
    data: data
  })
}
