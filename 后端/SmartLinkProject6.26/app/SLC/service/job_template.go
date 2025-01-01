package service

import (
	"SmartLinkProject/app/SLC/models"
	"SmartLinkProject/app/SLC/service/dto"
	"SmartLinkProject/common/actions"
	"errors"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/service"
)

type JobTemplate struct {
	service.Service
}

// GetByTypeId 方法用于根据 TypeId 查询作业模板信息
func (e *JobTemplate) GetByTypeId(d *dto.JobTemplateGetReq, p *actions.DataPermission, model *[]models.JobTemplate, count *int64) error {
	// 使用 GORM 构建查询，查找所有 TypeId 匹配的记录
	var data []models.JobTemplate // 声明一个用于接收查询结果的切片
	err := e.Orm.Model(&data).Debug().
		Scopes(
			actions.Permission("job_template", p),
		).Where("type_id = ?", d.GetId()).Find(&data).Count(count).Error // 将查询结果赋值给 data

	// 检查查询是否出错
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	// 将查询结果赋值给传入的 model 指针指向的切片
	*model = data
	return nil
}

// Insert 创建JobTemplate对象
func (e *JobTemplate) Insert(c *dto.JobTemplateInsertReq) error {
	var err error
	var data models.JobTemplate

	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Update 修改JobTemplate对象
func (e *JobTemplate) Update(c *dto.JobTemplateUpdateReq, p *actions.DataPermission) error {
	var err error
	var model models.JobTemplate
	db := e.Orm.First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateJobTemplate error: %s", err)
		return err
	}

	c.Generate(&model)
	update := e.Orm.Model(&model).Where("template_id = ?", &model.TemplateId).Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update jobtemplate error")
		log.Warnf("db update error")
		return err
	}
	return nil
}

// Remove 删除JobTemplate对象
func (e *JobTemplate) Remove(c *dto.TemplateById, p *actions.DataPermission) error {
	var err error
	var data models.JobTemplate

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission("job_template", p),
		).Delete(&data, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Error found in  RemoveJobTemplate : %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
