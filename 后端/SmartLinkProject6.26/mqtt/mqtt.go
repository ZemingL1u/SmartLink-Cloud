package mqtt

import (
	"SmartLinkProject/app/SLC/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MqttClient mqtt.Client

func InitMqttClient(broker string) {
	// 设置 MQTT 客户端选项
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID("go-admin-smartLink-dev")
	opts.SetUsername("server")
	opts.SetPassword("123456")
	opts.SetDefaultPublishHandler(messageHandler) // 设置消息处理器

	// 创建并启动 MQTT 客户端
	MqttClient = mqtt.NewClient(opts)

	// 连接到 MQTT 服务器
	if token := MqttClient.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("MQTT Connect error:", token.Error())
		os.Exit(1)
	} else {
		// 连接成功
		fmt.Println("Connected to MQTT broker")
	}
	InitialSubscription()
}
func InitialSubscription() {

	// 初始化数据库连接
	dsn := "root:Slc123456@tcp(120.79.59.158:3306)/admin?charset=utf8&parseTime=True&loc=Local&timeout=1000ms"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// 处理数据库连接错误
		panic("failed to connect to database")
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			// 处理错误，例如记录日志或打印错误信息
			log.Println("Error getting underlying database connection:", err)
			return
		}
		sqlDB.Close()
	}()

	// 定义一个切片来存储查询到的 device_id
	var deviceIds []int

	// 使用 GORM 的 Select 方法查询所有的 device_id
	db.Model(&models.Device{}).Where("deleted_at IS NULL").Select("device_id").Find(&deviceIds)
	if err != nil {
		panic(err)
	}

	// 为每个 ID 生成主题并订阅
	for _, id := range deviceIds {
		idStr := strconv.Itoa(id)
		topic := "getData/" + idStr
		// 订阅主题
		SubscribeToTopic(topic)
	}

}
func SubscribeToTopic(topic string) {
	if token := MqttClient.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println("Subscribe error:", token.Error())
	} else {
		fmt.Printf("Subscribed to topic: %s\n", topic)
	}
}
func UnsubscribeFromTopic(topic string) {
	// 取消订阅指定的主题
	if token := MqttClient.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		fmt.Println("Unsubscribe error:", token.Error())
	} else {
		fmt.Printf("Unsubscribed from topic: %s\n", topic)
	}
}
func messageHandler(client mqtt.Client, msg mqtt.Message) {
	payloadStr := string(msg.Payload())
	fmt.Printf("Received message: %s from topic: %s\n", payloadStr, msg.Topic())
	// 从主题中解析设备ID
	deviceID, err := ParseDeviceIDFromTopic(msg.Topic())
	if err != nil {
		fmt.Printf("Error parsing device ID from topic: %s\n", err)
		return
	}

	// 保存数据到OSS
	SaveData(payloadStr, msg.Topic())

	// 更新设备数据到数据库
	UpdateDeviceData(deviceID, payloadStr)

}

// SaveData 将收到的设备数据历史记录保存到oss
func SaveData(data, topic string) {
	endpoint := "https://oss-cn-shenzhen.aliyuncs.com"
	accessKeyId := "LTAI5tSQqUnkr4cWG5HNCUmM"
	accessKeySecret := "7WjY5dwGYjYGv9Gg06DHG6ay2JwgC0"
	bucketName := "smartlinkcloud-devicedata"

	// 创建OSSClient实例
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 从topic中提取DeviceID
	parts := strings.Split(topic, "/")
	deviceID := parts[len(parts)-1] // 最后一个部分是DeviceID

	// 将data解析为map
	var dataMap map[string]interface{}
	err = json.Unmarshal([]byte(data), &dataMap)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// 获取当前时间，并格式化为 "20060102_150405" 形式
	currentTime := time.Now().Format("20060102_150405")

	// 创建文件名，格式为 "DeviceID_当前时间"
	objectKey := fmt.Sprintf("Device%s_%s.json", deviceID, currentTime)

	// 将map转换回JSON字符串
	jsonData, err := json.Marshal(dataMap)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// 创建Bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 上传文件到OSS，这里直接使用内存中的数据
	err = bucket.PutObject(objectKey, strings.NewReader(string(jsonData)))
	if err != nil {
		fmt.Println("Error uploading object:", err)
		return
	}

}

// UpdateDeviceData 更新设备数据到数据库mysql
func UpdateDeviceData(deviceID int, data string) {
	// 初始化数据库连接
	dsn := "root:Slc123456@tcp(120.79.59.158:3306)/admin?charset=utf8&parseTime=True&loc=Local&timeout=1000ms"
	var db *gorm.DB
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// 处理数据库连接错误
		panic("failed to connect to database")
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			// 处理错误，例如记录日志或打印错误信息
			log.Println("Error getting underlying database connection:", err)
			return
		}
		sqlDB.Close()
	}()

	// 使用First方法查找特定的设备
	if err := db.Where("device_id = ?", deviceID).First(&models.Device{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果记录未找到，可以选择返回错误或进行其他处理
			fmt.Println("Device not found")
		} else {
			// 处理其他错误
			fmt.Println("Error occurred while finding the device:", err)
		}
		return
	}

	// 更新device_data字段
	if err := db.Model(&models.Device{}).Where("device_id = ?", deviceID).Updates(map[string]interface{}{"device_data": data}).Error; err != nil {
		// 处理更新错误
		fmt.Println("Error occurred while updating the device data:", err)
		return
	}

}

// ParseDeviceIDFromTopic 尝试从 MQTT 主题中解析出设备 ID 并转换为 int 型
func ParseDeviceIDFromTopic(topic string) (int, error) {
	// 检查主题是否符合预期格式
	if !strings.HasPrefix(topic, "getData/") {
		return 0, fmt.Errorf("topic does not start with 'getData/'")
	}

	// 移除前缀并获取设备 ID 字符串
	deviceIDStr := strings.TrimPrefix(topic, "getData/")

	// 检查设备 ID 是否有效（例如，非空）
	if deviceIDStr == "" {
		return 0, fmt.Errorf("empty device ID in topic")
	}

	// 将设备 ID 字符串转换为 int 类型
	deviceID, err := strconv.Atoi(deviceIDStr)
	if err != nil {
		return 0, fmt.Errorf("invalid device ID format: %s", err.Error())
	}

	return deviceID, nil
}
