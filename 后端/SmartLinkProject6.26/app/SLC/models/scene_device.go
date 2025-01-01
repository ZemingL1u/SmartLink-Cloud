package models

import (
	"SmartLinkProject/common/models"
	// "gorm.io/gorm"
)

type SceneDevice struct {
	SceneId   int    `json:"sceneId" gorm:"comment:任务ID"`
	ElementId string `json:"elementId" gorm:"comment:设备ID"`
	DeviceId  int    `json:"deviceId" gorm:"comment:设备ID"`

	models.ModelTime
}

func (*SceneDevice) TableName() string {
	return "scene_device"
}

func (e *SceneDevice) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SceneDevice) GetId() interface{} {
	return e.DeviceId
}
