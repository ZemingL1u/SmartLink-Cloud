package models

import "SmartLinkProject/common/models"

// Scene 场景模型
type Scene struct {
	models.ModelTime
	SceneID   int    `gorm:"primaryKey;autoIncrement;comment:场景ID" json:"sceneId"`
	SceneName string `gorm:"size:255;comment:场景名称" json:"sceneName"`
	DeviceNum int    `gorm:"comment:设备数量" json:"deviceNum"`
	LocName   string `gorm:"size:255;comment:位置名称" json:"locName"`
	LocID     int    `gorm:"comment:位置ID" json:"locId"`
	ImgSrc    string `gorm:"size:255;comment:图片来源" json:"imgSrc"`
	DataSrc   string `gorm:"size:255;comment:数据来源" json:"dataSrc"`
	Data      string `gorm:"-" json:"data"`
	Image     string `gorm:"-" json:"image"`
}

func (*Scene) TableName() string {
	return "scene"
}

func (e *Scene) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Scene) GetId() interface{} {
	return e.SceneID
}
