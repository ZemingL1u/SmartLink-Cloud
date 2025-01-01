package service

import (
	"SmartLinkProject/app/SLC/models"
	"SmartLinkProject/app/SLC/service/dto"
	"SmartLinkProject/mqtt"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"SmartLinkProject/common/actions"
	cDto "SmartLinkProject/common/dto"
)

type Device struct {
	service.Service
}

// GetPage 获取Device列表
func (e *Device) GetPage(c *dto.DeviceGetPageReq, p *actions.DataPermission, list *[]models.Device, count *int64) error {
	var err error
	var data models.Device

	err = e.Orm.Debug().
		Preload("Loc").
		Preload("Type").
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

// Get 获取Device对象
func (e *Device) Get(d *dto.DeviceById, p *actions.DataPermission, model *models.Device) error {
	var data models.Device

	err := e.Orm.Model(&data).Debug().Preload("Loc").Preload("Type").
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

	// 检查文件路径是否存在
	if _, err := os.Stat(model.ImageSrc); os.IsNotExist(err) {
		e.Log.Warnf("The directory does not exist: %s", model.ImageSrc)
		model.Image = []string{} // 设置为空切片
		return nil               // 直接返回，不进行遍历
	}
	err = filepath.Walk(model.ImageSrc, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // 处理错误
		}
		if !info.IsDir() { // 确保是文件
			// 读取图片文件并进行Base64编码
			file, err := os.Open(path)
			if err != nil {
				model.ImageSrc = ""
				defer file.Close()
				return nil
			}
			defer file.Close()

			// 创建一个限流的 reader 来读取文件，以避免一次性将整个文件读入内存
			reader := io.LimitReader(file, 10*1024*1024) // 限制文件大小为10MB

			// 读取文件内容
			imageData, err := io.ReadAll(reader)
			if err != nil {
				e.Log.Errorf("failed to read image file: %s", err)
				return err
			}

			// 进行Base64编码
			encodedImage := base64.StdEncoding.EncodeToString(imageData)
			model.Image = append(model.Image, encodedImage) // 添加到切片中
		}
		return nil
	})

	if err != nil {
		e.Log.Errorf("failed to walk through the directory: %s", err)
		return err
	}
	return nil
}

// Insert 创建Device对象
func (e *Device) Insert(c *dto.DeviceInsertReq) error {
	var err error
	var data models.Device
	var i int64
	err = e.Orm.Model(&data).Where("device_name = ?", c.DeviceName).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("设备名已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	data.ImageSrc = "static/deviceuploadfile/" + strconv.Itoa(data.DeviceId)
	update := e.Orm.Model(&data).Where("device_id = ?", &data.DeviceId).Updates(&data)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update device error")
		e.Log.Warnf("db update error")
		return err
	}
	c.DeviceId = data.DeviceId

	topic := "getData/" + strconv.Itoa(data.DeviceId)
	// 订阅MQTT主题
	mqtt.SubscribeToTopic(topic)

	return nil
}

// Update 修改Device对象
func (e *Device) Update(c *dto.DeviceUpdateReq, p *actions.DataPermission) error {
	var err error
	var model models.Device
	db := e.Orm.First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateDevice error: %s", err)
		return err
	}

	c.Generate(&model)
	update := e.Orm.Model(&model).Where("device_id = ?", &model.DeviceId).Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update device error")
		log.Warnf("db update error")
		return err
	}
	return nil
}

// UpdateStatus 更新设备状态
func (e *Device) UpdateStatus(c *dto.UpdateDeviceStatusReq, p *actions.DataPermission) error {
	var err error
	var model models.Device
	db := e.Orm.First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateDevice error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	err = e.Orm.Table(model.TableName()).Where("device_id =? ", c.DeviceId).Updates(c).Error
	if err != nil {
		e.Log.Errorf("Service UpdateDevice error: %s", err)
		return err
	}
	return nil
}

// Remove 删除Device对象
func (e *Device) Remove(c *dto.DeviceById, p *actions.DataPermission) error {
	var err error
	var data models.Device

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Error found in RemoveDevice: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}

	ids := c.GetId()
	// 检查获取的 ID 是单个 ID 还是 ID 切片
	switch v := ids.(type) {
	case int:
		// 构建单个 ID 的文件夹路径并删除文件夹
		folderPath := "static/deviceuploadfile/" + fmt.Sprint(v) // 假设文件夹以设备ID命名
		// 删除文件夹及其内容
		err = os.RemoveAll(folderPath)
		if err != nil {
			e.Log.Errorf("Failed to delete folder %s: %s", folderPath, err)
			return err
		}
		topic := "getData/" + strconv.Itoa(v)
		mqtt.UnsubscribeFromTopic(topic)
	case []int:
		// 如果是 ID 切片，遍历每个 ID 删除对应的设备记录和文件夹
		for _, id := range v {
			// 构建文件夹路径
			folderPath := "static/deviceuploadfile/" + fmt.Sprint(id) // 假设文件夹以设备ID命名
			// 删除文件夹及其内容
			err = os.RemoveAll(folderPath)
			if err != nil {
				e.Log.Errorf("Failed to delete folder %s: %s", folderPath, err)
				return err
			}

			topic := "getData/" + strconv.Itoa(id)
			mqtt.UnsubscribeFromTopic(topic)
		}
	default:
		return errors.New("无法识别的 ID 类型")
	}

	return nil
}

func (e *Device) GetProfile(c *dto.DeviceById, device *models.Device) error {
	err := e.Orm.Preload("Loc").Preload("Type").First(device, c.GetId()).Error
	if err != nil {
		return err
	}
	return nil
}

// 根据父区域id得到该区域下所有任务id
func (e *Device) GetDeviceIdsByLocId(c *dto.DLocJoin) ([]int, error) {
	locParentId := c.LocId
	sysLoc := &SysLoc{
		Service: service.Service{
			Orm: e.Orm,
			Log: e.Log,
		},
	}

	allLocIDs, err := sysLoc.getAllChildLocIDs(locParentId)
	if err != nil {
		e.Log.Errorf("Error getting all child location IDs: %v", err)
		return nil, err
	}
	allLocIDs = append(allLocIDs, locParentId)

	var deviceIds []int
	Devices := make([]models.Device, 0)
	if err := e.Orm.Model(&models.Device{}).Where("loc_id IN (?)", allLocIDs).Find(&Devices).Error; err != nil {
		e.Log.Errorf("db error: %v", err)
		return nil, err
	}

	for _, device := range Devices {
		deviceIds = append(deviceIds, device.DeviceId)
	}

	return deviceIds, nil
}

// GetPageByDeviceIds 根据提供的设备ID列表和分页信息获取设备数据
func (e *Device) GetPageByDeviceIds(deviceIds []int, pageReq *dto.DeviceGetPageReqByDeviceIds, p *actions.DataPermission, list *[]models.Device, count *int64) error {
	// 检查deviceIds是否为空
	if len(deviceIds) == 0 {
		return errors.New("设备ID列表为空")
	}

	// 检查pageReq是否符合要求
	if pageReq == nil {
		return errors.New("分页请求对象为空")
	}

	// 使用传入的设备ID列表和分页请求查询设备数据
	err := e.Orm.Debug().
		Where("device_id IN (?)", deviceIds).
		Scopes(
			cDto.Paginate(pageReq.GetPageSize(), pageReq.GetPageIndex()), // 应用分页
			actions.Permission("device", p),                              // 应用数据权限
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	return nil
}
