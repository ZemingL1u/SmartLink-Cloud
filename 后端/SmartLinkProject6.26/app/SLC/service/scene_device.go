package service

import (
	"SmartLinkProject/app/SLC/models"
	"SmartLinkProject/app/SLC/service/dto"
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"SmartLinkProject/common/actions"
)

type SceneDevice struct {
	service.Service
}

func (e *SceneDevice) Get(c *dto.SceneDeviceGetReq, p *actions.DataPermission, model *models.SceneDevice) error {
	// 复制请求参数到模型，设置查询条件
	c.Generate(model)

	// 使用ORM查询数据库
	err := e.Orm.Model(model).Debug().
		Scopes(
			actions.Permission(model.TableName(), p),
		).
		Where("scene_id = ? AND element_id = ?", model.SceneId, model.ElementId).
		First(model).Error

	// 检查查询结果
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("查看对象不存在或无权查看")
		}
		e.Log.Errorf("db error: %s", err.Error())
		return err
	}

	return nil
}

func (e *SceneDevice) Insert(c *dto.SceneDeviceInsertReq) error {
	var err error
	var data models.SceneDevice
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err.Error())
		return err
	}
	return nil
}

// Remove方法需要根据SceneId和ElementId来删除记录
func (e *SceneDevice) Remove(c *dto.SceneDeviceRemoveReq, p *actions.DataPermission) error {
	var err error
	var data models.SceneDevice

	// 使用Generate方法设置model的字段
	c.Generate(&data)

	// 构建删除查询，这里使用Where代替Delete来定位记录
	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Where("scene_id = ? AND element_id = ?", data.SceneId, data.ElementId).Delete(&data)

	if err = db.Error; err != nil {
		e.Log.Errorf("Error found in RemoveSceneDevice: %s", err.Error())
		return err
	}

	// 检查是否有记录被删除
	if db.RowsAffected == 0 {
		e.Log.Errorf("Error found in RemoveSceneDevice: %s", err.Error())
		return err
	}

	return nil
}
