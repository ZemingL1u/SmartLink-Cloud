package dto

import (
	"SmartLinkProject/app/SLC/models"
)

type SceneDeviceInsertReq struct {
	SceneId   int    `json:"sceneId" comment:"任务ID"` // 任务ID
	DeviceId  int    `json:"deviceId" comment:"设备ID" vd:"$>0" `
	ElementId string `json:"elementId" comment:"图元ID" `
}

func (d *SceneDeviceInsertReq) Generate(model *models.SceneDevice) {
	if d.SceneId != 0 {
		model.SceneId = d.SceneId
	}
	model.ElementId = d.ElementId
	model.DeviceId = d.DeviceId
}

func (s *SceneDeviceInsertReq) GetId() interface{} {
	return s.DeviceId
}

type SceneDeviceGetReq struct {
	SceneId   int    `json:"sceneId" comment:"任务ID"` // 任务ID
	ElementId string `json:"elementId" comment:"图元ID" `
}

func (d *SceneDeviceGetReq) Generate(model *models.SceneDevice) {
	model.SceneId = d.SceneId
	model.ElementId = d.ElementId
}

type SceneDeviceRemoveReq struct {
	SceneId   int    `json:"sceneId" comment:"任务ID"` // 任务ID
	ElementId string `json:"elementId" comment:"图元ID" `
}

func (d *SceneDeviceRemoveReq) Generate(model *models.SceneDevice) {
	model.ElementId = d.ElementId
	model.SceneId = d.SceneId
}

func (s *SceneDeviceRemoveReq) GetId() interface{} {
	return s.SceneId
}
