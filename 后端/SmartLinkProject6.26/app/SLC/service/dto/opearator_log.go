package dto

import (
	"SmartLinkProject/app/SLC/models"
	"SmartLinkProject/common/dto"
	common "SmartLinkProject/common/models"
)

type OperatorLogGetPageReq struct {
	dto.Pagination `search:"-"`
	OplId          int    `form:"oplId" search:"type:exact;column:opl_id;table:operator_log" comment:""`
	JobId          int    `form:"jobId" search:"type:exact;column:job_id;table:operator_log" comment:""`
	Name           string `form:"name" search:"type:exact;column:name;table:operator_log" comment:""`
	Tel            string `form:"tel" search:"type:exact;column:tel;table:operator_log" comment:""`
	Operator       string `form:"operator" search:"type:exact;column:operator;table:operator_log" comment:""`
}

func (m *OperatorLogGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type OperatorLogInsertReq struct {
	OplId    int    `json:"oplId" comment:"操作日志ID"` // 操作日志ID
	JobId    int    `json:"jobId" comment:"任务ID"`   // 任务ID
	Name     string `json:"name" `
	Tel      string `json:"tel" `
	Operator string `json:"operator" `
}

func (d *OperatorLogInsertReq) Generate(model *models.OperatorLog) {
	if d.OplId != 0 {
		model.OplId = d.OplId
	}
	model.JobId = d.JobId
	model.Name = d.Name
	model.Tel = d.Tel
	model.Operator = d.Operator
}

func (s *OperatorLogInsertReq) GetId() interface{} {
	return s.OplId
}

type OperatorLogById struct {
	dto.ObjectById
}

func (s *OperatorLogById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

func (s *OperatorLogById) GenerateM() (common.ActiveRecord, error) {
	return &models.OperatorLog{}, nil
}
