package service

import (
	"SmartLinkProject/app/SLC/models"
	"SmartLinkProject/app/SLC/service/dto"
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"SmartLinkProject/common/actions"
	cDto "SmartLinkProject/common/dto"
)

type OperatorLog struct {
	service.Service
}

func (e *OperatorLog) GetPage(c *dto.OperatorLogGetPageReq, p *actions.DataPermission, list *[]models.OperatorLog, count *int64) error {
	var err error
	var data models.OperatorLog

	err = e.Orm.Debug().
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

func (e *OperatorLog) Get(d *dto.OperatorLogById, p *actions.DataPermission, model *models.OperatorLog) error {
	var data models.OperatorLog

	err := e.Orm.Model(&data).Debug().
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

func (e *OperatorLog) Insert(c *dto.OperatorLogInsertReq) error {
	var err error
	var data models.OperatorLog
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

func (e *OperatorLog) Remove(c *dto.OperatorLogById, p *actions.DataPermission) error {
	var err error
	var data models.OperatorLog

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Error found in  RemoveOperatorLog : %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
