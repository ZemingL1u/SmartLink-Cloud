// apis包定义了SmartLinkProject的API接口
package apis

import (
	"SmartLinkProject/app/SLC/models" // 导入项目中的设备模型
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http" // 导入net包，用于处理HTTP相关功能
	"strconv"
	"time"

	"github.com/gin-gonic/gin"         // 导入Gin框架，用于处理HTTP请求
	"github.com/gin-gonic/gin/binding" // 导入Gin框架的binding功能，用于数据绑定

	"github.com/go-admin-team/go-admin-core/sdk/api" // 导入go-admin-core的api模块
	// 导入go-admin-core的response模块，虽然未使用，但可能用于响应处理
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"SmartLinkProject/app/SLC/service"     // 导入项目中的设备服务层
	"SmartLinkProject/app/SLC/service/dto" // 导入项目中的设备数据传输对象（DTO）

	"SmartLinkProject/common/actions" // 导入项目中的权限操作
	"SmartLinkProject/mqtt"
)

type Device struct {
	api.Api
}

func (e Device) GetPage(c *gin.Context) {
	// 实例化设备服务
	s := service.Device{}
	// 初始化设备分页请求数据传输对象（DTO）
	req := dto.DeviceGetPageReq{}
	// 绑定请求数据到req DTO，创建ORM和上下文，并将服务实例赋值给s
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		// 记录错误日志并返回500错误响应
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 从Gin上下文中获取当前用户的权限信息
	p := actions.GetPermissionFromContext(c)

	// 准备用于接收查询结果的切片和计数器变量
	list := make([]models.Device, 0)
	var count int64

	// 调用服务层的GetPage方法来获取设备分页列表和总记录数
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		// 如果服务层返回错误，返回500错误响应
		e.Error(500, err, "查询失败")
		return
	}
	// 返回分页成功的响应，包括设备列表、当前页码、每页大小和总记录数
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// 其他方法（如Get、Insert、Update、Delete、UpdateStatus、GetProfile）遵循类似的模式：
// 1. 实例化服务层对象
// 2. 初始化请求数据传输对象（DTO）
// 3. 绑定请求数据到DTO并创建服务层上下文
// 4. 执行服务层方法，并处理错误
// 5. 根据服务层执行结果返回相应的HTTP响应

func (e Device) Get(c *gin.Context) {
	s := service.Device{}
	req := dto.DeviceById{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.Device
	//数据权限检查
	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}

func (e Device) Insert(c *gin.Context) {
	s := service.Device{}
	req := dto.DeviceInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.Insert(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req.GetId(), "创建成功")
}

func (e Device) Update(c *gin.Context) {
	s := service.Device{}
	req := dto.DeviceUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), "更新成功")
}

func (e Device) Delete(c *gin.Context) {
	s := service.Device{}
	req := dto.DeviceById{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), "删除成功")
}

func (e Device) UpdateStatus(c *gin.Context) {
	s := service.Device{}
	req := dto.UpdateDeviceStatusReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.UpdateStatus(&req, p)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), "更新成功")
}

func (e Device) GetProfile(c *gin.Context) {
	s := service.Device{}
	req := dto.DeviceById{}

	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	device := models.Device{}
	err = s.GetProfile(&req, &device)
	if err != nil {
		e.Logger.Errorf("get device profile error, %s", err.Error())
		e.Error(500, err, "获取设备信息失败")
		return
	}
	e.OK(gin.H{
		"device": device,
	}, "查询成功")
}

func (e Device) GetDeviceIdsByLocId(c *gin.Context) {
	s := service.Device{}
	req := dto.DLocJoin{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	deviceIds, err := s.GetDeviceIdsByLocId(&req)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(deviceIds, "查询成功")
}

func (e Device) GetPageByDeviceIds(c *gin.Context) {
	s := service.Device{}
	req := dto.DeviceGetPageReqByDeviceIds{}

	// 绑定请求数据到req DTO，创建ORM和上下文，并将服务实例赋值给s
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 准备用于接收查询结果的切片和计数器变量
	list := make([]models.Device, 0)
	var count int64

	// 从Gin上下文中获取当前用户的权限信息
	p := actions.GetPermissionFromContext(c)

	// 调用服务层的GetPage方法来获取设备分页列表和总记录数
	// 注意：这里的实现取决于service包中Device服务的GetPage方法是否支持通过ID列表进行查询
	err = s.GetPageByDeviceIds(req.DeviceIds, &req, p, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	// 返回分页成功的响应，包括设备列表、当前页码、每页大小和总记录数
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// 随机生成指定范围内的数据点
func generateRandomData(min, max int, count int) []float64 {
	data := make([]float64, count)
	for i := 0; i < count; i++ {
		data[i] = rand.Float64()*(float64(max-min)+1) + float64(min)
	}
	return data
}

func predictData(data []float64, predictCount int) []float64 {
	if len(data) < 2 {
		return nil
	}

	var sumX, sumY, sumXSquare, sumXY float64
	sumY = 0
	for i, y := range data {
		sumX += float64(i)
		sumY += y
		sumXSquare += float64(i) * float64(i)
		sumXY += float64(i) * y
	}

	meanX := sumX / float64(len(data))
	meanY := sumY / float64(len(data))

	// 计算线性回归系数
	slope := (sumXY - (sumX * meanY)) / (sumXSquare - (sumX * meanX))
	intercept := meanY - (slope * meanX)

	// 预测后续数据点
	var predictions []float64
	lastIndex := len(data) - 1
	for i := 1; i <= predictCount; i++ {
		predictedY := slope*float64(lastIndex+i) + intercept
		predictions = append(predictions, predictedY)
	}
	return predictions
}

// GetTemp 函数生成数据并进行预测
func (e Device) GetTemp(c *gin.Context) {
	// 生成3组数据，每组7个
	hourData := generateRandomData(25, 32, 7)
	dayData := generateRandomData(25, 32, 7)
	monthData := generateRandomData(25, 32, 7)

	// 预测每组的后续3个数据点
	predictedHourData := predictData(hourData, 3)
	predictedDayData := predictData(dayData, 3)
	predictedMonthData := predictData(monthData, 3)

	// 准备JSON响应，包含过去7个数据点和未来3个预测数据点
	hourlyData := make(map[string][]float64)
	dailyData := make(map[string][]float64)
	monthlyData := make(map[string][]float64)

	// 填充hourlyData
	now := time.Now()
	currentHour := now.Hour()
	currentDay := now.Day()
	currentMonth := now.Month()

	// 填充hourlyData,dailyData,monthData
	// 填充hourlyData
	for i := -6; i <= 3; i++ {
		hour := (currentHour + i) % 24
		hourStr := fmt.Sprintf("%02d:00", hour)
		if i <= 0 {
			// 过去的时间点使用原始数据
			hourlyData[hourStr] = []float64{hourData[6+i]} // +6 因为 -7 到 -1 相当于从尾部添加
		} else {
			// 未来的时间点使用预测数据
			hourlyData[hourStr] = []float64{predictedHourData[i-1]}
		}
	}

	// 填充dailyData
	for i := -6; i <= 3; i++ {
		day := currentDay + i
		if day == 0 {
			day = 31 // 如果是月初的前几天，用月末的日期
		} else if day > 31 {
			day = 1 // 同理，如果超出31天，转到下个月的开始
		}
		dayStr := fmt.Sprintf("%02d", day)
		if i <= 0 {
			dailyData[dayStr] = []float64{dayData[6+i]}
		} else {
			dailyData[dayStr] = []float64{predictedDayData[i-1]}
		}
	}

	// 填充monthlyData
	for i := -6; i <= 3; i++ {
		month := int(currentMonth) + i
		if month == 0 {
			month = 12 // 如果是年初的前几个月，转到年末
		} else if month > 12 {
			month = 1 // 同理，如果超出12月，转到下一年的开始
		}
		monthStr := fmt.Sprintf("%02d", month)
		if i <= 0 {
			monthlyData[monthStr] = []float64{monthData[6+i]}
		} else {
			monthlyData[monthStr] = []float64{predictedMonthData[i-1]}
		}
	}
	// 准备最终的JSON响应
	response := map[string]interface{}{
		"hour":  hourlyData,
		"day":   dailyData,   // 根据需要填充
		"month": monthlyData, // 根据需要填充
	}

	// 将响应编码为JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		c.JSON(-1, gin.H{"error": "Error encoding JSON", "details": err.Error()})
		return
	}

	// 设置状态码和响应体
	c.Data(200, gin.MIMEJSON, jsonResponse)
}

// GetTemp 函数生成数据并进行预测
func (e Device) GetHumi(c *gin.Context) {
	// 生成3组数据，每组7个
	hourData := generateRandomData(40, 80, 7)
	dayData := generateRandomData(40, 80, 7)
	monthData := generateRandomData(40, 80, 7)

	// 预测每组的后续3个数据点
	predictedHourData := predictData(hourData, 3)
	predictedDayData := predictData(dayData, 3)
	predictedMonthData := predictData(monthData, 3)

	// 准备JSON响应，包含过去7个数据点和未来3个预测数据点
	hourlyData := make(map[string][]float64)
	dailyData := make(map[string][]float64)
	monthlyData := make(map[string][]float64)

	// 填充hourlyData
	now := time.Now()
	currentHour := now.Hour()
	currentDay := now.Day()
	currentMonth := now.Month()

	// 填充hourlyData,dailyData,monthData
	// 填充hourlyData
	for i := -6; i <= 3; i++ {
		hour := (currentHour + i) % 24
		hourStr := fmt.Sprintf("%02d:00", hour)
		if i <= 0 {
			// 过去的时间点使用原始数据
			hourlyData[hourStr] = []float64{hourData[6+i]} // +6 因为 -7 到 -1 相当于从尾部添加
		} else {
			// 未来的时间点使用预测数据
			hourlyData[hourStr] = []float64{predictedHourData[i-1]}
		}
	}

	// 填充dailyData
	for i := -6; i <= 3; i++ {
		day := currentDay + i
		if day == 0 {
			day = 31 // 如果是月初的前几天，用月末的日期
		} else if day > 31 {
			day = 1 // 同理，如果超出31天，转到下个月的开始
		}
		dayStr := fmt.Sprintf("%02d", day)
		if i <= 0 {
			dailyData[dayStr] = []float64{dayData[6+i]}
		} else {
			dailyData[dayStr] = []float64{predictedDayData[i-1]}
		}
	}

	// 填充monthlyData
	for i := -6; i <= 3; i++ {
		month := int(currentMonth) + i
		if month == 0 {
			month = 12 // 如果是年初的前几个月，转到年末
		} else if month > 12 {
			month = 1 // 同理，如果超出12月，转到下一年的开始
		}
		monthStr := fmt.Sprintf("%02d", month)
		if i <= 0 {
			monthlyData[monthStr] = []float64{monthData[6+i]}
		} else {
			monthlyData[monthStr] = []float64{predictedMonthData[i-1]}
		}
	}
	// 准备最终的JSON响应
	response := map[string]interface{}{
		"hour":  hourlyData,
		"day":   dailyData,   // 根据需要填充
		"month": monthlyData, // 根据需要填充
	}

	// 将响应编码为JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		c.JSON(-1, gin.H{"error": "Error encoding JSON", "details": err.Error()})
		return
	}

	// 设置状态码和响应体
	c.Data(200, gin.MIMEJSON, jsonResponse)
}

// controlLed 方法通过Gin的Context接收HTTP请求并处理
func (e *Device) ControlLed(c *gin.Context) {
	// 从请求中获取state参数，假设它作为查询参数传来
	state := c.DefaultQuery("state", "0") // 默认为0，如果未提供state参数

	// 将state参数转换为整数
	stateInt, err := strconv.Atoi(state)
	if err != nil {
		// 如果转换失败，返回错误响应
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state parameter"})
		return
	}

	// 根据state参数确定要发布的payload
	var payload []byte
	if stateInt == 1 {
		payload = []byte(`{"Led":1}`) // LED 开
	} else {
		payload = []byte(`{"Led":0}`) // LED 关
	}

	// 定义MQTT消息的主题
	topic := "Operate/3"

	// 使用MqttClient的Publish方法发送消息
	if token := mqtt.MqttClient.Publish(topic, 0, false, payload); token.Wait() && token.Error() != nil {
		// 如果发布失败，返回错误响应
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Publish error", "details": token.Error()})
	} else {
		// 如果发布成功，返回成功响应
		c.JSON(http.StatusOK, gin.H{"message": "Published to topic", "topic": topic, "payload": string(payload)})
	}
}

// controlLed 方法通过Gin的Context接收HTTP请求并处理
func (e *Device) ControlBeep(c *gin.Context) {
	// 从请求中获取state参数，假设它作为查询参数传来
	state := c.DefaultQuery("state", "0") // 默认为0，如果未提供state参数

	// 将state参数转换为整数
	stateInt, err := strconv.Atoi(state)
	if err != nil {
		// 如果转换失败，返回错误响应
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state parameter"})
		return
	}

	// 根据state参数确定要发布的payload
	var payload []byte
	if stateInt == 1 {
		payload = []byte(`{"Beep":1}`) // LED 开
	} else {
		payload = []byte(`{"Beep":0}`) // LED 关
	}

	// 定义MQTT消息的主题
	topic := "Operate/3"

	// 使用MqttClient的Publish方法发送消息
	if token := mqtt.MqttClient.Publish(topic, 0, false, payload); token.Wait() && token.Error() != nil {
		// 如果发布失败，返回错误响应
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Publish error", "details": token.Error()})
	} else {
		// 如果发布成功，返回成功响应
		c.JSON(http.StatusOK, gin.H{"message": "Published to topic", "topic": topic, "payload": string(payload)})
	}
}
