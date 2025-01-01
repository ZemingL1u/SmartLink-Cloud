package dto

import (
	"SmartLinkProject/app/SLC/models"
	"SmartLinkProject/common/dto"
	common "SmartLinkProject/common/models"
)

// SceneGetPageReq 分页查询场景请求参数
type SceneGetPageReq struct {
	dto.Pagination `search:"-"`
	SceneId        int    `json:"sceneId" search:"type:exact;column:scene_id;table:scene"`
	LocIdName      string `json:"locName" search:"type:contains;column:loc_name;table:scene"`
	LocId          int    `json:"locId" search:"type:exact;column:loc_id;table:scene"`
	SceneName      string `json:"sceneName" search:"type:contains;column:scene_name;table:scene"`
	DeviceNum      int    `json:"deviceNum" search:"type:exact;column:device_num;table:scene"`
}

func (m *SceneGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// SceneGetPageReqBySceneIds 通过场景ID列表获取场景列表的请求参数
type SceneGetPageReqBySceneIds struct {
	dto.Pagination `search:"-"`
	SceneIds       []int `json:"sceneIds" binding:"required"` // 场景ID列表
}

// SceneInsertReq 新增场景请求参数
type SceneInsertReq struct {
	SceneId   int    `json:"sceneId"`   // 场景ID
	LocId     int    `json:"locId"`     // 位置ID
	LocName   string `json:"locName"`   // 位置ID
	SceneName string `json:"sceneName"` // 场景名称
	ImgSrc    string `json:"imgSrc"`    // 图片来源
	DataSrc   string `json:"dataSrc"`   // 数据来源
	DeviceNum int    `json:"deviceNum"` // 设备数量
}

func (d *SceneInsertReq) Generate(model *models.Scene) {
	model.LocID = d.LocId
	model.SceneName = d.SceneName
	model.LocName = d.LocName
	model.DeviceNum = 0
}

func (s *SceneInsertReq) GetId() interface{} {
	return s.SceneId
}

// SceneUpdateReq 更新场景请求参数
type SceneUpdateReq struct {
	SceneId   int    `json:"sceneId"`   // 场景ID
	LocId     int    `json:"locId"`     // 位置ID
	LocName   string `json:"locName"`   // 位置ID
	SceneName string `json:"sceneName"` // 场景名称
	DeviceNum int    `json:"deviceNum"` // 设备数量
}

func (d *SceneUpdateReq) Generate(model *models.Scene) {
	if d.SceneId != 0 {
		model.SceneID = d.SceneId
	}
	model.LocID = d.LocId
	model.SceneName = d.SceneName
	model.LocName = d.LocName
	model.DeviceNum = d.DeviceNum
}
func (d *SceneUpdateReq) GetId() interface{} {
	return d.SceneId
}

type SceneById struct {
	dto.ObjectById
}

func (s *SceneById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

func (s *SceneById) GenerateM() (common.ActiveRecord, error) {
	return &models.Scene{}, nil
}
