package models

import "SmartLinkProject/common/models"

type JobTemplate struct {
	TemplateId             int    `json:"templateId" gorm:"primaryKey;autoIncrement"`    // 主键，自动递增
	TypeId                 int    `json:"typeId" gorm:"comment:类型ID"`                    // 类型ID
	TemplateName           string `json:"templateName" gorm:"comment:模版名称"`              // 模版名称
	JobDesc                string `json:"jobDesc" gorm:"type:varchar(255);comment:作业描述"` // 作业描述
	ScheduleType           string `json:"scheduleType" gorm:"comment:调度类型"`              // 调度类型
	ScheduleConf           string `json:"scheduleConf" gorm:"type:text;comment:调度配置"`    // 调度配置
	ConGenDisplay          string `json:"conGenDisplay" gorm:"type:text;comment:生成展示"`   // 生成展示
	ExecutorHandler        string `json:"executorHandler" gorm:"comment:执行器处理器"`         // 执行器处理器
	ExecutorParam          string `json:"executorParam" gorm:"type:text;comment:执行器参数"`  // 执行器参数
	ExecutorTimeout        int    `json:"executorTimeout" gorm:"comment:执行器超时时间"`        // 执行器超时时间（单位：秒）
	ScheduleConfCron       string `json:"scheduleConfCron" gorm:"comment:CRON表达式"`       // CRON表达式
	ScheduleConfFixRate    string `json:"scheduleConfFixRate" gorm:"comment:固定速率配置"`     // 固定速率配置
	ScheduleConfFixDelay   string `json:"scheduleConfFixDelay" gorm:"comment:固定延迟配置"`    // 固定延迟配置
	GlueType               string `json:"glueType" gorm:"comment:粘合类型"`                  // 粘合类型
	ExecutorFailRetryCount int    `json:"executorFailRetryCount" gorm:"comment:失败重试次数"`  // 失败重试次数

}

func (*JobTemplate) TableName() string {
	return "job_template"
}

func (e *JobTemplate) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *JobTemplate) GetId() interface{} {
	return e.TemplateId
}
