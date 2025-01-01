package service

import (
	"SmartLinkProject/app/SLC/models"
	"SmartLinkProject/app/SLC/service/dto"
	"encoding/base64"
	"errors"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"SmartLinkProject/common/actions"
	cDto "SmartLinkProject/common/dto"
)

type Scene struct {
	service.Service
}

func (e *Scene) encodeImageToBase64(scene *models.Scene) error {
	imagePath := scene.ImgSrc // 图片文件的路径
	file, err := os.Open(imagePath)
	if err != nil {
		scene.ImgSrc = "" // 如果文件打开失败，可以设置为空或者错误信息
		return nil
	}
	defer file.Close()

	// 创建一个限流的reader来读取文件，以避免一次性将整个文件读入内存
	reader := io.LimitReader(file, 10*1024*1024) // 限制文件大小为10MB

	// 读取文件内容
	imageData, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	// 进行Base64编码
	scene.Image = base64.StdEncoding.EncodeToString(imageData)
	return nil
}

// copyFile copies a file from src to dst
func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

// GetPage 获取Scene列表
func (e *Scene) GetPage(c *dto.SceneGetPageReq, p *actions.DataPermission, list *[]models.Scene, count *int64) error {
	var err error
	var data models.Scene

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
	// 再为每个场景的图片生成Base64编码
	for i := range *list {
		if err := e.encodeImageToBase64(&(*list)[i]); err != nil {
			e.Log.Errorf("error encoding image to base64: %s", err)
		}
	}
	return nil
}

// Get 获取Scene对象
func (e *Scene) Get(d *dto.SceneById, p *actions.DataPermission, model *models.Scene) error {
	var data models.Scene

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

	// 读取 JSON 文件
	jsonFilePath := model.DataSrc
	fileContent, err := os.ReadFile(jsonFilePath)
	if err != nil {
		log.Printf("failed to read JSON file: %s", err)
		return err
	}

	// 将 JSON 文件内容转换为字符串，并赋值给 model.DataSrc
	model.Data = string(fileContent)

	return nil
}

// Insert 创建Scene对象
func (e *Scene) Insert(c *dto.SceneInsertReq) error {
	var err error
	var data models.Scene
	var i int64
	err = e.Orm.Model(&data).Where("scene_name = ?", c.SceneName).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("场景名已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	data.ImgSrc = "static/scene/image/" + strconv.Itoa(data.SceneID) + ".png"
	data.DataSrc = "static/scene/data/" + strconv.Itoa(data.SceneID) + ".json"

	// 为新创建的场景生成默认文件路径
	defaultImagePath := "static/scene/image/0.png"
	defaultDataPath := "static/scene/data/0.json"
	// 复制默认图片文件到新路径
	if err := copyFile(defaultImagePath, data.ImgSrc); err != nil {
		e.Log.Errorf("failed to copy image file: %s", err)
		return err
	}

	// 复制默认数据文件到新路径
	if err := copyFile(defaultDataPath, data.DataSrc); err != nil {
		e.Log.Errorf("failed to copy data file: %s", err)
		return err
	}

	update := e.Orm.Model(&data).Where("scene_id = ?", &data.SceneID).Updates(&data)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update Scene error")
		e.Log.Warnf("db update error")
		return err
	}
	c.SceneId = data.SceneID
	return nil
}

// Update 修改Scene对象
func (e *Scene) Update(c *dto.SceneUpdateReq, p *actions.DataPermission) error {
	var err error
	var model models.Scene
	db := e.Orm.First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateScene error: %s", err)
		return err
	}
	c.Generate(&model)
	update := e.Orm.Model(&model).Where("scene_id = ?", &model.SceneID).Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update Scene error")
		e.Log.Warnf("db update error")
		return err
	}

	return nil
}

// Remove 删除Scene对象
func (e *Scene) Remove(c *dto.SceneById, p *actions.DataPermission) error {
	var err error
	var data models.Scene

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Error found in RemoveScene: %s", err)
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
		dataPath := "static/scene/data/" + strconv.Itoa(v) + ".json"  //
		imagePath := "static/scene/image/" + strconv.Itoa(v) + ".png" //
		// 删除文件夹及其内容
		err = os.Remove(dataPath)
		if err != nil {
			e.Log.Errorf("Failed to delete file %s: %s", dataPath, err)
		}
		err = os.Remove(imagePath)
		if err != nil {
			e.Log.Errorf("Failed to delete file %s: %s", imagePath, err)
		}
	case []int:
		// 如果是 ID 切片，遍历每个 ID 删除对应的设备记录和文件夹
		for _, id := range v {
			// 构建单个 ID 的文件夹路径并删除文件夹
			dataPath := "static/scene/data" + strconv.Itoa(id) + ".json"   //
			imagePath := "static/scene/image/" + strconv.Itoa(id) + ".png" //
			// 删除文件夹及其内容
			err = os.Remove(dataPath)
			if err != nil {
				e.Log.Errorf("Failed to delete file %s: %s", dataPath, err)
			}
			err = os.Remove(imagePath)
			if err != nil {
				e.Log.Errorf("Failed to delete file %s: %s", imagePath, err)
			}
		}
	default:
		return errors.New("无法识别的 ID 类型")
	}
	return nil
}

// 根据父区域id得到该区域下所有场景id
func (e *Scene) GetSceneIdsByLocId(c *dto.DLocJoin) ([]int, error) {
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

	var SceneIds []int
	Scenes := make([]models.Scene, 0)
	if err := e.Orm.Model(&models.Scene{}).Where("loc_id IN (?)", allLocIDs).Find(&Scenes).Error; err != nil {
		e.Log.Errorf("db error: %v", err)
		return nil, err
	}

	for _, Scene := range Scenes {
		SceneIds = append(SceneIds, Scene.SceneID)
	}

	return SceneIds, nil
}

// GetPageBySceneIds 根据提供的设备ID列表和分页信息获取设备数据
func (e *Scene) GetPageBySceneIds(sceneIds []int, pageReq *dto.SceneGetPageReqBySceneIds, p *actions.DataPermission, list *[]models.Scene, count *int64) error {
	// 检查SceneIds是否为空
	if len(sceneIds) == 0 {
		return errors.New("设备ID列表为空")
	}

	// 检查pageReq是否符合要求
	if pageReq == nil {
		return errors.New("分页请求对象为空")
	}

	// 使用传入的设备ID列表和分页请求查询设备数据
	err := e.Orm.Debug().
		Where("scene_id IN (?)", sceneIds).
		Scopes(
			cDto.Paginate(pageReq.GetPageSize(), pageReq.GetPageIndex()), // 应用分页
			actions.Permission("Scene", p),                               // 应用数据权限
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	return nil
}
