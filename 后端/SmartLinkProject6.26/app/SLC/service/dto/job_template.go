package dto

import (
	"SmartLinkProject/app/SLC/models"
	"SmartLinkProject/common/dto"
)

// JobTemplateGetReq 结构体用于获取作业模板请求
type JobTemplateGetReq struct {
	dto.Pagination `search:"-"`
	TypeId         int `comment:"设备类型ID" form:"typeId"` // 设备类型的唯一标识
}

// JobTemplateInsertReq 结构体用于插入作业模板请求
type JobTemplateInsertReq struct {
	TemplateId             int    `json:"templateId" comment:"作业模板ID"`
	TypeId                 int    `json:"typeId" comment:"设备类型ID"`                 // 设备类型的唯一标识
	TemplateName           string `json:"templateName" comment:"模板名称"`             // 作业模板的名称
	JobDesc                string `json:"jobDesc" comment:"作业描述"`                  // 对作业的详细描述
	ScheduleType           string `json:"scheduleType" comment:"调度类型"`             // 作业的调度类型，如 "cron", "fixedRate", "fixedDelay"
	ScheduleConf           string `json:"scheduleConf" comment:"调度配置"`             // 具体的调度配置信息
	ConGenDisplay          string `json:"conGenDisplay" comment:"生成展示"`            // 用于展示的配置信息生成
	ExecutorHandler        string `json:"executorHandler" comment:"执行器处理器"`        // 执行器处理作业的处理器
	ExecutorParam          string `json:"executorParam" comment:"执行器参数"`           // 传递给执行器的参数
	ExecutorTimeout        int    `json:"executorTimeout" comment:"执行器超时时间"`       // 执行器的超时时间（单位通常为秒）
	ScheduleConfCron       string `json:"scheduleConfCron" comment:"CRON表达式"`      // CRON 表达式用于定时任务
	ScheduleConfFixRate    string `json:"scheduleConfFixRate" comment:"固定速率配置"`    // 固定速率调度的配置
	ScheduleConfFixDelay   string `json:"ScheduleConfFixDelay" comment:"固定延迟配置"`   // 固定延迟调度的配置
	GlueType               string `json:"glueType" comment:"粘合类型"`                 // 作业使用的粘合类型，可能指的是作业执行的上下文或环境
	ExecutorFailRetryCount int    `json:"executorFailRetryCount" comment:"失败重试次数"` // 作业失败后的重试次数
}

// JobTemplateUpdateReq 结构体更新请求
type JobTemplateUpdateReq struct {
	TemplateId             int    `json:"templateId" comment:"作业模板ID"`             // 作业模板的唯一标识
	TypeId                 int    `json:"typeId" comment:"设备类型ID"`                 // 设备类型的唯一标识
	TemplateName           string `json:"templateName" comment:"模板名称"`             // 作业模板的名称
	JobDesc                string `json:"jobDesc" comment:"作业描述"`                  // 对作业的详细描述
	ScheduleType           string `json:"scheduleType" comment:"调度类型"`             // 作业的调度类型，如 "cron", "fixedRate", "fixedDelay"
	ScheduleConf           string `json:"scheduleConf" comment:"调度配置"`             // 具体的调度配置信息
	ConGenDisplay          string `json:"conGenDisplay" comment:"生成展示"`            // 用于展示的配置信息生成
	ExecutorHandler        string `json:"executorHandler" comment:"执行器处理器"`        // 执行器处理作业的处理器
	ExecutorParam          string `json:"executorParam" comment:"执行器参数"`           // 传递给执行器的参数
	ExecutorTimeout        int    `json:"executorTimeout" comment:"执行器超时时间"`       // 执行器的超时时间（单位通常为秒）
	ScheduleConfCron       string `json:"scheduleConfCron" comment:"CRON表达式"`      // CRON 表达式用于定时任务
	ScheduleConfFixRate    string `json:"scheduleConfFixRate" comment:"固定速率配置"`    // 固定速率调度的配置
	ScheduleConfFixDelay   string `json:"ScheduleConfFixDelay" comment:"固定延迟配置"`   // 固定延迟调度的配置
	GlueType               string `json:"glueType" comment:"粘合类型"`                 // 作业使用的粘合类型，可能指的是作业执行的上下文或环境
	ExecutorFailRetryCount int    `json:"executorFailRetryCount" comment:"失败重试次数"` // 作业失败后的重试次数
}

func (j *JobTemplateInsertReq) Generate(model *models.JobTemplate) {
	model.TypeId = j.TypeId
	model.TemplateName = j.TemplateName
	model.JobDesc = j.JobDesc
	model.ScheduleType = j.ScheduleType
	model.ScheduleConf = j.ScheduleConf
	model.ConGenDisplay = j.ConGenDisplay
	model.ExecutorHandler = j.ExecutorHandler
	model.ExecutorParam = j.ExecutorParam
	model.ExecutorTimeout = j.ExecutorTimeout
	model.ScheduleConfCron = j.ScheduleConfCron
	model.ScheduleConfFixRate = j.ScheduleConfFixRate
	model.ScheduleConfFixDelay = j.ScheduleConfFixDelay
	model.GlueType = j.GlueType
	model.ExecutorFailRetryCount = j.ExecutorFailRetryCount
}

func (j *JobTemplateUpdateReq) Generate(model *models.JobTemplate) {
	model.TemplateId = j.TemplateId
	model.TypeId = j.TypeId
	model.TemplateName = j.TemplateName
	model.JobDesc = j.JobDesc
	model.ScheduleType = j.ScheduleType
	model.ScheduleConf = j.ScheduleConf
	model.ConGenDisplay = j.ConGenDisplay
	model.ExecutorHandler = j.ExecutorHandler
	model.ExecutorParam = j.ExecutorParam
	model.ExecutorTimeout = j.ExecutorTimeout
	model.ScheduleConfCron = j.ScheduleConfCron
	model.ScheduleConfFixRate = j.ScheduleConfFixRate
	model.ScheduleConfFixDelay = j.ScheduleConfFixDelay
	model.GlueType = j.GlueType
	model.ExecutorFailRetryCount = j.ExecutorFailRetryCount
}

// GetId 方法用于获取设备类型的ID
func (d *JobTemplateGetReq) GetId() interface{} {
	return d.TypeId
}

// GetId 获取数据对应的ID
func (d *JobTemplateUpdateReq) GetId() interface{} {
	return d.TemplateId
}

// GetId 获取数据对应的ID
func (d *JobTemplateInsertReq) GetId() interface{} {
	return d.TemplateId
}

type JobTemplateDeleteReq struct {
	TemplateIds []int `json:"template_ids"`
}

func (d *JobTemplateDeleteReq) GetId() interface{} {
	return d.TemplateIds
}

type TemplateById struct {
	dto.ObjectById
}

func (s *TemplateById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}
