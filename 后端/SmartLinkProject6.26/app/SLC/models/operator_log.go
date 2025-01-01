package models

import (
	"SmartLinkProject/common/models"
	// "gorm.io/gorm"
)

type OperatorLog struct {
	OplId    int    `json:"oplId" gorm:"primaryKey;autoIncrement;comment:任务ID"`
	JobId    int    `json:"jobId"`
	Name     string `json:"name" gorm:"comment:操作人名字" `
	Tel      string `json:"tel" gorm:"comment:电话"`
	Operator string `json:"operator" gorm:"comment:操作内容"`

	models.ModelTime
}

func (*OperatorLog) TableName() string {
	return "operator_log"
}

func (e *OperatorLog) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OperatorLog) GetId() interface{} {
	return e.OplId
}
