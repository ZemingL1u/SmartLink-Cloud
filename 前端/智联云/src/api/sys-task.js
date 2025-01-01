import request from '@/utils/request'

// 查询id封装成ids
export function listfirstTask(params) {
  return request({
    url: '/api/v1/jobinfo/pageList',
    method: 'get',
    params: params
  })
}

// 查询ids对应的任务列表
export function listsecondTask(params) {
  return request({
    url: 'http://120.79.59.158:8080/xxl-job-admin/jobinfo/pageLists',
    method: 'post',
    params: params
  })
}

// 启动任务
export function JobStart(params) {
  return request({
    url: 'http://120.79.59.158:8080/xxl-job-admin/jobinfo/starts',
    method: 'post',
    params: params
  })
}

// 停止任务
export function JobStop(params) {
  return request({
    url: 'http://120.79.59.158:8080/xxl-job-admin/jobinfo/stops',
    method: 'post',
    params: params
  })
}

// 单次执行
export function Single(params) {
  return request({
    url: 'http://120.79.59.158:8080/xxl-job-admin/jobinfo/triggers',
    method: 'post',
    params: params
  })
}

// 获取操作日志
export function GetLogs(params) {
  return request({
    url: '/api/v1/operatorLog/pageList',
    method: 'get',
    params: params
  })
}

// 删除任务
export function JobfirstRemove(data) {
  return request({
    url: '/api/v1/jobinfo',
    method: 'delete',
    data: data
  })
}
export function JobsecondRemove(params) {
  return request({
    url: 'http://120.79.59.158:8080/xxl-job-admin/jobinfo/removes',
    method: 'post',
    params: params
  })
}

// 删除日志
export function logremove(data) {
  return request({
    url: '/api/v1/operatorLog',
    method: 'delete',
    data: data
  })
}
// 增加日志
export function logadd(data) {
  return request({
    url: '/api/v1/operatorLog',
    method: 'post',
    data: data
  })
}

export function TreeChild(params) {
  return request({
    url: '/api/v1/jobinfo/getJobIdByLocId',
    method: 'get',
    params: params
  })
}

// 获取下次执行时间
export function Findtime(params){
  return request({
    url: 'http://120.79.59.158:8080/xxl-job-admin/jobinfo/nextTriggerTime',
    method: 'post',
    params: params
  })
}
